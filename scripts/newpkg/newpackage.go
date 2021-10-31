package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// generates folder and starter files for a new package
func main() {
	flag.Parse()
	name := flag.Arg(0)

	if name == "" {
		log.Fatalln("please provide single word package name as first argument")
	}

	packageDir := filepath.Join("pkg", name)
	_, err := os.Stat(packageDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(packageDir, 0755)
		if err != nil {
			log.Fatalln(err)
		}
		s := &bytes.Buffer{}
		imports := []string{
			"github.com/maxence-charriere/go-app/v9/pkg/app",
		}

		s.WriteString("package " + name + "\n")
		s.WriteString("import (\n")
		for i := range imports {
			s.WriteString(fmt.Sprintf("\t%q\n", imports[i]))
		}
		s.WriteString(")\n")
		s.WriteString(fmt.Sprintf("type %s struct {\n", strings.Title(name)))
		s.WriteString("\tapp.Compo\n")
		s.WriteString("\tId string\n")
		s.WriteString("}\n")
		s.WriteString(fmt.Sprintf("func (%s *%s)Render() app.UI {\n return app.Div()\n}\n", name[0:1], strings.Title(name)))
		s.WriteString(fmt.Sprintf("func (%s *%s)OnMount(_ app.Context) {\n}\n", name[0:1], strings.Title(name)))

		err = WriteSource(filepath.Join(packageDir, fmt.Sprintf("%s.go", name)), s)
		if err != nil {
			log.Fatalln(err)
		}
		t := &bytes.Buffer{}
		t.WriteString(fmt.Sprintf(TestSource, name, strings.Title(name), strings.Title(name)))

		err = WriteSource(filepath.Join(packageDir, fmt.Sprintf("%s_test.go", name)), t)
		if err != nil {
			log.Fatalln(err)
		}
		create, err := os.Create(filepath.Join(packageDir, "index.html"))
		if err != nil {
			panic(err)
		}
		io.Copy(create, bytes.NewBufferString(IndexSource))
		create.Close()

	} else {
		log.Fatalln(packageDir, "exists already")
	}
}

func WriteSource(path string, buffer *bytes.Buffer) (err error) {
	var formattedSource []byte
	if formattedSource, err = format.Source(buffer.Bytes()); err != nil {
		return
	}
	return os.WriteFile(path, formattedSource, 0755)
}

const TestSource = `
package %s

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
)

func Test%s_Render(t *testing.T) {
	id := "testID"
	compo := &%s{Id: id}
	ct := base.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

}
`

const IndexSource = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <link rel="stylesheet" href="https://unpkg.com/material-components-web@v13.0.0/dist/material-components-web.css"
          type="text/css">
    <script src="https://unpkg.com/material-components-web@13.0.0/dist/material-components-web.js"></script>
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons" type="text/css">

</head>
<body>


<script>

</script>

</body>
</html>`
