PACKAGE    = game_03
DATE      ?= $(shell date +%FT%T%z)
VERSION   ?= $(shell echo $(shell cat $(PWD)/.version)-$(shell git describe --tags --always))

GO         = go
GOROOT     ?= $(shell go env GOROOT)
GODOC      = godoc
GOFMT      = gofmt

# For CI
ifneq ($(wildcard ./bin/golangci-lint),)
	GOLINT = ./bin/golangci-lint
else
	GOLINT = golangci-lint
endif

V          = 0
Q          = $(if $(filter 1,$V),,@)
M          = $(shell printf "\033[0;35m▶\033[0m")

GO_PACKAGE        = github.com/elojah/game_03
API               = api
WEB               = web

GEN_PB            = protoc -I=$(GOPATH)/src --gogoslick_out=$(GOPATH)/src
GEN_PB_SERVICE    = protoc -I=$(GOPATH)/src --gogoslick_out=plugins=grpc,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:$(GOPATH)/src

.PHONY: all
all: api

.PHONY: api
api:  ## Build api binary
	$(info $(M) building executable api…) @
	$Q cd cmd/$(API) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(API)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(API)_$(VERSION) bin/$(PACKAGE)_$(API)

.PHONY: web
web:  ## Build web binary
	$(info $(M) building executable web…) @
	$Q cd cmd/$(WEB) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(WEB)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(WEB)_$(VERSION) bin/$(PACKAGE)_$(WEB)

# Utils
.PHONY: proto
proto: ## Regenerate protobuf files
	$(info $(M) running protobuf…) @
	$(info $(M) generate utils…) @
	$(info $(M) generate domain…) @
	$(info $(M) generate clients…) @
	$(info $(M) generate dto…) @
	$(info $(M) generate services…) @
	$Q $(GEN_PB_SERVICE) $(GO_PACKAGE)/cmd/$(API)/grpc/$(API).proto

# Vendor
.PHONY: vendor
vendor:
	$(info $(M) running go mod vendor…) @
	$Q $(GO) mod vendor

# Tidy
.PHONY: tidy
tidy:
	$(info $(M) running go mod tidy…) @
	$Q $(GO) mod tidy

# Check
.PHONY: check
check: vendor test lint

# Lint
.PHONY: lint
lint:
	$(info $(M) running $(GOLINT)…)
	$Q $(GOLINT) run

# Test
.PHONY: test
test:
	$(info $(M) running go test…) @
	$Q $(GO) test -cover -race -v ./...

# Helpers
.PHONY: go-version
go-version: ## Print go version used in this makefile
	$Q echo $(GO)

.PHONY: fmt
fmt: ## Format code
	$(info $(M) running $(GOFMT)…) @
	$Q $(GOFMT) ./...

.PHONY: doc
doc: ## Generate project documentation
	$(info $(M) running $(GODOC)…) @
	$Q $(GODOC) ./...

.PHONY: version
version: ## Print current project version
	@echo $(VERSION)

.PHONY: help
help: ## Print this
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
