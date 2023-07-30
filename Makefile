.PHONY: server test

server:
	@go run cmd/main.go

test:
	@go test ./...