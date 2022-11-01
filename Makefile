build:
	go build -o bin/channels

run: build
	go run ./bin/channels

test:
	go test -v  ./... --race