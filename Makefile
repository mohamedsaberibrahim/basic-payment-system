BINARY_NAME := payment-api

build:
	go build -o bin/$(BINARY_NAME) -v

run:
	go run main.go

test:
	go test -v ./...

clean:
	go clean
	rm -f bin/$(BINARY_NAME)
