package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/mlctrez/goapp-mdc/demo"
)

func main() {
	var output string
	flag.StringVar(&output, "output", "demo/markup/code.go", "the output file")

	fmt.Println(output)

	getwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	paths := make(map[string]bool)

	err = filepath.Walk(getwd, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".go") {
			var rel string
			rel, err = filepath.Rel(getwd, path)
			if err != nil {
				return err
			}
			// add root files
			if !strings.Contains(rel, "/") {
				paths[rel] = true
			}
			if strings.HasPrefix(rel, "demo/") &&
				!strings.HasPrefix(rel, "demo/older") &&
				!strings.HasPrefix(rel, "demo/markup") {
				paths[rel] = true
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

	var orderedPaths []string

	demo.Routes()
	// navigation items as they appear on the page, match these
	// with the demo code path based on demo/demo_<name>.go
	// where <name> will be the href
	for _, item := range demo.NavigationItems {
		cp := strings.TrimPrefix(item.Href, "/")
		if cp == "" {
			cp = "index"
		}
		codePath := fmt.Sprintf("demo/demo_%s.go", cp)
		if paths[codePath] {
			delete(paths, codePath)
			orderedPaths = append(orderedPaths, codePath)
		}
	}
	for k := range paths {
		if strings.HasPrefix(k, "demo/") {
			delete(paths, k)
			orderedPaths = append(orderedPaths, k)
		}
	}

	var mainPackage []string
	for k := range paths {
		mainPackage = append(mainPackage, k)
	}
	sort.Strings(mainPackage)
	orderedPaths = append(orderedPaths, mainPackage...)

	open, err := os.Create(output)
	if err != nil {
		panic(err)
	}

	buff := bytes.Buffer{}
	buff.WriteString("package markup\n")

	buff.WriteString("type CodeDetails struct {\n")
	buff.WriteString("	Name string\n")
	buff.WriteString("	Code string\n")
	buff.WriteString("}\n")

	for i, path := range orderedPaths {
		buff.WriteString(fmt.Sprintf("// %2d %s\n", i, path))
	}

	buff.WriteString("var Code = []CodeDetails{\n")

	for _, path := range orderedPaths {
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

		path = strings.Replace(path, "demo/", "", 1)
		path = strings.Replace(path, "demo_", "", 1)

		buff.WriteString(fmt.Sprintf("    CodeDetails{Name:%q,Code:`%s`},\n", path, string(html)))
	}
	buff.WriteString("}\n")

	_, err = open.Write(buff.Bytes())
	if err != nil {
		panic(err)
	}

}
