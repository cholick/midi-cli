SRC=$(shell go list ./... | grep -v fake)

.PHONY: test

default: fmt generate vet test build

build:
	go build ./cmd/midi-cli

test:
	go test $(SRC)

generate:
	COUNTERFEITER_NO_GENERATE_WARNING=true go generate ./...

vet:
	go vet $(SRC)

fmt:
	gofmt -l -w -s .
	go tool goimports -l -w .
