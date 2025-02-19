.PHONY: all get build buildmac

get:
	go get -v ./...

build:
	GOARCH=amd64 GOOS=linux go build -o bin/shortener-linux-amd64 cmd/shortener/*.go

buildmac:
	GOARCH=arm64 GOOS=darwin go build -o bin/shortener-darwin-arm64 cmd/shortener/*.go

builddocker:
	docker build -t shortener .

rundocker:
	docker run -p 8080:8080 shortener

all: get build
