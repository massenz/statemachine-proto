// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: api/statemachine.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StatemachineServiceClient is the client API for StatemachineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatemachineServiceClient interface {
	// Creates an immutable Configuration, which can then be used to create FSMs whose
	// state transitions are defined in the configuration.
	// `Configuration`s are immutable and cannot be modified once created: attempting
	// to `PUT` a `Configuration` with the same `name` and `version` will cause an `AlreadyExists` error.
	PutConfiguration(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*PutResponse, error)
	// Retrieves all Configuration names, or versions.
	//
	// If the `StringValue` passed in is non-empty, the returned list will be of all
	// the fully qualified (`name:version`) configurations available on this server, where
	// the `Configuration.name` matches the value.
	//
	// If the value is empty, all the `Configuration` names (without the `version`) will be
	// returned.
	GetAllConfigurations(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*ListResponse, error)
	// Streams all the `Configuration` whose name matches the passed in StringValue.
	// Use this (instead of the `GetAllConfigurations()) if you want the full contents
	// of the configurations.
	StreamAllConfigurations(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (StatemachineService_StreamAllConfigurationsClient, error)
	// Retrieves a Configuration by its ID.
	GetConfiguration(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Configuration, error)
	// Creates a new FSM, using the Configuration identified by `config_id`.
	// The FSM itself is immutable, only its `state` can be mutated by processing Events.
	PutFiniteStateMachine(ctx context.Context, in *PutFsmRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// Retrieves an FSM by its Configuration `name` and Statemachine `ID`.
	GetFiniteStateMachine(ctx context.Context, in *GetFsmRequest, opts ...grpc.CallOption) (*FiniteStateMachine, error)
	// Looks up all the FSMs, with the given `config_name` that are in the given `state`
	// and returns their IDs.
	GetAllInState(ctx context.Context, in *GetFsmRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Streams the full contents of the `FiniteStateMachine` that match the request,
	// including their `id`.
	//
	// For this method call, the `state` may optionally be empty, all the FSMs that match
	// the configuration name will be returned.
	//
	// The server may optionally choose not to implement this feature, and return an
	// `InvalidRequest` if `state` is empty.
	StreamAllInstate(ctx context.Context, in *GetFsmRequest, opts ...grpc.CallOption) (StatemachineService_StreamAllInstateClient, error)
	// Process an Event for an FSM, identified by `id`.
	//
	// Events are processed asynchronously; the `outcome` will be empty, but
	// the returned `event_id` can be used later on to query the outcome of the
	// event processing (see `GetEventOutcome`).
	SendEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	// Get the outcome of an event processing, identified by the `event_id` returned by `ProcessEvent`.
	GetEventOutcome(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type statemachineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatemachineServiceClient(cc grpc.ClientConnInterface) StatemachineServiceClient {
	return &statemachineServiceClient{cc}
}

func (c *statemachineServiceClient) PutConfiguration(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/statemachine.v1beta.StatemachineService/PutConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statemachineServiceClient) GetAllConfigurations(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/statemachine.v1beta.StatemachineService/GetAllConfigurations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statemachineServiceClient) StreamAllConfigurations(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (StatemachineService_StreamAllConfigurationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &StatemachineService_ServiceDesc.Streams[0], "/statemachine.v1beta.StatemachineService/StreamAllConfigurations", opts...)
	if err != nil {
		return nil, err
	}
	x := &statemachineServiceStreamAllConfigurationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StatemachineService_StreamAllConfigurationsClient interface {
	Recv() (*Configuration, error)
	grpc.ClientStream
}

type statemachineServiceStreamAllConfigurationsClient struct {
	grpc.ClientStream
}

func (x *statemachineServiceStreamAllConfigurationsClient) Recv() (*Configuration, error) {
	m := new(Configuration)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *statemachineServiceClient) GetConfiguration(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/statemachine.v1beta.StatemachineService/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statemachineServiceClient) PutFiniteStateMachine(ctx context.Context, in *PutFsmRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/statemachine.v1beta.StatemachineService/PutFiniteStateMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statemachineServiceClient) GetFiniteStateMachine(ctx context.Context, in *GetFsmRequest, opts ...grpc.CallOption) (*FiniteStateMachine, error) {
	out := new(FiniteStateMachine)
	err := c.cc.Invoke(ctx, "/statemachine.v1beta.StatemachineService/GetFiniteStateMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statemachineServiceClient) GetAllInState(ctx context.Context, in *GetFsmRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/statemachine.v1beta.StatemachineService/GetAllInState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statemachineServiceClient) StreamAllInstate(ctx context.Context, in *GetFsmRequest, opts ...grpc.CallOption) (StatemachineService_StreamAllInstateClient, error) {
	stream, err := c.cc.NewStream(ctx, &StatemachineService_ServiceDesc.Streams[1], "/statemachine.v1beta.StatemachineService/StreamAllInstate", opts...)
	if err != nil {
		return nil, err
	}
	x := &statemachineServiceStreamAllInstateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StatemachineService_StreamAllInstateClient interface {
	Recv() (*PutResponse, error)
	grpc.ClientStream
}

type statemachineServiceStreamAllInstateClient struct {
	grpc.ClientStream
}

func (x *statemachineServiceStreamAllInstateClient) Recv() (*PutResponse, error) {
	m := new(PutResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *statemachineServiceClient) SendEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/statemachine.v1beta.StatemachineService/SendEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statemachineServiceClient) GetEventOutcome(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/statemachine.v1beta.StatemachineService/GetEventOutcome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatemachineServiceServer is the server API for StatemachineService service.
// All implementations must embed UnimplementedStatemachineServiceServer
// for forward compatibility
type StatemachineServiceServer interface {
	// Creates an immutable Configuration, which can then be used to create FSMs whose
	// state transitions are defined in the configuration.
	// `Configuration`s are immutable and cannot be modified once created: attempting
	// to `PUT` a `Configuration` with the same `name` and `version` will cause an `AlreadyExists` error.
	PutConfiguration(context.Context, *Configuration) (*PutResponse, error)
	// Retrieves all Configuration names, or versions.
	//
	// If the `StringValue` passed in is non-empty, the returned list will be of all
	// the fully qualified (`name:version`) configurations available on this server, where
	// the `Configuration.name` matches the value.
	//
	// If the value is empty, all the `Configuration` names (without the `version`) will be
	// returned.
	GetAllConfigurations(context.Context, *wrapperspb.StringValue) (*ListResponse, error)
	// Streams all the `Configuration` whose name matches the passed in StringValue.
	// Use this (instead of the `GetAllConfigurations()) if you want the full contents
	// of the configurations.
	StreamAllConfigurations(*wrapperspb.StringValue, StatemachineService_StreamAllConfigurationsServer) error
	// Retrieves a Configuration by its ID.
	GetConfiguration(context.Context, *wrapperspb.StringValue) (*Configuration, error)
	// Creates a new FSM, using the Configuration identified by `config_id`.
	// The FSM itself is immutable, only its `state` can be mutated by processing Events.
	PutFiniteStateMachine(context.Context, *PutFsmRequest) (*PutResponse, error)
	// Retrieves an FSM by its Configuration `name` and Statemachine `ID`.
	GetFiniteStateMachine(context.Context, *GetFsmRequest) (*FiniteStateMachine, error)
	// Looks up all the FSMs, with the given `config_name` that are in the given `state`
	// and returns their IDs.
	GetAllInState(context.Context, *GetFsmRequest) (*ListResponse, error)
	// Streams the full contents of the `FiniteStateMachine` that match the request,
	// including their `id`.
	//
	// For this method call, the `state` may optionally be empty, all the FSMs that match
	// the configuration name will be returned.
	//
	// The server may optionally choose not to implement this feature, and return an
	// `InvalidRequest` if `state` is empty.
	StreamAllInstate(*GetFsmRequest, StatemachineService_StreamAllInstateServer) error
	// Process an Event for an FSM, identified by `id`.
	//
	// Events are processed asynchronously; the `outcome` will be empty, but
	// the returned `event_id` can be used later on to query the outcome of the
	// event processing (see `GetEventOutcome`).
	SendEvent(context.Context, *EventRequest) (*EventResponse, error)
	// Get the outcome of an event processing, identified by the `event_id` returned by `ProcessEvent`.
	GetEventOutcome(context.Context, *EventRequest) (*EventResponse, error)
	mustEmbedUnimplementedStatemachineServiceServer()
}

// UnimplementedStatemachineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatemachineServiceServer struct {
}

func (UnimplementedStatemachineServiceServer) PutConfiguration(context.Context, *Configuration) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutConfiguration not implemented")
}
func (UnimplementedStatemachineServiceServer) GetAllConfigurations(context.Context, *wrapperspb.StringValue) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConfigurations not implemented")
}
func (UnimplementedStatemachineServiceServer) StreamAllConfigurations(*wrapperspb.StringValue, StatemachineService_StreamAllConfigurationsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAllConfigurations not implemented")
}
func (UnimplementedStatemachineServiceServer) GetConfiguration(context.Context, *wrapperspb.StringValue) (*Configuration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedStatemachineServiceServer) PutFiniteStateMachine(context.Context, *PutFsmRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutFiniteStateMachine not implemented")
}
func (UnimplementedStatemachineServiceServer) GetFiniteStateMachine(context.Context, *GetFsmRequest) (*FiniteStateMachine, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiniteStateMachine not implemented")
}
func (UnimplementedStatemachineServiceServer) GetAllInState(context.Context, *GetFsmRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllInState not implemented")
}
func (UnimplementedStatemachineServiceServer) StreamAllInstate(*GetFsmRequest, StatemachineService_StreamAllInstateServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAllInstate not implemented")
}
func (UnimplementedStatemachineServiceServer) SendEvent(context.Context, *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEvent not implemented")
}
func (UnimplementedStatemachineServiceServer) GetEventOutcome(context.Context, *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventOutcome not implemented")
}
func (UnimplementedStatemachineServiceServer) mustEmbedUnimplementedStatemachineServiceServer() {}

// UnsafeStatemachineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatemachineServiceServer will
// result in compilation errors.
type UnsafeStatemachineServiceServer interface {
	mustEmbedUnimplementedStatemachineServiceServer()
}

func RegisterStatemachineServiceServer(s grpc.ServiceRegistrar, srv StatemachineServiceServer) {
	s.RegisterService(&StatemachineService_ServiceDesc, srv)
}

func _StatemachineService_PutConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Configuration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatemachineServiceServer).PutConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statemachine.v1beta.StatemachineService/PutConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatemachineServiceServer).PutConfiguration(ctx, req.(*Configuration))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatemachineService_GetAllConfigurations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatemachineServiceServer).GetAllConfigurations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statemachine.v1beta.StatemachineService/GetAllConfigurations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatemachineServiceServer).GetAllConfigurations(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatemachineService_StreamAllConfigurations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrapperspb.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StatemachineServiceServer).StreamAllConfigurations(m, &statemachineServiceStreamAllConfigurationsServer{stream})
}

type StatemachineService_StreamAllConfigurationsServer interface {
	Send(*Configuration) error
	grpc.ServerStream
}

type statemachineServiceStreamAllConfigurationsServer struct {
	grpc.ServerStream
}

func (x *statemachineServiceStreamAllConfigurationsServer) Send(m *Configuration) error {
	return x.ServerStream.SendMsg(m)
}

func _StatemachineService_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatemachineServiceServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statemachine.v1beta.StatemachineService/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatemachineServiceServer).GetConfiguration(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatemachineService_PutFiniteStateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutFsmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatemachineServiceServer).PutFiniteStateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statemachine.v1beta.StatemachineService/PutFiniteStateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatemachineServiceServer).PutFiniteStateMachine(ctx, req.(*PutFsmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatemachineService_GetFiniteStateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFsmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatemachineServiceServer).GetFiniteStateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statemachine.v1beta.StatemachineService/GetFiniteStateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatemachineServiceServer).GetFiniteStateMachine(ctx, req.(*GetFsmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatemachineService_GetAllInState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFsmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatemachineServiceServer).GetAllInState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statemachine.v1beta.StatemachineService/GetAllInState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatemachineServiceServer).GetAllInState(ctx, req.(*GetFsmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatemachineService_StreamAllInstate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFsmRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StatemachineServiceServer).StreamAllInstate(m, &statemachineServiceStreamAllInstateServer{stream})
}

type StatemachineService_StreamAllInstateServer interface {
	Send(*PutResponse) error
	grpc.ServerStream
}

type statemachineServiceStreamAllInstateServer struct {
	grpc.ServerStream
}

func (x *statemachineServiceStreamAllInstateServer) Send(m *PutResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StatemachineService_SendEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatemachineServiceServer).SendEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statemachine.v1beta.StatemachineService/SendEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatemachineServiceServer).SendEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatemachineService_GetEventOutcome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatemachineServiceServer).GetEventOutcome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statemachine.v1beta.StatemachineService/GetEventOutcome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatemachineServiceServer).GetEventOutcome(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatemachineService_ServiceDesc is the grpc.ServiceDesc for StatemachineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatemachineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "statemachine.v1beta.StatemachineService",
	HandlerType: (*StatemachineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutConfiguration",
			Handler:    _StatemachineService_PutConfiguration_Handler,
		},
		{
			MethodName: "GetAllConfigurations",
			Handler:    _StatemachineService_GetAllConfigurations_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _StatemachineService_GetConfiguration_Handler,
		},
		{
			MethodName: "PutFiniteStateMachine",
			Handler:    _StatemachineService_PutFiniteStateMachine_Handler,
		},
		{
			MethodName: "GetFiniteStateMachine",
			Handler:    _StatemachineService_GetFiniteStateMachine_Handler,
		},
		{
			MethodName: "GetAllInState",
			Handler:    _StatemachineService_GetAllInState_Handler,
		},
		{
			MethodName: "SendEvent",
			Handler:    _StatemachineService_SendEvent_Handler,
		},
		{
			MethodName: "GetEventOutcome",
			Handler:    _StatemachineService_GetEventOutcome_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAllConfigurations",
			Handler:       _StatemachineService_StreamAllConfigurations_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamAllInstate",
			Handler:       _StatemachineService_StreamAllInstate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/statemachine.proto",
}
