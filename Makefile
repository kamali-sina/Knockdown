all: build

build:
	go build -o knockdown.out

run: build
	./knockdown.out