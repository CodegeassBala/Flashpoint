APP_NAME=flashpoint
CMD_PATH=./cmd/main.go

build:
	go build -o $(APP_NAME) $(CMD_PATH)

run: build
	 ./$(APP_NAME)

dev: 
	go run $(CMD_PATH)

test:
	go test ./...

test-cover:
	go test -cover ./...