# Copyright (c) 2022 AlertAvert.com.  All rights reserved.
# Licensed under the Apache License, Version 2.0
#
# Author: Marco Massenzio (marco@alertavert.com)

go-dir=golang

$(go-dir)/api/statemachine.pb.go: api/statemachine.proto
	@mkdir -p $(go-dir)
	protoc --go_out=./$(go-dir) --go_opt=paths=source_relative \
         -I thirdparty -I . api/*.proto

$(go-dir)/api/statemachine_grpc.pb.go: api/statemachine.proto
	protoc --go-grpc_out=./$(go-dir) --go-grpc_opt=paths=source_relative \
         -I thirdparty -I . api/*.proto

protos: $(go-dir)/api/statemachine.pb.go \
        $(go-dir)/api/statemachine_grpc.pb.go

$(go-dir)/go.mod:
	@mkdir -p $(go-dir)
	cd $(go-dir) && go mod init github.com/massenz/statemachine-proto/golang

$(go-dir)/go.sum: $(go-dir)/go.mod
	cd $(go-dir) && go mod tidy

mod: $(go-dir)/go.sum $(go-dir)/go.mod

build: protos mod

clean:
	@rm -rf $(go-dir)
