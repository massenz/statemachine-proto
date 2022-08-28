# statemachine-protos

The Protobuf definitions for the statemachine API.

[![Author](https://img.shields.io/badge/Author-M.%20Massenzio-green)](https://github.com/massenz)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release Automation](https://github.com/massenz/statemachine-proto/actions/workflows/release.yaml/badge.svg)](https://github.com/massenz/statemachine-proto/actions/workflows/release.yaml)

### Copyright & Licensing

**The code is copyright (c) 2022 AlertAvert.com. All rights reserved**<br>

# Overview

This repository contains the Protobuf definitions for the statemachine API.

The following is divided in two main sections, depending on whether you need to [build](#build-protocol-buffers) the Protobuf definitions from scratch, or whether you just need to [use](#usage) the generated code.

If you plan to [contribute](#contributing) to the project, please read the [Contributing](#contributing) section.

# Build Protocol Buffers

## Prerequisites

To generate the code, you will need to install `protoc` and, for Go code, the `protoc-gen-go` plugins; please follow the instructions below before attempting to running any of the `make` commands.

All the base classes are defined in the `api` folder and are used to (de)serialize state machines for storage in the database.

See [installation instructions](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers) for compiling protobufs for Go.

It mostly boils down to the following:

* Install `protoc` (the compiler) from [here](https://github.com/protocolbuffers/protobuf/releases), and install it somewhere on your `PATH` (and place the included `.proto` packages in the `include` folder somewhere where they will be found by `protoc`).
* Install the `protoc-gen-go` plugin, by running `go get -u github.com/golang/protobuf/protoc-gen-go`.
* Install the `protoc-gen-go-grpc` plugin, by running `go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc`.


## Build

We use `make` to build the code:

```shell
make protos && make mod
```
will do the trick; or just `make build` will combine the two.

## Release

The current version is kept in `build.settings`, remember **to update the version number before releasing**; the actual release is handled by the `release.yaml` workflow, which will automatically create a new release and upload the generated code, upon pushing to the `release` branch.

# Contributing

Contributions are welcome, and are greatly appreciated! Every little bit helps, and credit will always be given.

Please make sure you follow [Google Code Style guides](https://developers.google.com/protocol-buffers/docs/style) when contributing, and submit a pull request referencing the issue you are addressing.

# Usage

## Design

The main class is the `FiniteStateMachine` which, however, is really simple, as it just carries a reference to its `Configuration`, the current `state` and a list of `Events` which have occurred (possibly empty).

A `Configuration` describes all the `states` an FSM could go through, as well as the `Transitions` from one state to the next; a `Transition` is defined simply as the `event` (as an opaque string) which will cause the FSM to move `from` state `to` state.

While Protobufs cannot define (and even less, enforce) behaviors, the following should be respected by implementations:

* a `Configuration` should be immutable; if changes are desired, simply increase the `version` accordingly (using, e.g., [semantic versioning](https://semver.org/))
* equally, `FiniteStateMachine` should be deemed "immutable" and only be changed via ingesting `Event`s (hence, `config_id` once set should never change)
* `Event`s are only appended to the `history` which is never "rewritten"
* we would expect `event_id`s to be unique (within a defined namespace or domain)


## Importing

To use the generated code from a given release, you will need to import the package in your `go.mod` file:

```shell
└─( go get github.com/massenz/statemachine-proto/golang@v0.4.1-gdb3d0ac
```

making sure that the referenced version is one of the [Releases](https://github.com/massenz/statemachine-proto/releases); then in the go code, you can import the `api` package as usual:

```go
package main

import (
    "github.com/massenz/statemachine-proto/golang/api"
)

func main() {
    // ...
    response, err := client.ConsumeEvent(context.Background(),
        &api.EventRequest{
                Event: &api.Event{
                    EventId:   uuid.NewString(),
                    Timestamp: timestamppb.Now(),
                    Transition: &api.Transition{
                        Event: *event,
                    },
                    Originator: "gRPC Client",
                },
                Dest: *fsmId,
        })
}
```

see the [Go Statemachine project](https://github.com/massenz/go-statemachine) for an example of how to use the generated code.
