GO := $(HOME)/go/bin/go1.20.5

build:
	$(GO) build -o ./bin/app

run:build
	./bin/app

tidy:
	$(GO) mod tidy

format:
	$(GO) fmt ./...
