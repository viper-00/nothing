// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MonitorDataServiceClient is the client API for MonitorDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitorDataServiceClient interface {
	HandlePing(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*Message, error)
	IsUp(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*IsActive, error)
	InitAgent(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*Message, error)
	HandleMonitorData(ctx context.Context, in *MonitorData, opts ...grpc.CallOption) (*Message, error)
	HandleCustomMonitorData(ctx context.Context, in *MonitorData, opts ...grpc.CallOption) (*Message, error)
	HandleMonitorDataRequest(ctx context.Context, in *MonitorDataRequest, opts ...grpc.CallOption) (*MonitorData, error)
	HandleCustomMetricNameRequest(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*Message, error)
	HandleAgentIdsRequest(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Message, error)
}

type monitorDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorDataServiceClient(cc grpc.ClientConnInterface) MonitorDataServiceClient {
	return &monitorDataServiceClient{cc}
}

func (c *monitorDataServiceClient) HandlePing(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.MonitorDataService/HandlePing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorDataServiceClient) IsUp(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*IsActive, error) {
	out := new(IsActive)
	err := c.cc.Invoke(ctx, "/api.MonitorDataService/IsUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorDataServiceClient) InitAgent(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.MonitorDataService/InitAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorDataServiceClient) HandleMonitorData(ctx context.Context, in *MonitorData, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.MonitorDataService/HandleMonitorData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorDataServiceClient) HandleCustomMonitorData(ctx context.Context, in *MonitorData, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.MonitorDataService/HandleCustomMonitorData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorDataServiceClient) HandleMonitorDataRequest(ctx context.Context, in *MonitorDataRequest, opts ...grpc.CallOption) (*MonitorData, error) {
	out := new(MonitorData)
	err := c.cc.Invoke(ctx, "/api.MonitorDataService/HandleMonitorDataRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorDataServiceClient) HandleCustomMetricNameRequest(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.MonitorDataService/HandleCustomMetricNameRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorDataServiceClient) HandleAgentIdsRequest(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.MonitorDataService/HandleAgentIdsRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorDataServiceServer is the server API for MonitorDataService service.
// All implementations must embed UnimplementedMonitorDataServiceServer
// for forward compatibility
type MonitorDataServiceServer interface {
	HandlePing(context.Context, *ServerInfo) (*Message, error)
	IsUp(context.Context, *ServerInfo) (*IsActive, error)
	InitAgent(context.Context, *ServerInfo) (*Message, error)
	HandleMonitorData(context.Context, *MonitorData) (*Message, error)
	HandleCustomMonitorData(context.Context, *MonitorData) (*Message, error)
	HandleMonitorDataRequest(context.Context, *MonitorDataRequest) (*MonitorData, error)
	HandleCustomMetricNameRequest(context.Context, *ServerInfo) (*Message, error)
	HandleAgentIdsRequest(context.Context, *Void) (*Message, error)
	mustEmbedUnimplementedMonitorDataServiceServer()
}

// UnimplementedMonitorDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMonitorDataServiceServer struct {
}

func (UnimplementedMonitorDataServiceServer) HandlePing(context.Context, *ServerInfo) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandlePing not implemented")
}
func (UnimplementedMonitorDataServiceServer) IsUp(context.Context, *ServerInfo) (*IsActive, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUp not implemented")
}
func (UnimplementedMonitorDataServiceServer) InitAgent(context.Context, *ServerInfo) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitAgent not implemented")
}
func (UnimplementedMonitorDataServiceServer) HandleMonitorData(context.Context, *MonitorData) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleMonitorData not implemented")
}
func (UnimplementedMonitorDataServiceServer) HandleCustomMonitorData(context.Context, *MonitorData) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleCustomMonitorData not implemented")
}
func (UnimplementedMonitorDataServiceServer) HandleMonitorDataRequest(context.Context, *MonitorDataRequest) (*MonitorData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleMonitorDataRequest not implemented")
}
func (UnimplementedMonitorDataServiceServer) HandleCustomMetricNameRequest(context.Context, *ServerInfo) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleCustomMetricNameRequest not implemented")
}
func (UnimplementedMonitorDataServiceServer) HandleAgentIdsRequest(context.Context, *Void) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleAgentIdsRequest not implemented")
}
func (UnimplementedMonitorDataServiceServer) mustEmbedUnimplementedMonitorDataServiceServer() {}

// UnsafeMonitorDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitorDataServiceServer will
// result in compilation errors.
type UnsafeMonitorDataServiceServer interface {
	mustEmbedUnimplementedMonitorDataServiceServer()
}

func RegisterMonitorDataServiceServer(s grpc.ServiceRegistrar, srv MonitorDataServiceServer) {
	s.RegisterService(&MonitorDataService_ServiceDesc, srv)
}

func _MonitorDataService_HandlePing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorDataServiceServer).HandlePing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MonitorDataService/HandlePing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorDataServiceServer).HandlePing(ctx, req.(*ServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorDataService_IsUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorDataServiceServer).IsUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MonitorDataService/IsUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorDataServiceServer).IsUp(ctx, req.(*ServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorDataService_InitAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorDataServiceServer).InitAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MonitorDataService/InitAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorDataServiceServer).InitAgent(ctx, req.(*ServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorDataService_HandleMonitorData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorDataServiceServer).HandleMonitorData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MonitorDataService/HandleMonitorData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorDataServiceServer).HandleMonitorData(ctx, req.(*MonitorData))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorDataService_HandleCustomMonitorData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorDataServiceServer).HandleCustomMonitorData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MonitorDataService/HandleCustomMonitorData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorDataServiceServer).HandleCustomMonitorData(ctx, req.(*MonitorData))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorDataService_HandleMonitorDataRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorDataServiceServer).HandleMonitorDataRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MonitorDataService/HandleMonitorDataRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorDataServiceServer).HandleMonitorDataRequest(ctx, req.(*MonitorDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorDataService_HandleCustomMetricNameRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorDataServiceServer).HandleCustomMetricNameRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MonitorDataService/HandleCustomMetricNameRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorDataServiceServer).HandleCustomMetricNameRequest(ctx, req.(*ServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorDataService_HandleAgentIdsRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorDataServiceServer).HandleAgentIdsRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MonitorDataService/HandleAgentIdsRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorDataServiceServer).HandleAgentIdsRequest(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// MonitorDataService_ServiceDesc is the grpc.ServiceDesc for MonitorDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MonitorDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.MonitorDataService",
	HandlerType: (*MonitorDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandlePing",
			Handler:    _MonitorDataService_HandlePing_Handler,
		},
		{
			MethodName: "IsUp",
			Handler:    _MonitorDataService_IsUp_Handler,
		},
		{
			MethodName: "InitAgent",
			Handler:    _MonitorDataService_InitAgent_Handler,
		},
		{
			MethodName: "HandleMonitorData",
			Handler:    _MonitorDataService_HandleMonitorData_Handler,
		},
		{
			MethodName: "HandleCustomMonitorData",
			Handler:    _MonitorDataService_HandleCustomMonitorData_Handler,
		},
		{
			MethodName: "HandleMonitorDataRequest",
			Handler:    _MonitorDataService_HandleMonitorDataRequest_Handler,
		},
		{
			MethodName: "HandleCustomMetricNameRequest",
			Handler:    _MonitorDataService_HandleCustomMetricNameRequest_Handler,
		},
		{
			MethodName: "HandleAgentIdsRequest",
			Handler:    _MonitorDataService_HandleAgentIdsRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
