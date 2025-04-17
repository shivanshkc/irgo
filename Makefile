SHELL=/usr/bin/env bash

# Builds the project.
build:
	@echo "+$@"
	@go build -o bin/irgo cmd/irgo/main.go

# Runs the project after linting and building it anew.
run: tidy build
	@echo "+$@"
	@echo "########### Running the application binary ############"
	@bin/irgo nums/phi-1m.txt

# Tests the whole project.
test:
	@echo "+$@"
	@CGO_ENABLED=1 go test -race -coverprofile=coverage.out -covermode=atomic ./...

# Runs the "go mod tidy" command.
tidy:
	@echo "+$@"
	@go mod tidy

# Runs golang-ci-lint over the project.
lint:
	@echo "+$@"
	@golangci-lint run ./...
