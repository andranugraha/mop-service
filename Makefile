build:
	go build -o bin/mop-service cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/*

.PHONY: build run test clean