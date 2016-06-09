default: clean build clean_build

clean:
	@echo "Clean all..."
	@rm dynamock_*.zip || true
	@rm dynamock || true

build: linux_64 linux_32 windows_32 windows_64 darwin_osx
	@echo "Complete"

clean_build:
	@rm dynamock || true
	@rm dynamock.exe || true

linux_64: clean_build
	@echo "Building for linux 64bits..."
	@GOOS=linux GOARCH=amd64 go build && zip dynamock_linux_64.zip dynamock

linux_32: clean_build
	@echo "Building for linux 32bits..."
	@GOOS=linux GOARCH=386 go build && zip dynamock_linux_32.zip dynamock

windows_32: clean_build
	@echo "Building for windows 32bits..."
	@GOOS=windows GOARCH=386 go build && zip dynamock_windows_32.zip dynamock.exe

windows_64: clean_build
	@echo "Building for windows 64bits..."
	@GOOS=windows GOARCH=amd64 go build && zip dynamock_windows_64.zip dynamock.exe

darwin_osx: clean_build
	@echo "Building for Darwin/OSX"
	@GOOS=darwin go build && zip dynamock_darwin_osx.zip dynamock
