# Copyright (c) 2022 AlertAvert.com.  All rights reserved.
# (Reluctantly) Created by M. Massenzio, 2022-03-14

protos:
	protoc --proto_path=api/ \
               --go_out=./golang \
               --go-grpc_out=./golang \
               --go_opt=paths=source_relative \
               --go-grpc_opt=paths=source_relative \
               api/*.proto

clean:
	@rm -f golang/api/*.pb.go
