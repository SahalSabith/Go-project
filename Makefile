build:
	@go build -o bin/GoProject

run: build
	@./bin/GoProject

test:
	@go test -v ./...