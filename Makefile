export PATH := $(CURDIR)/.bin:$(PATH)

TARGETS = mission-reward-server

GOLANGCI_LINT = golangci-lint run
TEST = ./...
PKGNAME = $(shell go list -m)
GIT_COMMIT = $(shell git rev-parse HEAD)
VERSION =$(shell git symbolic-ref -q --short HEAD || git describe --tags --exact-match)
BUILD=local
LDFLAGS = -ldflags "-X $(PKGNAME)/pkg/version.gitCommit=$(GIT_COMMIT) \
										-X $(PKGNAME)/pkg/version.version=$(VERSION)\
										-X $(PKGNAME)/pkg/version.build=$(BUILD)"
DB_ADDRESS := localhost
GOOSE_OPTION = -dir ./db/migrate mysql "user:password@tcp(${DB_ADDRESS}:3306)/mission_reward?parseTime=true&loc=Asia%2FTokyo"

# command
defualt: tools help

## Install dependency and tools
setup: deps tools

## Install dependency
deps:
	go get ./...
deps.update.minor:
	go get -t -u ./...
deps.update.patch:
	go get -t -u=patch ./...
deps.tidy:
	go mod tidy

## Install tools
tools:
	export GOBIN=$(CURDIR)/.bin &&\
	go install github.com/Songmu/make2help/cmd/make2help@v0.2.1 &&\
	go install github.com/kyoh86/richgo@v0.3.12 &&\
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3 &&\
	go install github.com/pressly/goose/v3/cmd/goose@latest &&\
	go install github.com/cosmtrek/air@latest &&\
	go install github.com/go-delve/delve/cmd/dlv@latest &&\
	go install google.golang.org/protobuf/cmd/protoc-gen-go



## Remove build target
clean:
	rm -f $(TARGETS)
	rm -rf dist
	rm -rf tmp

## Build app
build: clean deps
	go build $(LDFLAGS) -o dist/$(TARGETS) ./cmd/server/main.go

## run service app
run-server:
	air

## run test client
run-client:
	go run ./cmd/client/main.go


## Check code format
check:
	$(GOLANGCI_LINT) ./...

## Fix code
fix:
	$(GOLANGCI_LINT) --fix ./...

## Run test
test: tools
	mkdir -p tmp
	richgo test -race -coverprofile=tmp/coverage.txt -covermode=atomic $(TEST)

## generate protoc file
protoc:
	protoc -I./proto --go_out=./pkg/grpc --go_opt=paths=source_relative \
          --go-grpc_out=./pkg/grpc --go-grpc_opt=paths=source_relative \
          ./proto/*.proto

## Run migrate help
migrate-help:
	goose -h $(GOOSE_OPTION)

## Run migrate
migrate-%:
	goose  $(GOOSE_OPTION) ${@:migrate-%=%}

## ssh service container
ssh-server:
	docker compose exec server /bin/bash

## Show help
help:
	@make2help $(MAKEFILE_LIST)

NO_PHONY = /^:/
PHONY := $(shell cat $(MAKEFILE_LIST) | awk -F':' '/^[a-z0-9_.-]+:/ && !$(NO_PHONY) {print $$1}')
.PHONY: $(PHONY)

show_phony:
	@echo $(PHONY)
