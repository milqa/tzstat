PROJECT_NAME := "tzstat"
PKG := "github.com/milQA/$(PROJECT_NAME)"
CMD := "$(PKG)/cmd/$(PROJECT_NAME)"
BIN := "bin/$(PROJECT_NAME)"
PIDFILE=$(BIN).pid

PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

GIT_COMMIT = $(shell git log --pretty=format:'%h' -n 1)
VERSION = $(shell date +'%Y%m%d%H%M.${GIT_COMMIT}')

LDFLAGS_VERSION := "main.Version"
LDFLAGS_APPLICATION := "main.Application"
LDFLAGS := "-X $(LDFLAGS_VERSION)=$(VERSION) -X $(LDFLAGS_APPLICATION)=$(PROJECT_NAME)"

.PHONY: all dep build test coverage coverhtml lint

all: build

## TODO need add install golint
lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: dep ## Run unittests
	@go test -short ${PKG_LIST}

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

coverage: ## Generate global code coverage report
	shell ./tools/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	shell ./tools/coverage.sh html;

dep: ## Get the dependencies
	@go get -v -d ./...

build: dep ## Build the binary file
	@go build -ldflags $(LDFLAGS) -v -o $(BIN) $(CMD)

clean: ## Remove previous build
	rm -f $(BIN)
	find ./ -type d -name _mocks -exec rm -rf '{}' \;
	find ./ -type f -name "*.out" -delete
	find ./ -type f -name "coverage.html" -delete

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
