# global go tools
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# go build flags
GOFLAGS = -x -v

# binary name
TARGET_NAME=main
TARGET_ENTRY=src/sorter/main.go

all: test build

build:
	$(GOBUILD) -o bin/$(TARGET_NAME) $(GOFLAGS) $(TARGET_ENTRY)

test:
	$(GOTEST) -v ./...

.PHONY: clean

clean:
	$(GOCLEAN)
	rm -f $(TARGET_NAME)

