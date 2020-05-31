all: clean build test image

build:
	go build "./..."

test:
	go fmt ./...
	go vet ./...
	golint "-set_exit_status" ./...
	go test ./...
    
run:
	go run cmd/server/main.go

image:
	docker build . --tag quick-calc-service:latest

clean:
	go clean "./..."
