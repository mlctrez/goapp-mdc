//go:build ignore
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func snakeToCamelCase(in string) string {
	var output string
	parts := strings.Split(in, "_")
	for _, part := range parts {
		output += strings.Title(part)
	}
	return output
}

func findLeadingBracket(b *bytes.Buffer) (err error) {
	if _, err = b.ReadBytes('{'); err != nil {
		return
	}
	err = b.UnreadByte()
	return
}

func readAndCache(dir, path, url string) (data io.ReadCloser, err error) {
	var userCacheDir, ourCacheDir, cacheFile = "", "", ""
	if userCacheDir, err = os.UserCacheDir(); err != nil {
		return
	}
	ourCacheDir = filepath.Join(userCacheDir, dir)
	if err = os.MkdirAll(ourCacheDir, 0755); err != nil {
		return
	}
	cacheFile = filepath.Join(ourCacheDir, path)

	if _, err = os.Stat(cacheFile); os.IsNotExist(err) {

		fmt.Println("creating cache file", cacheFile)

		// cache file does not exist, download it and save
		var response *http.Response
		if response, err = http.Get(url); err != nil {
			return
		}
		buffer := &bytes.Buffer{}
		if _, err = io.Copy(buffer, response.Body); err != nil {
			return
		}
		if err = findLeadingBracket(buffer); err != nil {
			return
		}
		if err = os.WriteFile(cacheFile, buffer.Bytes(), 0755); err != nil {
			return
		}
	}
	return os.Open(cacheFile)
}

type IconMetadata struct {
	Host            string   `json:"host"`
	AssetUrlPattern string   `json:"asset_url_pattern"`
	Families        []string `json:"families"`
	Icons           []*Icon  `json:"icons"`
}

type Icon struct {
	Name       string   `json:"name"`
	Version    uint16   `json:"version"`
	Popularity uint32   `json:"popularity"`
	Codepoint  uint16   `json:"codepoint"`
	Categories []string `json:"categories"`
	Tags       []string `json:"tags"`
	SizesPx    []uint16 `json:"sizes_px"`
}

const FontMetadataUrl = "https://fonts.google.com/metadata/icons"
const CacheFile = "material-icons-metadata.json"
const CacheSubDir = "goapp_cache"

func getIconMetadata() (md *IconMetadata, err error) {
	var cachedData io.ReadCloser
	if cachedData, err = readAndCache(CacheSubDir, CacheFile, FontMetadataUrl); err != nil {
		return
	}
	defer cachedData.Close()
	md = &IconMetadata{}
	err = json.NewDecoder(cachedData).Decode(md)
	return
}

func prepareIconData() (sortedGroups []string, groupMap map[string][]string, err error) {

	metadata, err := getIconMetadata()
	if err != nil {
		panic(err)
	}

	// map of icon names by group
	groupMap = make(map[string][]string)
	for _, k := range metadata.Icons {
		for _, category := range k.Categories {
			groupMap[category] = append(groupMap[category], k.Name)
		}
	}

	for k := range groupMap {
		sortedGroups = append(sortedGroups, k)
	}
	sort.Strings(sortedGroups)

	for k := range groupMap {
		//items := groupMap[k]
		sort.Strings(groupMap[k])
		//groupMap[k] = items
	}

	return
}

type SourceWriter struct {
	bytes.Buffer
}

func (sw *SourceWriter) Println(line string) {
	_, err := sw.WriteString(line + "\n")
	if err != nil {
		// bytes.Buffer
		panic(err)
	}
}

func (sw *SourceWriter) WriteSource(path string) error {

	source, err := format.Source(sw.Bytes())
	if err != nil {
		return err
	}

	return os.WriteFile(path, source, 0755)
}

func main() {
	var output string
	var pkg string
	flag.StringVar(&output, "output", "pkg/icon/material.go", "output file icon constants")
	flag.StringVar(&pkg, "package", "icon", "output file package name")
	flag.Parse()

	sortedGroups, groupMap, err := prepareIconData()
	if err != nil {
		panic(err)
	}

	w := &SourceWriter{}
	w.Println(fmt.Sprintf("package %s", pkg))
	w.Println("// DO NOT EDIT - generated from " + FontMetadataUrl)
	w.Println("")
	w.Println("type MaterialIcon string")
	w.Println("")

	for _, g := range sortedGroups {

		for _, icon := range groupMap[g] {
			w.Println(fmt.Sprintf("const MI%s MaterialIcon = %q", snakeToCamelCase(icon), icon))
		}

		w.Println(fmt.Sprintf("func All%s() []MaterialIcon {", snakeToCamelCase(g)))
		w.Println("\t return []MaterialIcon{")
		for _, icon := range groupMap[g] {
			w.Println(fmt.Sprintf("\t\tMI%s,", snakeToCamelCase(icon)))
		}
		w.Println("\t}")
		w.Println("}")

	}

	w.Println("")
	w.Println("func AllGroupFunctions() map[string]func() []MaterialIcon {")
	w.Println("\t return map[string]func() []MaterialIcon{")
	for _, g := range sortedGroups {
		w.Println(fmt.Sprintf("\t%q: All%s,", g, snakeToCamelCase(g)))
	}
	w.Println("\t}")
	w.Println("}")
	w.Println("")
	w.Println("var _ = AllGroupFunctions")

	err = w.WriteSource(output)
	if err != nil {
		panic(err)
	}

}
