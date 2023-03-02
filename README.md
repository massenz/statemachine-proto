# statemachine-protos

The Protobuf definitions for the statemachine API.

[![Author](https://img.shields.io/badge/Author-M.%20Massenzio-green)](https://github.com/massenz)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://github.com/massenz/statemachine-proto/actions/workflows/release.yaml/badge.svg)](https://github.com/massenz/statemachine-proto/actions/workflows/release.yaml)

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

# gRPC API

> For a full description and documentation of the gRPC API, please see the [Protocol Buffer definition](api/statemachine.proto).

These are the methods that are currently defined (as of version `v1beta` of the API, release `v1.1.0-beta-g1fc5dd8`) of the `statemachine-proto` repository.

> **NOTE**
>
> It appears that it is currently not possible to export a Postman Collection containing gRPC method calls; we are looking into ways to share example method calls.


#### Creation methods

```
  // Creates an immutable Configuration.
  rpc PutConfiguration(Configuration) returns (PutResponse);

  // Creates a new FSM, using the Configuration identified by `config_id`.
  rpc PutFiniteStateMachine(PutFsmRequest) returns (PutResponse);
```

Examples:

1. Creating a Configuration

```
{
    "name": "test.orders",
    "version": "v3",
    "starting_state": "start",
    "states": ["start", "stop"],
    "transitions": [
        {"event": "run", "from": "start", "to": "stop"}
    ]
}
```

2. Create a FiniteStateMachine

```
{
    "id": "12345",
    "fsm": {"config_id": "test.orders:v4"}
}
```

3. Create a FiniteStateMachine omitting the ID, specifying the starting state

```
PutFiniteStateMachine

{
    "fsm": {
        "config_id": "test.orders:v2",
        "state": "shipped"
    }
}
```
response:
```
Status code: 0 (OK)

{
    "id": "b71e000e-3463-4ca9-8719-9bc74e75c173",
    "fsm": {
        "history": [],
        "config_id": "test.orders:v2",
        "state": "shipped"
    }
}
```

#### Lookup methods

```
  // Retrieves all Configuration names, or versions.
  rpc GetAllConfigurations(google.protobuf.StringValue) returns (ListResponse);

  // Retrieves a Configuration by its ID.
  rpc GetConfiguration(google.protobuf.StringValue) returns (Configuration);

  // Retrieves an FSM by its Configuration `name` and Statemachine `ID`.
  rpc GetFiniteStateMachine(GetFsmRequest) returns (FiniteStateMachine);

  // Looks up all the FSMs in the given `state`
  rpc GetAllInState(GetFsmRequest) returns (ListResponse);
```

Examples:

1. To get all FSMs in the `shipped` state:

```
GetAllInState
{
    "config": "test.orders",
    "state": "shipped"
}
```
```
Status Code: 0 (OK)
{
    "ids": [
        "b71e000e-3463-4ca9-8719-9bc74e75c173"
    ]
}
```

2. To find all configurations available in the store use the `GetAllConfigurations` with an empty request. The response will have all configuration names:

```
Status Code: 0 (OK)
{
    "ids": [
        "returns",
        "users",
        "devices",
        "test.orders"
    ]
}
```
and then to retrieve all versions, pass in the `name` of the configuration:

```
GetAllConfigurations
{"value": "test.orders"}
```
```
Status Code: 0 (OK)
{
    "ids": [
        "test.orders:v1",
        "test.orders:v2"
    ]
}
```

#### Streaming methods

```
  // Streams all the `Configuration` whose name matches the passed in StringValue.
  rpc StreamAllConfigurations(google.protobuf.StringValue) returns (stream Configuration);

  // Streams the full contents of `FiniteStateMachines` in the given `state`
  rpc StreamAllInstate(GetFsmRequest) returns (stream PutResponse);
```

#### Event Management

```
  // Process an Event for an FSM, identified by `id`.
  rpc SendEvent(EventRequest) returns (EventResponse);

  // Get the outcome of an event processing, identified by the `event_id` returned by `ProcessEvent`.
  rpc GetEventOutcome(EventRequest) returns (EventResponse);
```
