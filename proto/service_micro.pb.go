// Code generated by protoc-gen-micro
// source: service.proto
package service

import (
	"context"

	micro_api "github.com/unistack-org/micro/v3/api"
	micro_client "github.com/unistack-org/micro/v3/client"
	micro_server "github.com/unistack-org/micro/v3/server"
)

// NewLoggerEndpoints provides api endpoints metdata for Logger service
func NewLoggerEndpoints() []*micro_api.Endpoint {
	var endpoints []*micro_api.Endpoint
	return endpoints
}

// LoggerService interface
type LoggerService interface {
	Log(context.Context, *Message, ...micro_client.CallOption) (*Empty, error)
}

// Micro server stuff

// LoggerHandler server handler
type LoggerHandler interface {
	Log(context.Context, *Message, *Empty) error
}

// RegisterLoggerHandler registers server handler
func RegisterLoggerHandler(s micro_server.Server, sh LoggerHandler, opts ...micro_server.HandlerOption) error {
	type logger interface {
		Log(context.Context, *Message, *Empty) error
	}
	type Logger struct {
		logger
	}
	h := &loggerHandler{sh}
	for _, endpoint := range NewLoggerEndpoints() {
		opts = append(opts, micro_api.WithEndpoint(endpoint))
	}
	return s.Handle(s.NewHandler(&Logger{h}, opts...))
}