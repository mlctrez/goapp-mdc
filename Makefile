
GIT_COMMIT := $(shell git describe --long --always 2> /dev/null)
SCR := scripts
PROG := bin/goappmdc

build: icons wasm binary

run: build
	VERSION=dynamic $(PROG)

static: build
	rm -rf static
	go run scripts/genstatic/genstatic.go
	cp -a web/* static/web

icons: bin/material
	bin/material -output pkg/icon/material.go -package icon

upload: static bin/upload
	bin/upload static mlctrez-goapp-mdc

wasm:
	GOARCH=wasm GOOS=js go build -ldflags="-s -w" -o web/app.wasm
	ls -l web/app.wasm
#	rm -f web/app.wasm.br
#	brotli -q 6 -j web/app.wasm
#	mv web/app.wasm.br web/app.wasm

binary: bin
	go build -ldflags "-X main.GitCommit=$(GIT_COMMIT)" -o $(PROG)

bin:
	mkdir -p bin

bin/material: $(SCR)/material/*.go
	cd $(SCR) && go build -o ../$@ ../$<

bin/newpkg: $(SCR)/newpkg/*.go
	cd $(SCR) && go build -o ../$@ ../$<

bin/upload: $(SCR)/upload/*.go
	cd $(SCR) && go build -o ../$@ ../$<
