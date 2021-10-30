
GIT_COMMIT := $(shell git describe --long --always 2> /dev/null)

build: icons wasm binary

run: build
	VERSION=dynamic ./bin/example

static: build
	rm -rf static
	GEN_STATIC="true" ./bin/example
	cp -a web/* static/web

icons:
	go run scripts/material/generate.go -output pkg/icon/material.go -package icon

upload: static
	aws s3 sync --delete --exclude web/app.wasm ./static/ s3://mlctrez-goapp-mdc/
	aws s3 cp ./static/web/app.wasm s3://mlctrez-goapp-mdc/web/app.wasm --content-encoding br

wasm:
	GOARCH=wasm GOOS=js go build -ldflags="-s -w" -o web/app.wasm
	rm -f web/app.wasm.br
	brotli -q 6 -j web/app.wasm
	mv web/app.wasm.br web/app.wasm

binary:
	mkdir -p bin
	go build -ldflags "-X main.GitCommit=$(GIT_COMMIT)" -o bin/example

