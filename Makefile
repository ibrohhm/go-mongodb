run:
	go run main.go

mod:
	go mod download

tidy:
	go mod tidy

test:
	go test ./...