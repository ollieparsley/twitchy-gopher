GOBIN = go
#GOBIN = go1.11.5
#GOBIN = go1.12.5
#GOBIN = go1.13.5
#GOBIN = go1.14.5

default: help

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  deps                Fetch dependencies"
	@echo "  test                Run unit tests"
	@echo "  coverage            Run test coverage"
	@echo "  coverage-travis-ci  Run test coverage specific to travis ci"

deps:
	GO111MODULE=on $(GOBIN) get -t -v ./...
	GO111MODULE=on $(GOBIN) get -v github.com/mattn/goveralls@v0.0.4

test:
	$(GOBIN) test -count=1 -v ./twitch

coverage:
	$(GOBIN) test -count=1 -cover ./twitch

coverage-travis-ci:
	$(GOPATH)/bin/goveralls -service=travis-ci
