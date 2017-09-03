all: test

.PHONY: test test_local govendor

test:
	go vet ./...
	go test -v -cover ./...

test_local:
	go run ./cmd/wotreplay-parse/main.go _replays/*.wotreplay

govendor: $(GOPATH)/bin/govendor
	govendor sync
	govendor list

$(GOPATH)/bin/govendor:
	go get -u github.com/kardianos/govendor
