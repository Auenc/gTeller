default: run

name = gTeller
package = github.com/auenc/$(name)
ifndef $(GOPATH)
GOPATH = /home/auenc/Documents/workspace/go/
endif
.PHONY: default run

clean:
	GOPATH=$(GOPATH) go clean

run:format
	CGO_ENABLED=0 go run $(name).go

test:
	GOPATH=$(GOPATH) CGO_ENABLED=0 go test -i $(package)/...
	GOPATH=$(GOPATH) CGO_ENABLED=0 go test $(package)/...

docgen:
	echo godoc -goroot=$(GOPATH) -html src/$(package)/...

build:test
	GOPATH=$(GOPATH) go build -v

install:test
		GOPATH=$(GOPATH) go install -v

format:
	GOPATH=$(GOPATH) go fmt
