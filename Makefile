all:
	go build "./..."

test:
	.githooks/pre-commit

run:
	go run main.go

clean:
	go clean "./..."
