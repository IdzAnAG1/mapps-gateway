#######################################################################################################################
#                                    Environment                                                                      #
#######################################################################################################################

GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
REPO=https://github.com/IdzAnAG1/mapps-contracts.git#branch=main
BUF_GEN=buf generate

#######################################################################################################################
#                              Proto files discovery (OS)                                                             #
#######################################################################################################################

ifeq ($(GOHOSTOS), windows)
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

#######################################################################################################################
#                                    Targets                                                                          #
#######################################################################################################################

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

#######################################################################################################################
#                              Protobuf: internal config                                                              #
#######################################################################################################################

.PHONY: config
# generate internal proto config
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

#######################################################################################################################
#                                 Protobuf: API                                                                       #
#######################################################################################################################

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
	       --go_out=paths=source_relative:./api \
	       --go-http_out=paths=source_relative:./api \
	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

#######################################################################################################################
#                                    Build                                                                            #
#######################################################################################################################

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

#######################################################################################################################
#                           Go generate / dependencies                                                                #
#######################################################################################################################

.PHONY: generate
# generate
generate:
	go generate ./...
	go mod tidy

#######################################################################################################################
#                                      Wire                                                                           #
#######################################################################################################################

.PHONY: wire
# regenerate wire_gen.go
wire:
	cd cmd/mApps_gateway && wire

#######################################################################################################################
#                                      Local run                                                                      #
#######################################################################################################################

.PHONY: local
# run locally
local:
	go run ./cmd/mApps_gateway/... -conf ./configs

#######################################################################################################################
#                                      All                                                                            #
#######################################################################################################################

.PHONY: all
# generate all
all:
	make api
	make config
	make generate

#######################################################################################################################
#                        Contracts generation (remote buf)                                                            #
#######################################################################################################################

.PHONY: generate_contracts
generate_contracts: auth products asset_manager

# generate auth go files from remote repository which contains .proto contracts using buf
auth:
	$(BUF_GEN) "$(REPO)" --template buf.gen.yaml --path proto/auth/v1

# generate products go files from remote repository which contains .proto contracts using buf
products:
	$(BUF_GEN) "$(REPO)" --template buf.gen.yaml --path proto/products/v1

# generate asset_manager go files from remote repository which contains .proto contracts using buf
asset_manager:
	$(BUF_GEN) "$(REPO)" --template buf.gen.yaml --path proto/asset_manager/v1

#######################################################################################################################
#                                           Help                                                                      #
#######################################################################################################################

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
	   if (helpMessage) { \
	      helpCommand = substr($$1, 0, index($$1, ":")); \
	      helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
	      printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
	   } \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
