all:
	go build "./..."

test:
	.githooks/pre-commit

run:
	go run cmd/server/main.go

clean:
	go clean "./..."
