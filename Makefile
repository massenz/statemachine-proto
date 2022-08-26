# Copyright (c) 2022 AlertAvert.com.  All rights reserved.
# (Reluctantly) Created by M. Massenzio, 2022-03-14

go-dir=golang

$(go-dir)/api/statemachine.pb.go: api/statemachine.proto
	mkdir -p $(go-dir)
	protoc --go_out=./$(go-dir) --go_opt=paths=source_relative \
         api/*.proto

$(go-dir)/api/statemachine_grpc.pb.go: api/statemachine.proto
	protoc --go-grpc_out=./$(go-dir) --go-grpc_opt=paths=source_relative \
         api/*.proto

protos: $(go-dir)/api/statemachine.pb.go \
        $(go-dir)/api/statemachine_grpc.pb.go

golang/go.mod:
	cd $(go-dir) && go mod init github.com/massenz/statemachine-proto/golang

golang/go.sum: golang/go.mod
		cd $(go-dir) && go mod tidy

mod: golang/go.sum golang/go.mod

clean:
	@rm -rf $(go-dir)
