test:
	go test

run:
	go run main.go

build:
	go build -o bin/$(APP_NAME) main.go
