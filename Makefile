GOPATH=$(shell readlink -f $(shell pwd)/../../../../)

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  docs                    to build SDK documentation"
	@echo "  test                    to run unit tests"
	@echo "  deps                    to go get the dependencies"

deps:
	GOPATH=$(GOPATH) go get -v ./...
	GOPATH=$(GOPATH) go get -v github.com/stretchr/testify/assert
	GOPATH=$(GOPATH) go get -v github.com/stretchr/testify/mock

test: deps
	GOPATH=$(GOPATH) go test ./twitch

docs:
	@echo "Generate documentation docs"
	@mkdir -p target
	GOPATH=$(GOPATH) godoc -html github.com/ollieparsley/twitchy-gopher > target/docs.html
