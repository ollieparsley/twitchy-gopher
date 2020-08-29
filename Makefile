default: help

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  deps                Fetch dependencies"
	@echo "  test                Run unit tests"
	@echo "  coverage            Run test coverage"
	@echo "  coverage-travis-ci  Run test coverage specific to travis ci"

deps:
	export GO111MODULE=1
	go get -v ./...
	go get -v github.com/stretchr/testify/assert@v1.2.2
	go get -v github.com/stretchr/testify/mock@v1.2.2
	go get -v github.com/jarcoal/httpmock
	go get -v github.com/mattn/goveralls
	go get -u golang.org/x/tools/...

test:
	go test -v ./twitch

coverage:
	go test -cover ./twitch

coverage-travis-ci:
	$(GOPATH)/bin/goveralls -service=travis-ci
