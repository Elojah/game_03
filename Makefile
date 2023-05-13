PACKAGE            = game_03
DATE              ?= $(shell date +%FT%T%z)
VERSION           ?= $(shell echo $(shell cat $(PWD)/.version)-$(shell git describe --tags --always))
DIR                = $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

GO                 = go
GOROOT             ?= $(shell go env GOROOT)
GODOC              = godoc
GOFMT              = gofmt

# For CI
ifneq ($(wildcard ./bin/golangci-lint),)
	GOLINT = ./bin/golangci-lint
else
	GOLINT = golangci-lint
endif

V                 = 0
Q                 = $(if $(filter 1,$V),,@)
M                 = $(shell printf "\033[0;35m▶\033[0m")

GO_PACKAGE        = github.com/elojah/game_03
API               = api
ADMIN             = admin
AUTH              = auth
CORE              = core
WEB_CLIENT        = web_client
CLIENT            = client
DASHBOARD         = dashboard
WEB_DASHBOARD     = web_dashboard

# Static directory name for client
STATIC            = static

GEN_PARENT_PATH    = $(GOPATH)/src
PROTOC_GEN_TS      = $(DIR)/cmd/$(CLIENT)/node_modules/.bin/protoc-gen-ts
GEN_PB_GO          = protoc -I=$(GEN_PARENT_PATH) --gogoslick_out=$(GOPATH)/src
GEN_PB_TS          = protoc -I=$(GEN_PARENT_PATH) --plugin=protoc-gen-ts=$(PROTOC_GEN_TS) --js_out=import_style=commonjs,binary:$(GOPATH)/src --ts_out="service=grpc-web:$(GOPATH)/src"
GEN_PB_SERVICE_GO  = protoc -I=$(GEN_PARENT_PATH) -I=$(GOPATH)/src --gogoslick_out=plugins=grpc,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:$(GOPATH)/src
GEN_PB_SERVICE_TS  = protoc -I=$(GEN_PARENT_PATH) -I=$(GOPATH)/src --plugin=protoc-gen-ts=$(PROTOC_GEN_TS) --js_out=import_style=client,binary:$(GOPATH)/src --ts_out=service=grpc-web:$(GOPATH)/src

.PHONY: all
all: api admin auth core client web_client dashboard web_dashboard

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

.PHONY: core
core:  ## Build core binary
	$(info $(M) building executable core…) @
	$Q cd cmd/$(CORE) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(CORE)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(CORE)_$(VERSION) bin/$(PACKAGE)_$(CORE)

.PHONY: web_client
web_client: ## Build web_client binary
	$(info $(M) building executable web_client…) @
	$Q cd cmd/$(WEB_CLIENT) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(WEB_CLIENT)_$(VERSION)
	$Q yes | cp -rf bin/$(PACKAGE)_$(WEB_CLIENT)_$(VERSION) bin/$(PACKAGE)_$(WEB_CLIENT)

.PHONY: web_dashboard
web_dashboard: ## Build web_dashboard binary
	$(info $(M) building executable web_dashboard…) @
	$Q cd cmd/$(WEB_DASHBOARD) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(WEB_DASHBOARD)_$(VERSION)
	$Q yes | cp -rf bin/$(PACKAGE)_$(WEB_DASHBOARD)_$(VERSION) bin/$(PACKAGE)_$(WEB_DASHBOARD)

.PHONY: client
client:  ## Build client content
	$(info $(M) building bundle client…) @
	$Q cd cmd/$(CLIENT) && npx webpack --config webpack.config.js
	$Q mkdir -p bin && rm -rf bin/$(STATIC) && mkdir -p bin/$(CLIENT)/$(STATIC)/
	$Q yes | cp -rf cmd/$(CLIENT)/dist/. bin/$(CLIENT)/$(STATIC)/

.PHONY: dashboard
dashboard:  ## Build dashboard content
	$(info $(M) building bundle dashboard…) @
	$Q cd cmd/$(DASHBOARD) && npx webpack --config webpack.config.js
	$Q mkdir -p bin && rm -rf bin/$(STATIC) && mkdir -p bin/$(DASHBOARD)/$(STATIC)/
	$Q yes | cp -rf cmd/$(DASHBOARD)/dist/. bin/$(DASHBOARD)/$(STATIC)/

.PHONY: populate
populate:  ## Populate initial content
	$(info $(M) populate initial content…) @
	$Q ./scripts/upload_default_images.sh
	$Q ./scripts/create_default_templates.sh
	$Q ./scripts/create_default_tilesets.sh
	$Q ./scripts/create_default_animations.sh

.PHONY: regen-assets
regen-assets:  ## Regenerate assets from external raw
	$(info $(M) regenerate assets …) @
	$Q go run ./scripts/write_animation/main.go 'assets/animations'
	$Q go run ./scripts/write_tileset/main.go 'assets/external/Tilesets' 'assets/tilesets'

# Proto lang
.PHONY: proto-go proto-ts
proto-go:    PB_LANG = GO
proto-ts:    PB_LANG = TS
proto-go proto-ts: ## Regenerate protobuf files
	$(info $(M) running protobuf $(PB_LANG)…) @
	$(info $(M) generate utils…) @
	$Q $(GEN_PB_$(PB_LANG)) github.com/gogo/protobuf/gogoproto/gogo.proto
	$(info $(M) generate domain…) @
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/ability/ability.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/ability/cast.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/cookie/keys.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/ability.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/animation.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/entity.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/pc.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/spawn.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/template.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/event/event.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/faction/faction.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/geometry/geometry.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/cell.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/room.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/user.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/world.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/spawn.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/twitch/follow.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/twitch/user.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/user/user.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/user/role.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/user/session.proto
	$(info $(M) generate clients…) @
	$(info $(M) generate dto…) @
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/ability/dto/ability.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/dto/animation.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/dto/entity.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/dto/pc.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/entity/dto/template.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/dto/cell.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/dto/room.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/dto/user.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/room/dto/world.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/rtc/dto/rtc.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/tile/dto/set.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/twitch/dto/follow.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/user/dto/session.proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/user/dto/user.proto
	$(info $(M) generate services…) @
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(ADMIN)/grpc/$(ADMIN).proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(API)/grpc/$(API).proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(AUTH)/grpc/$(AUTH).proto
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(CORE)/grpc/$(CORE).proto
	$Q $(GEN_PB_SERVICE_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(ADMIN)/grpc/$(ADMIN).proto
	$Q $(GEN_PB_SERVICE_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(API)/grpc/$(API).proto
	$Q $(GEN_PB_SERVICE_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(AUTH)/grpc/$(AUTH).proto
	$Q $(GEN_PB_SERVICE_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(CORE)/grpc/$(CORE).proto

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
	$(info $(M) cleaning binaries…) @
	$Q rm -rf bin/$(PACKAGE)_$(API)*
	$Q rm -rf bin/$(PACKAGE)_$(ADMIN)*
	$Q rm -rf bin/$(PACKAGE)_$(AUTH)*
	$Q rm -rf bin/$(PACKAGE)_$(CORE)*
	$Q rm -rf bin/$(PACKAGE)_$(WEB_CLIENT)*
	$Q rm -rf bin/$(PACKAGE)_$(CLIENT)*
	$Q rm -rf bin/$(PACKAGE)_$(WEB_DASHBOARD)*
	$Q rm -rf bin/$(PACKAGE)_$(DASHBOARD)*
	$Q rm -rf bin/$(CLIENT)/$(STATIC)
	$Q rm -rf bin/$(DASHBOARD)/$(STATIC)

# Clean
.PHONY: clean-assets
clean-assets:
	$(info $(M) cleaning assets…) @
	$Q rm -rf cmd/$(CLIENT)/dist/img/assets/*
	$Q rm -rf cmd/$(CLIENT)/dist/json/assets/*
	$Q rm -rf cmd/$(DASHBOARD)/dist/img/assets/*

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
