.PHONY: build test run clean

build:
	go build -o bin/loctree cmd/loctree/main.go

test:
	go test ./...

run:
	go run cmd/loctree/main.go

clean:
	rm -rf bin/