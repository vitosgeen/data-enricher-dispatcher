include ./.env
build:
	go build ./main.go

run:
	go run ./main.go

test:
	go test -v -cover ./...

run-linter:
	echo "Starting linters"
	golangci-lint run ./...
