build:
	@go build -o bin/gopostgres

run: build
	@./bin/gopostgres

test:
	@go test ./... -v
