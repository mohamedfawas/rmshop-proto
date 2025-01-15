.PHONY: generate
generate:
	buf generate

.PHONY: lint
lint:
	buf lint

.PHONY: clean
clean:
	rm -rf gen/*

.PHONY: all
all: clean generate

.PHONY: init
init:
	go mod init github.com/mohamedfawas/rmshop-proto
	go mod tidy