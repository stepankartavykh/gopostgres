build:
	@go build -o bin/gopostgres

run: build
	@./bin/gopostgres

test:
	@go test ./... -v

run-several-clients-test:
	@go test -v mypostgres/test/highload -run TestConnectSeveralClientsToDatabase