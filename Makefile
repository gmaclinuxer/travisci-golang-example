# global go tools
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=gofmt

# define build version
VERSION := 1.0.0
BUILD_COMMIT := `git rev-parse HEAD`
BUILD_TIME := `date`
GO_VERSION := `go version`

# go build flags
LDFLAGS = -ldflags "-X main.Version=$(VERSION) -X 'main.BuildCommit=$(BUILD_COMMIT)' \
-X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GoVersion=$(GO_VERSION)'"

# binary name
TARGET=main
TARGET_DIR=./src/sorter
TARGET_SRC=$(shell find $(TARGET_DIR) -type f -name "*.go" -not -path "./pkg/*")
GOFILES=$(shell find . -type f -name "*.go" -not -path "./pkg/*")

all: test build

.PHONY: clean

$(TARGET): $(TARGET_SRC)
	$(GOBUILD) $(LDFLAGS) -o $(TARGET) $(TARGET_SRC)
	cp $(TARGET) ./bin

build: $(TARGET)
	@true

test:
	$(GOTEST) -v -cover=true ./...

fmt:
	$(GOFMT) -l -w $(GOFILES)

clean:
	$(GOCLEAN)
	rm -f $(TARGET) bin/$(TARGET)

check:
	@test -z $(shell gofmt -l $(GOFILES) | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${GOFILES}
