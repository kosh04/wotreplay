TARGET   := bin/wotreplay-parse
VERSION  := $(shell git describe --tag)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS  := -ldflags "-X main.version=$(VERSION) -X main.revision=$(REVISION)"
#LDFLAGS += -ldflags "-w -s"

.PHONY: run test test_local build govendor clean

default: test $(TARGET)

run: build
	./$(TARGET) -version

build: $(TARGET)

$(TARGET):
	go build -gcflags "-N -l" $(LDFLAGS) -o $@ ./cmd/wotreplay-parse

test: build
	go vet ./...
	go test -v -cover ./...

test_local: build
	./$(TARGET) testdata/replays/*.wotreplay

govendor: $(GOPATH)/bin/govendor
	govendor sync
	govendor list

$(GOPATH)/bin/govendor:
	go get -u github.com/kardianos/govendor

clean:
	go clean
	$(RM) $(TARGET)
