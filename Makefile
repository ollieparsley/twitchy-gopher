GOPATH ?= $(shell readlink -f $(shell pwd)/../../../../)

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  docs                    to build SDK documentation"
	@echo "  test                    to run unit tests"
	@echo "  deps                    to go get the dependencies"

deps:
	GOPATH=$(GOPATH) go get -v ./...
	GOPATH=$(GOPATH) go get -v github.com/stretchr/testify/assert
	GOPATH=$(GOPATH) go get -v github.com/stretchr/testify/mock
	GOPATH=$(GOPATH) go get -v github.com/jarcoal/httpmock
	GOPATH=$(GOPATH) go get -v github.com/mattn/goveralls

test: deps
	GOPATH=$(GOPATH) go test -v ./twitch

coverage: deps
	GOPATH=$(GOPATH) go test -cover ./twitch

coverage-travis-ci: deps
	$(GOPATH)/bin/goveralls -service=travis-ci

docs:
	@echo "Generate documentation docs"
	@mkdir -p target
	GOPATH=$(GOPATH) godoc -html github.com/ollieparsley/twitchy-gopher > target/docs.html
