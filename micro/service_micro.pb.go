// Code generated by protoc-gen-go-micro. DO NOT EDIT.
// protoc-gen-go-micro version: v3.5.3
// source: service.proto

package servicepb

import (
	context "context"
	proto "go.unistack.org/micro-logger-service/v3/proto"
	api "go.unistack.org/micro/v3/api"
	client "go.unistack.org/micro/v3/client"
)

var (
	LoggerServiceName = "LoggerService"

	LoggerServiceEndpoints = []api.Endpoint{}
)

func NewLoggerServiceEndpoints() []api.Endpoint {
	return LoggerServiceEndpoints
}

type LoggerServiceClient interface {
	Log(ctx context.Context, req *proto.LogReq, opts ...client.CallOption) (*proto.LogRsp, error)
}

type LoggerServiceServer interface {
	Log(ctx context.Context, req *proto.LogReq, rsp *proto.LogRsp) error
}
