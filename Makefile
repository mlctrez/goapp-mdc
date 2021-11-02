
GIT_COMMIT := $(shell git describe --long --always 2> /dev/null)
SCR := scripts
PROG := bin/goappmdc

build: icons markup wasm binary

run: build
	$(PROG) dev

static: build bin/genstatic
	rm -rf static
	bin/genstatic
	cp -a web/* static/web

icons: bin/material
	bin/material -output pkg/icon/material.go -package icon

upload: bin/upload
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

bin/material: $(SCR)/material/*.go
	cd $(SCR) && go build -o ../$@ ../$<

bin/newpkg: $(SCR)/newpkg/*.go
	cd $(SCR) && go build -o ../$@ ../$<

bin/upload: $(SCR)/upload/*.go
	cd $(SCR) && go build -o ../$@ ../$<

bin/genstatic: $(SCR)/genstatic/*.go ./demo/*.go
	cd $(SCR) && go build -o ../$@ ../$<

bin/markup: $(SCR)/markup/*.go
	cd $(SCR) && go build -o ../$@ ../$<