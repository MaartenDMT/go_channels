build: 
	go build -o ./bin/channels

run: build 
	./bin/channels

test:
	go test -v ./...
	