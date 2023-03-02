# Copyright (c) 2022 AlertAvert.com.  All rights reserved.
# Licensed under the Apache License, Version 2.0
#
# Author: Marco Massenzio (marco@alertavert.com)

go-dir := golang
version := 1.2.0
release := v$(version)-g$(shell git rev-parse --short HEAD)

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

version:
	@echo $(version)

release:
	@echo $(release)

##@ Golang

$(go-dir)/api/statemachine.pb.go: api/statemachine.proto
	@mkdir -p $(go-dir)
	protoc --go_out=./$(go-dir) --go_opt=paths=source_relative \
         api/*.proto

$(go-dir)/api/statemachine_grpc.pb.go: api/statemachine.proto
	protoc --go-grpc_out=./$(go-dir) --go-grpc_opt=paths=source_relative \
         api/*.proto

protos: $(go-dir)/api/statemachine.pb.go $(go-dir)/api/statemachine_grpc.pb.go  ## Generates Golang Protobuf bindings

$(go-dir)/go.mod:
	@mkdir -p $(go-dir)
	cd $(go-dir) && go mod init github.com/massenz/statemachine-proto/golang

$(go-dir)/go.sum: $(go-dir)/go.mod
	cd $(go-dir) && go mod tidy

mod: $(go-dir)/go.sum $(go-dir)/go.mod  ## Initializes the go.mod

build: protos mod	## Builds the Golang project structure

clean:	## Cleans up the Golang generated files
	@rm -rf $(go-dir)
