build: 
	@go build -o bin/go-expense-tracker-api cmd/main.go

run: build
	@./bin/go-expense-tracker-api

test:
	@go test -v ./..