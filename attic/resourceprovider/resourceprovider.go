package resourceprovider

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ResourceProvider struct {
}

func (r ResourceProvider) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if strings.HasPrefix(request.RequestURI, "/web/") {
		file, err := os.Open(request.RequestURI[1:])
		if os.IsNotExist(err) {
			writer.WriteHeader(404)
			return
		}
		ext := filepath.Ext(request.RequestURI)

		contentType := mime.TypeByExtension(ext)
		writer.Header().Set("Content-Type", contentType)
		if ext == ".wasm" {
			writer.Header().Set("Content-Encoding", "br")
		}

		io.Copy(writer, file)
		return
	}

	fmt.Println("404")
	writer.WriteHeader(404)
}

func (r ResourceProvider) Package() string { return "" }
func (r ResourceProvider) Static() string  { return "" }
func (r ResourceProvider) AppWASM() string { return "/web/app.wasm" }

var _ app.ResourceProvider = (*ResourceProvider)(nil)
