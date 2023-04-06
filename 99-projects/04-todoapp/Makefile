# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

# Binary name
BINARY_NAME=./build/todoapp

all: clean build

build:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) ./cmd/todoapp
	ls $(BINARY_NAME)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) ./cmd/todoapp
	./$(BINARY_NAME)
