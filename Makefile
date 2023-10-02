build:
	@go build -o bin/go-bank-jwt

run: build
	@./bin/go-bank-jwt

test:
	@go test -v ./...
