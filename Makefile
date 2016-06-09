ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

test:
	GOPATH=$(ROOT_DIR) go test ./...
