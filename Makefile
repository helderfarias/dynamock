# General
PKG_LIST_ALL_TESTS	:= $(shell go list ./... | grep -v /vendor)
GIT_BRANCH			:= $(shell git symbolic-ref HEAD | sed -e 's/^refs\/heads\///')
GIT_LAST_COMMIT		:= $(shell git rev-parse --short HEAD)

# Version
VMAJOR_MINOR 		:= $(or ${VMAJOR_MINOR}, ${GIT_BRANCH})
VBUILD 				:= $(or ${VBUILD}, 0)
VREV 				:= $(or ${VREV}, ${GIT_LAST_COMMIT})
VERSION				:= "$(VMAJOR_MINOR).$(VBUILD).$(shell echo ${VREV} | cut -c 1-8)"


all: help

version:
	@echo "$(VERSION)"

static:
	@echo ">> runnung static analysis"
	@go get -v honnef.co/go/tools/cmd/staticcheck
	@cd api && staticcheck . && cd ..

ast:
	@echo ">> inspecting source code for security problems"
	@go get -v github.com/GoASTScanner/gas/cmd/gas/...
	@cd api && gas . && cd ..

test:
	@echo ">> running unit tests"
	@go test -cover $(PKG_LIST_ALL_TESTS)

build: build_alpine

docker:
	@echo ">> docker image"
	docker build -t helderfarias/dynamock .

build_alpine:
	@echo ">> release for linux alpine"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -a -installsuffix cgo -o release/alpine/dynamock

build_linux:
	@echo ">> release for linux"
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o release/linux/dynamock

build_osx:
	@echo ">> release for osx"
	GOOS=darwin go build -ldflags "-X main.version=$(VERSION)" -o release/osx/dynamock

help:
	@echo 'Usage: '
	@echo ''
	@echo 'make build'
	@echo 'make test'
	@echo 'make ast'
	@echo 'make static'
	@echo 'make build_alpine'
	@echo 'make build_linux'
	@echo 'make build_osx'
