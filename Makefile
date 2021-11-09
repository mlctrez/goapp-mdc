
GIT_COMMIT := $(shell git describe --long --always 2> /dev/null)
PROG := bin/goappmdc
GOFILES := $(shell find . -name "*.go" -print)

build: icons markup wasm binary

run: build
	$(PROG) dev

static: build bin/genstatic
	rm -rf static
	bin/genstatic
	cp -a web/* static/web

dynstatic: bin/genstatic
	bin/genstatic
	cp -a web/* static/web

dynserver: bin/dynserver bin/genstatic
	bin/genstatic
	cp -a web/* static/web
	bin/dynserver

icons: bin/material
	bin/material -output pkg/icon/material.go -package icon

upload: bin/upload static
	bin/upload static mlctrez-goapp-mdc

markup: bin/markup ./*.go demo/*.go
	bin/markup -output demo/markup/code.go

wasm:
	GOARCH=wasm GOOS=js go build -ldflags="-s -w" -o web/app.wasm

binary: bin
	go build -ldflags "-X main.GitCommit=$(GIT_COMMIT)" -o $(PROG)

fmt:
	go fmt ./...

bin:
	mkdir -p bin

bin/material: scripts/material/*.go
	go build -o $@ $<

bin/newpkg: scripts/newpkg/*.go
	go build -o $@ $<

bin/upload: scripts/upload/*.go
	go build -o $@ $<

bin/genstatic: scripts/genstatic/*.go $(GOFILES)
	go build -o $@ $<

bin/markup: scripts/markup/*.go
	go build -o $@ $<

bin/dynserver: scripts/dynserver/*.go
	go build -o $@ $<