## which linter version to use?
GOLANGCI_LINT_VERSION = v1.31.0

GO = go
TOOLSDIR = $(TMPDIR)/wdy-go-byteorder
GOLANGCI_LINT = ${TOOLSDIR}/golangci-lint

all: generate lint test ## Runs lint and test

lint: ## Executes all linters
	${GOLANGCI_LINT} run --enable-all --exclude-use-default=false

test: ## Executes the tests
	${GO} test -race ./...

.PHONY: lint test setup

help: ## Shows this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

generate: ## Executes go generate
	${GO} generate ./...

setup: installGolangCi  ## Installs golangci-lint
	${GO} mod tidy


installGolangCi:
	mkdir -p ${TOOLSDIR}
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TOOLSDIR) $(GOLANGCI_LINT_VERSION)


.DEFAULT_GOAL := all
