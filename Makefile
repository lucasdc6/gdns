.PHONY: build run test clean dependencies dep

all: test build

build:
	bash scripts/build.sh

dependencies:
	@echo "Creating \`vendor\` directory"
	@go mod vendor

dep: dependencies

run:
	@go run cmd/gdns/main.go

test:
	@echo "Executing tests..."
	@go test ./...

clean:
	@echo "Clean environment"
	@-go clean
	@rm -r bin
