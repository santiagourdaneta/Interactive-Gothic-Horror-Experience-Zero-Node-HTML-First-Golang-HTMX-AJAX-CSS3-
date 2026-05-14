run:
	go run cmd/server/main.go

test:
	go test ./... -v

lint:
	# Instala: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

build:
	go build -o server cmd/server/main.go