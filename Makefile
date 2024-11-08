SWAG_BIN := $(shell go env GOPATH)/bin/swag

build:
	go mod tidy
	make docs
	go build -o bin/mop-service cmd/server/main.go

run:
	./bin/mop-service

test:
	go test -v ./...

clean:
	rm -rf bin/*

lint:
	gofumpt -l -w .

docs:
	go install github.com/swaggo/swag/cmd/swag@latest
	$(SWAG_BIN) init -g cmd/server/main.go

.PHONY: build run test clean lint docs