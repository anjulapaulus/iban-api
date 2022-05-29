BINARY_NAME=iban-api

WINDOWS=$(BINARY_NAME)-windows-amd64.exe
LINUX=$(BINARY_NAME)-linux-amd64
DARWIN=$(BINARY_NAME)-darwin-amd64

## to build binaries for each platform
## Build for Windows
windows:  
	GOARCH=amd64 go build -o $(WINDOWS) main.go
## Build for Linux
linux: 
	GOARCH=amd64 go build -o $(LINUX) main.go
## Build for Darwin (macOS)
darwin:
	GOARCH=amd64 go build -o $(DARWIN) main.go


## OS env is only in windows
ifeq ($(OS),Windows_NT)
    os := Windows
else
    os := $(shell uname -s)
endif

## Build for all platforms specified
build: 
ifeq ($(os), Windows)
    env GOOS=windows GOARCH=amd64 go build -o $(WINDOWS) main.go
endif
ifeq ($(os), Linux)
	env GOOS=linux GOARCH=amd64 go build -o $(LINUX) main.go
endif
ifeq ($(os), Darwin)
    env GOOS=darwin GOARCH=amd64 go build -o $(DARWIN) main.go
endif


run:
ifeq ($(os), Windows)
   ./$(WINDOWS)
endif
ifeq ($(os), Linux)
	./$(LINUX)
endif
ifeq ($(os), Darwin)
    ./$(DARWIN)
endif

test:
	go test -coverpkg=./... -coverprofile coverProfile.out ./...
	go tool cover -func coverProfile.out | grep total | awk '{print $3}'
	rm coverProfile.out

build_and_run: build run

clean:
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)
	go clean

docker_run:
	docker-compose up -d

docker_rm:
	docker stop iban-api
	docker rm iban-api

.PHONY: all test clean

all: test build run

