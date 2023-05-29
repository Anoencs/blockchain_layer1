build: 
	@ go build -o bin/blockchain_layer1
run: build
	@ ./bin/blockchain_layer1
test:
	@ go test ./...
	
