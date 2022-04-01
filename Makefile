
proto:
	protoc --go_out=paths=source_relative:. ./internal/proto/click.proto
run:
	go run cmd/main.go
build:
	go build -o main ./cmd/main.go
