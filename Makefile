PACKAGE    = game_03
DATE      ?= $(shell date +%FT%T%z)
VERSION   ?= $(shell echo $(shell cat $(PWD)/.version)-$(shell git describe --tags --always))
DIR        = $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

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
ADMIN             = admin
AUTH              = auth
WEB               = web
BROWSER           = browser
DASHBOARD         = dashboard

# Static directory name for browser
STATIC            = static

GEN_PARENT_PATH    = $(GOPATH)/src
PROTOC_GEN_TS      = $(DIR)/cmd/$(BROWSER)/node_modules/.bin/protoc-gen-ts
GEN_PB_GO          = protoc -I=$(GEN_PARENT_PATH) --gogoslick_out=$(GOPATH)/src
GEN_PB_TS          = protoc -I=$(GEN_PARENT_PATH) --plugin=protoc-gen-ts=$(PROTOC_GEN_TS) --js_out=import_style=commonjs,binary:$(GOPATH)/src --ts_out="service=grpc-web:$(GOPATH)/src"
GEN_PB_SERVICE_GO  = protoc -I=$(GEN_PARENT_PATH) -I=$(GOPATH)/src --gogoslick_out=plugins=grpc,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:$(GOPATH)/src
GEN_PB_SERVICE_TS  = protoc -I=$(GEN_PARENT_PATH) -I=$(GOPATH)/src --plugin=protoc-gen-ts=$(PROTOC_GEN_TS) --js_out=import_style=browser,binary:$(GOPATH)/src --ts_out=service=grpc-web:$(GOPATH)/src

.PHONY: all
all: api admin auth web

.PHONY: api
api:  ## Build api binary
	$(info $(M) building executable api…) @
	$Q cd cmd/$(API) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(API)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(API)_$(VERSION) bin/$(PACKAGE)_$(API)

.PHONY: admin
admin:  ## Build admin binary
	$(info $(M) building executable admin…) @
	$Q cd cmd/$(ADMIN) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(ADMIN)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(ADMIN)_$(VERSION) bin/$(PACKAGE)_$(ADMIN)

.PHONY: auth
auth:  ## Build auth binary
	$(info $(M) building executable auth…) @
	$Q cd cmd/$(AUTH) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(AUTH)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(AUTH)_$(VERSION) bin/$(PACKAGE)_$(AUTH)

.PHONY: web
web: ## Build web binary
	$(info $(M) building executable web…) @
	$Q cd cmd/$(WEB) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(WEB)_$(VERSION)
	$Q yes | cp -rf bin/$(PACKAGE)_$(WEB)_$(VERSION) bin/$(PACKAGE)_$(WEB)

.PHONY: browser
browser:  ## Build browser content
	$(info $(M) building bundle browser…) @
	$Q cd cmd/$(BROWSER) && npx webpack --config webpack.config.js
	$Q mkdir -p bin && rm -rf bin/$(STATIC) && mkdir -p bin/$(STATIC)
	$Q yes | cp -rf cmd/$(BROWSER)/dist/. bin/$(STATIC)/

.PHONY: dashboard
dashboard:  ## Build dashboard content
	$(info $(M) building bundle dashboard…) @
	$Q cd cmd/$(DASHBOARD) && npx webpack --config webpack.config.js
	$Q mkdir -p bin && rm -rf bin/$(STATIC) && mkdir -p bin/$(STATIC)
	$Q yes | cp -rf cmd/$(DASHBOARD)/dist/. bin/$(STATIC)/

# Proto lang
.PHONY: proto-go proto-ts
proto-go:    PB_LANG = GO
proto-ts:    PB_LANG = TS
proto-go proto-ts: ## Regenerate protobuf files
	$(info $(M) running protobuf $(PB_LANG)…) @
	$(info $(M) generate utils…) @
	$Q $(GEN_PB_$(PB_LANG)) github.com/gogo/protobuf/gogoproto/gogo.proto
	$(info $(M) generate domain…) @
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/animation.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/entity.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/npc.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/pc.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/template.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/geometry/geometry.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/cell.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/room.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/user.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/world.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/twitch/follow.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/twitch/user.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/user/user.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/user/role.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/user/session.proto
	$(info $(M) generate clients…) @
	$(info $(M) generate dto…) @
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/dto/animation.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/dto/entity.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/dto/pc.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/dto/template.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/dto/cell.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/dto/room.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/dto/world.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/twitch/dto/follow.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/tile/dto/set.proto
	$(info $(M) generate services…) @
	$Q $(GEN_PB_SERVICE_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(API)/grpc/$(API).proto
	$Q $(GEN_PB_SERVICE_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(ADMIN)/grpc/$(ADMIN).proto
	$Q $(GEN_PB_SERVICE_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(AUTH)/grpc/$(AUTH).proto

# Proto
.PHONY: proto
proto: proto-go proto-ts

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

# Clean
.PHONY: clean
clean:
	$(info $(M) cleaning bin…) @
	$Q rm -rf bin/$(PACKAGE)_$(API)_*
	$Q rm -rf bin/$(PACKAGE)_$(ADMIN)_*
	$Q rm -rf bin/$(PACKAGE)_$(AUTH)_*
	$Q rm -rf bin/$(PACKAGE)_$(WEB)_*
	$Q rm -rf bin/$(STATIC)

## Helpers

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
