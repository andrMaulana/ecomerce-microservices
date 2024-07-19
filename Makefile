default: build test

build:
    go build -o cmd/main cmd/main.go

test:
    go test ./...

deploy:
    docker-compose up -d