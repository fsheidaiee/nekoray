// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: libcore.proto

package gen

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

// LibcoreServiceClient is the client API for LibcoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibcoreServiceClient interface {
	Exit(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error)
	KeepAlive(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error)
	Start(ctx context.Context, in *LoadConfigReq, opts ...grpc.CallOption) (*ErrorResp, error)
	SetTun(ctx context.Context, in *SetTunReq, opts ...grpc.CallOption) (*ErrorResp, error)
	Stop(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*ErrorResp, error)
	Test(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestResp, error)
	QueryStats(ctx context.Context, in *QueryStatsReq, opts ...grpc.CallOption) (*QueryStatsResp, error)
	ListV2RayConnections(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*ListV2RayConnectionsResp, error)
}

type libcoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLibcoreServiceClient(cc grpc.ClientConnInterface) LibcoreServiceClient {
	return &libcoreServiceClient{cc}
}

func (c *libcoreServiceClient) Exit(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, "/libcore.LibcoreService/Exit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) KeepAlive(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, "/libcore.LibcoreService/KeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, "/libcore.LibcoreService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) Start(ctx context.Context, in *LoadConfigReq, opts ...grpc.CallOption) (*ErrorResp, error) {
	out := new(ErrorResp)
	err := c.cc.Invoke(ctx, "/libcore.LibcoreService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) SetTun(ctx context.Context, in *SetTunReq, opts ...grpc.CallOption) (*ErrorResp, error) {
	out := new(ErrorResp)
	err := c.cc.Invoke(ctx, "/libcore.LibcoreService/SetTun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) Stop(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*ErrorResp, error) {
	out := new(ErrorResp)
	err := c.cc.Invoke(ctx, "/libcore.LibcoreService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) Test(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestResp, error) {
	out := new(TestResp)
	err := c.cc.Invoke(ctx, "/libcore.LibcoreService/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) QueryStats(ctx context.Context, in *QueryStatsReq, opts ...grpc.CallOption) (*QueryStatsResp, error) {
	out := new(QueryStatsResp)
	err := c.cc.Invoke(ctx, "/libcore.LibcoreService/QueryStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) ListV2RayConnections(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*ListV2RayConnectionsResp, error) {
	out := new(ListV2RayConnectionsResp)
	err := c.cc.Invoke(ctx, "/libcore.LibcoreService/ListV2rayConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibcoreServiceServer is the server API for LibcoreService service.
// All implementations must embed UnimplementedLibcoreServiceServer
// for forward compatibility
type LibcoreServiceServer interface {
	Exit(context.Context, *EmptyReq) (*EmptyResp, error)
	KeepAlive(context.Context, *EmptyReq) (*EmptyResp, error)
	Update(context.Context, *UpdateReq) (*UpdateResp, error)
	Start(context.Context, *LoadConfigReq) (*ErrorResp, error)
	SetTun(context.Context, *SetTunReq) (*ErrorResp, error)
	Stop(context.Context, *EmptyReq) (*ErrorResp, error)
	Test(context.Context, *TestReq) (*TestResp, error)
	QueryStats(context.Context, *QueryStatsReq) (*QueryStatsResp, error)
	ListV2RayConnections(context.Context, *EmptyReq) (*ListV2RayConnectionsResp, error)
	mustEmbedUnimplementedLibcoreServiceServer()
}

// UnimplementedLibcoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLibcoreServiceServer struct {
}

func (UnimplementedLibcoreServiceServer) Exit(context.Context, *EmptyReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exit not implemented")
}
func (UnimplementedLibcoreServiceServer) KeepAlive(context.Context, *EmptyReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (UnimplementedLibcoreServiceServer) Update(context.Context, *UpdateReq) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLibcoreServiceServer) Start(context.Context, *LoadConfigReq) (*ErrorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedLibcoreServiceServer) SetTun(context.Context, *SetTunReq) (*ErrorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTun not implemented")
}
func (UnimplementedLibcoreServiceServer) Stop(context.Context, *EmptyReq) (*ErrorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedLibcoreServiceServer) Test(context.Context, *TestReq) (*TestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedLibcoreServiceServer) QueryStats(context.Context, *QueryStatsReq) (*QueryStatsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStats not implemented")
}
func (UnimplementedLibcoreServiceServer) ListV2RayConnections(context.Context, *EmptyReq) (*ListV2RayConnectionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListV2RayConnections not implemented")
}
func (UnimplementedLibcoreServiceServer) mustEmbedUnimplementedLibcoreServiceServer() {}

// UnsafeLibcoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibcoreServiceServer will
// result in compilation errors.
type UnsafeLibcoreServiceServer interface {
	mustEmbedUnimplementedLibcoreServiceServer()
}

func RegisterLibcoreServiceServer(s grpc.ServiceRegistrar, srv LibcoreServiceServer) {
	s.RegisterService(&LibcoreService_ServiceDesc, srv)
}

func _LibcoreService_Exit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).Exit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libcore.LibcoreService/Exit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).Exit(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libcore.LibcoreService/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).KeepAlive(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libcore.LibcoreService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libcore.LibcoreService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).Start(ctx, req.(*LoadConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_SetTun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTunReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).SetTun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libcore.LibcoreService/SetTun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).SetTun(ctx, req.(*SetTunReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libcore.LibcoreService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).Stop(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libcore.LibcoreService/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).Test(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_QueryStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStatsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).QueryStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libcore.LibcoreService/QueryStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).QueryStats(ctx, req.(*QueryStatsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_ListV2RayConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).ListV2RayConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libcore.LibcoreService/ListV2rayConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).ListV2RayConnections(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LibcoreService_ServiceDesc is the grpc.ServiceDesc for LibcoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LibcoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libcore.LibcoreService",
	HandlerType: (*LibcoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exit",
			Handler:    _LibcoreService_Exit_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _LibcoreService_KeepAlive_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LibcoreService_Update_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _LibcoreService_Start_Handler,
		},
		{
			MethodName: "SetTun",
			Handler:    _LibcoreService_SetTun_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _LibcoreService_Stop_Handler,
		},
		{
			MethodName: "Test",
			Handler:    _LibcoreService_Test_Handler,
		},
		{
			MethodName: "QueryStats",
			Handler:    _LibcoreService_QueryStats_Handler,
		},
		{
			MethodName: "ListV2rayConnections",
			Handler:    _LibcoreService_ListV2RayConnections_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "libcore.proto",
}
