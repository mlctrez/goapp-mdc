package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func main() {
	var output string
	flag.StringVar(&output, "output", "demo/markup/code.go", "the output file")

	fmt.Println(output)

	getwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var paths []string

	err = filepath.Walk(getwd, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".go") {
			var rel string
			rel, err = filepath.Rel(getwd, path)
			if err != nil {
				return err
			}
			// add root files
			if !strings.Contains(rel, "/") {
				paths = append(paths, rel)
			}
			if strings.HasPrefix(rel, "demo/") &&
				!strings.HasPrefix(rel, "demo/older") &&
				!strings.HasPrefix(rel, "demo/markup") {
				paths = append(paths, rel)
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Dir(output), 0755)
	if err != nil {
		panic(err)
	}

	open, err := os.Create(output)
	if err != nil {
		panic(err)
	}

	buff := bytes.Buffer{}
	buff.WriteString("package markup\n")

	buff.WriteString("var Code = map[string]string{\n")

	for _, path := range paths {
		buf := bytes.Buffer{}
		buf.WriteString("```go\n")
		file, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		buf.Write(file)
		buf.WriteString("```\n")
		p := parser.NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs)
		html := markdown.ToHTML(buf.Bytes(), p, nil)

		buff.WriteString(fmt.Sprintf("%q:%s", path, "`"+string(html)+"`,\n"))
		_ = html
	}
	buff.WriteString("}\n")

	_, err = open.Write(buff.Bytes())
	if err != nil {
		panic(err)
	}

}
