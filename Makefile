include ../.env
export

all: run

run:
	go run cmd/main.go

test:
	go test ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

