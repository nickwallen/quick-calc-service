all: build test image

build:
	go build "./..."

test:
	.githooks/pre-commit

run:
	go run cmd/server/main.go

image:
	docker build . --tag quick-calc-service:latest

clean:
	go clean "./..."
