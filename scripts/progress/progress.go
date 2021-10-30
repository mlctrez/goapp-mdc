package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	glob, err := filepath.Glob("/home/mattman/WebstormProjects/material-components-web/packages/*")
	if err != nil {
		panic(err)
	}

	create, err := os.Create("PackageProgress.md")
	if err != nil {
		panic(err)
	}

	create.WriteString(`## Package Mapping

| [material-components-web](https://github.com/material-components/material-components-web) | [goapp-mdc](https://github.com/mlctrez/goapp-mdc/) | Notes |
| --- | --- | --- |
`)

	for i, s := range glob {
		_ = i
		stat, err := os.Stat(s)
		if err != nil {
			panic(err)
		}

		if stat.IsDir() {

			mdcPkg := filepath.Base(s)
			href := "https://github.com/material-components/material-components-web/tree/master/packages/" + mdcPkg
			goappPkg := strings.TrimPrefix(mdcPkg, "mdc-")
			goappPkgDir := filepath.Join("pkg", goappPkg)

			mdcCol := fmt.Sprintf("[%s](%s)", mdcPkg, href)
			goappCol := fmt.Sprintf("[%s](%s)", goappPkg, goappPkgDir)
			_, err := os.Stat(goappPkgDir)
			if os.IsNotExist(err) {
				goappCol = goappPkgDir
			}

			create.WriteString("| " + mdcCol + " | " + goappCol + " | - |\n")

		}

	}
	create.Close()

}
