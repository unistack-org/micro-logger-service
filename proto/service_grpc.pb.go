// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service.proto

package servicepb

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

// LoggerServiceClient is the client API for LoggerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggerServiceClient interface {
	Log(ctx context.Context, in *LogReq, opts ...grpc.CallOption) (*LogRsp, error)
}

type loggerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggerServiceClient(cc grpc.ClientConnInterface) LoggerServiceClient {
	return &loggerServiceClient{cc}
}

func (c *loggerServiceClient) Log(ctx context.Context, in *LogReq, opts ...grpc.CallOption) (*LogRsp, error) {
	out := new(LogRsp)
	err := c.cc.Invoke(ctx, "/servicepb.LoggerService/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerServiceServer is the server API for LoggerService service.
// All implementations must embed UnimplementedLoggerServiceServer
// for forward compatibility
type LoggerServiceServer interface {
	Log(context.Context, *LogReq) (*LogRsp, error)
	mustEmbedUnimplementedLoggerServiceServer()
}

// UnimplementedLoggerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoggerServiceServer struct {
}

func (UnimplementedLoggerServiceServer) Log(context.Context, *LogReq) (*LogRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedLoggerServiceServer) mustEmbedUnimplementedLoggerServiceServer() {}

// UnsafeLoggerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggerServiceServer will
// result in compilation errors.
type UnsafeLoggerServiceServer interface {
	mustEmbedUnimplementedLoggerServiceServer()
}

func RegisterLoggerServiceServer(s grpc.ServiceRegistrar, srv LoggerServiceServer) {
	s.RegisterService(&LoggerService_ServiceDesc, srv)
}

func _LoggerService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicepb.LoggerService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServiceServer).Log(ctx, req.(*LogReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LoggerService_ServiceDesc is the grpc.ServiceDesc for LoggerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoggerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "servicepb.LoggerService",
	HandlerType: (*LoggerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _LoggerService_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
