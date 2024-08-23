build:
		@go build -o bin/fetch

run: build
		@./bin/fetch

test:
		@go test -b ./...