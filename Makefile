# https://www.gnu.org/software/make/manual/html_node/Automatic-Variables.html

VERSION ?= $$(./script/ci.sh version)
BUILD_FLAGS := -ldflags "\
	      -X \"main.Version=$(VERSION)\" \
	      "

GO_FILES := $(shell find . -type f -name "*.go")
GO_TESTED := ./report/test-go.txt

.PHONY: init dependencies install build test lint release

init: install lint build test

install:
	go get -t -v

build:
	go build -v $(BUILD_FLAGS)

test: $(GO_TESTED)

lint:
	gofmt -s -w .
	go vet .

release:
	go get -v github.com/mitchellh/gox
	VERSION="$(VERSION)" ASSET_DIR="$(ASSET_DIR)" ./script/release.sh
	touch pkg/.gitkeep

$(GO_TESTED): $(GO_FILES)
	go test -v ./...
	touch $(GO_TESTED)

include script/*.mk
