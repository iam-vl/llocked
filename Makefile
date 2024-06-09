.DEFAULT_GOAL := build

.PHONY:fmt vet build run clean
fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

build: vet 
	go build -o app 

run: build 
	./app 

clean: run 
	rm ./app