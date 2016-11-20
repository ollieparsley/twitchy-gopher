GOPATH=$(shell readlink -f $(shell pwd)/../../../../)

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  docs                    to build SDK documentation"
	@echo "  test                    to run unit tests"
	@echo "  deps                    to go get the dependencies"

deps:
	GOPATH=$(GOPATH) go get -v -u ./...
	GOPATH=$(GOPATH) go get -v -u github.com/golang/lint/golint
	GOPATH=$(GOPATH) go get -v -u github.com/jstemmer/go-junit-report
	GOPATH=$(GOPATH) go get -v -u github.com/stretchr/testify/assert
	GOPATH=$(GOPATH) go get -v -u github.com/stretchr/testify/mock

test: deps
	GOPATH=$(GOPATH) go test

docs:
	@echo "Generate documentation docs"
	@mkdir -p target
	GOPATH=$(GOPATH) godoc -html github.com/ollieparsley/twitchy-gopher > target/docs.html
