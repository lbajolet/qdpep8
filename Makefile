.PHONY: all

all: bin
	go build -o bin/qdpep8_cli ./qdpep8cli

bin:
	mkdir bin
