.PHONY: test build init

all: build

build: test
	@echo "Building"
	@cd src/github.com/iv-p/ssb; go build -o ../../../../bin/ssb main.go

test: init
	@echo "Running unit tests"
	@cd src/github.com/iv-p/ssb; go test ./...

init:
	@echo "Ensuring dependencies"
	@cd src/github.com/iv-p/ssb; go get

run: build
	./bin/ssb