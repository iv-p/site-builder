.PHONY: test build init

package_main = services/renderer/main.go
executable = site-renderer
all: build

build: test
	@echo "Building"
	@go build -o ./bin/$(executable) $(package_main)

test: init
	@echo "Running unit tests"
	@go test ./...

init:
	@echo "Ensuring dependencies"
	@go get

run: build
	./bin/$(executable)