run:
	go run ./cmd/forgebase

test:
	go test ./...

build:
	go build -o forgebase ./cmd/forgebase

docker:
	docker build -t forgebase .
