.PHONY: deps clean build deploy

deps:
	go get -u ./...

clean:
	rm -rf bin

build:
	GOOS=linux go build -o bin/main

deploy: clean build
	yarn serverless deploy
