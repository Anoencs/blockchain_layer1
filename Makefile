build:
	go build -o ./bin/projectx

run: build
	sudo ./bin/projectx

test:
	go test ./...
