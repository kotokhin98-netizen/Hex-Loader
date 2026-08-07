.PHONY: build run test clean

build:
	go build -o hex-loader cmd/hex-loader/main.go

run:
	go run cmd/hex-loader/main.go

test:
	go test ./...

clean:
	rm -f hex-loader