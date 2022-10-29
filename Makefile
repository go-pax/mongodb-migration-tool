SHELL := /bin/bash

clean:
	rm -rf dist/

# ==============================================================================
# Building binary on MAC
# ==============================================================================

BINARY_NAME := mongodb_migration

build: clean
	go build -o dist/$(BINARY_NAME) -gcflags all=-N
