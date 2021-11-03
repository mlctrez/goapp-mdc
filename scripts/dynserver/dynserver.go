package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

type ModifyResponseWriter struct {
	bytes.Buffer
	status int
	header http.Header
}

func (m *ModifyResponseWriter) Header() http.Header {
	return m.header
}

func (m *ModifyResponseWriter) WriteHeader(statusCode int) {
	m.status = statusCode
}

// ReWrite filters lines from html files matching the provided filterPattern.
func (m *ModifyResponseWriter) ReWrite(url string, resp http.ResponseWriter, filterPattern *regexp.Regexp) {

	if strings.Contains(m.header.Get("Content-Type"), "text/html") {
		newContent := &bytes.Buffer{}
		scanner := bufio.NewScanner(m)
		for scanner.Scan() {
			line := scanner.Text()
			if !filterPattern.MatchString(line) {
				newContent.WriteString(line + "\n")
			}
			if line == "</head>" {
				newContent.WriteString(`<style> #app-wasm-loader {	display: none; } </style>` + "\n")
			}
		}
		delete(m.header, "Accept-Ranges")
		delete(m.header, "Content-Length")
		m.header.Set("Content-Length", fmt.Sprintf("%d", newContent.Len()))
		m.Reset()
		m.Write(newContent.Bytes())
	}
	resp.WriteHeader(m.status)
	for key, strings := range m.header {
		for _, value := range strings {
			resp.Header().Set(key, value)
		}
	}
	resp.Write(m.Bytes())

}

func watchFiles(dirs string, cmd string) {

	var validDirs []string

	// test directories first
	for _, s := range strings.Split(dirs, ",") {
		dir := strings.TrimSpace(s)
		_, err := os.Stat(dir)
		if err != nil {
			log.Fatal("directory", dir, "cannot be watched:", err)
		}
		validDirs = append(validDirs, dir)
	}

	var watcher *fsnotify.Watcher
	var err error
	if watcher, err = fsnotify.NewWatcher(); err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	var runCmd = false

	go func() {
		for {
			if runCmd == true {
				time.Sleep(500 * time.Millisecond)
				split := strings.Split(cmd, " ")
				log.Println("running", cmd)
				output, err := exec.Command(split[0], split[1:]...).CombinedOutput()
				if err != nil {
					log.Println(string(output))
					log.Println(err)
				}
				runCmd = false
			}
			time.Sleep(100 * time.Millisecond)
		}

	}()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Println("watcher.Events select returned !ok")
					return
				}
				//log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					if !strings.HasSuffix(event.Name, "~") {
						//log.Println("modified file:", event.Name)
						runCmd = true
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					log.Println("watcher.Errors select returned !ok")
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	for _, dir := range validDirs {
		walkErr := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return watcher.Add(dir)
			}
			return nil
		})
		if walkErr != nil {
			log.Fatal(walkErr)
		}
	}
	<-done

}

func main() {
	var dir string
	var gen string
	var watch string
	flag.StringVar(&dir, "dir", "static", "directory to serve files from")
	flag.StringVar(&gen, "gen", "make dynstatic", "command to execute to re-generate")
	flag.StringVar(&watch, "watch", "pkg,demo", "directories to watch for file changes")
	flag.Parse()

	go watchFiles(watch, gen)

	fileServer := http.FileServer(http.Dir(dir))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if "/favicon.ico" == request.RequestURI {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		// re-write page urls without extension
		if request.RequestURI != "/" && filepath.Ext(request.RequestURI) == "" {
			request.RequestURI = request.RequestURI + ".html"
			parse, err := url.Parse(request.URL.String() + ".html")
			if err != nil {
				panic(err)
			}
			request.URL = parse
		}

		mrw := &ModifyResponseWriter{header: make(http.Header)}
		fileServer.ServeHTTP(mrw, request)
		mrw.ReWrite(request.RequestURI, writer, regexp.MustCompile(".*/manifest\\.webmanifest.*|.*/wasm_exec\\.js.*|.*\"/app\\.js.*"))
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
