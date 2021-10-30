package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
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
			//"github.com/mlctrez/goapp-mdc/pkg/jsutil",
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
		t.WriteString(fmt.Sprintf(TestSource,name,strings.Title(name),strings.Title(name)))

		err = WriteSource(filepath.Join(packageDir, fmt.Sprintf("%s_test.go", name)), t)
		if err != nil {
			log.Fatalln(err)
		}

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
	"github.com/mlctrez/goapp-mdc/pkg/internal/componenttest"
	"testing"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Test%s_Render(t *testing.T) {
	id := "testID"
	compo := &%s{Id: id}
	ct := componenttest.ComponentTest{T: t, Compo: compo, Dispatcher: app.NewServerTester(compo)}
	defer ct.Close()

}
`