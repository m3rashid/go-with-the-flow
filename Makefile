test:
	go test ./... -cover

run:
	go run main.go

build:
	rm -rf bin/main
	go build -o bin/main main.go
