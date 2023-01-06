build:
	go build -o bin/main src/*.go

build-linux:
	GOOS=linux go build -o bin/main src/*.go

build-windows:
	GOOS=windows go build -o bin/main src/*.go

run:
	go run src/*.go

all: build