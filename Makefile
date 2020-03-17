TARGET  := bin/wotreplay-parse
VERSION := v0.1-$(shell git rev-parse --short HEAD)
LDFLAGS := -ldflags "-X main.Version=$(VERSION)" #-ldflags "-w -s"

.PHONY: run test test_local build govendor clean

default: test $(TARGET)

run: build
#	go run ./cmd/wotreplay-parse/main.go
	./$(TARGET)

build: $(TARGET)

$(TARGET):
	go build -gcflags "-N -l" $(LDFLAGS) -o $@ ./cmd/wotreplay-parse/main.go

test:
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
