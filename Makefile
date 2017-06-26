GO_WORKSPACE := '/go/src/github.com/helderfarias/dynamock'

default: clean build

clean:
	@echo "Clean all..."
	@rm dynamock_*.zip || true
	@rm dynamock || true
	@rm dynamock.exe || true

build: linux_32 windows_32 windows_64 darwin_osx
	@echo "Complete"

docker:
	@echo "Building for alpine linux..."
	@docker run --rm -ti -v $(PWD):$(GO_WORKSPACE) -w $(GO_WORKSPACE) helderfarias/go:1.8-alpine go build -o dynamock_alpine
	@docker build -t helderfarias/dynamock .
	@rm -f dynamock_alpine || true

linux_64:
	@echo "Building for linux 64bits..."
	@GOOS=linux GOARCH=amd64 go build && zip dynamock_linux_64.zip dynamock

linux_32:
	@echo "Building for linux 32bits..."
	@GOOS=linux GOARCH=386 go build && zip dynamock_linux_32.zip dynamock

windows_32:
	@echo "Building for windows 32bits..."
	@GOOS=windows GOARCH=386 go build && zip dynamock_windows_32.zip dynamock.exe

windows_64:
	@echo "Building for windows 64bits..."
	@GOOS=windows GOARCH=amd64 go build && zip dynamock_windows_64.zip dynamock.exe

darwin_osx:
	@echo "Building for Darwin/OSX"
	@GOOS=darwin go build && zip dynamock_darwin_osx.zip dynamock
