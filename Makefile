.PHONY: dev build run test lint clean

dev:
	go mod tidy
	go run ./cmd/filemanager

build:
	@mkdir -p tmp/bin
	go build -o tmp/bin/filemanager ./cmd/filemanager

run:
	go run ./cmd/filemanager

lint:
	golangci-lint run

clean:
	rm -rf tmp/

test:
	go test ./... -v

