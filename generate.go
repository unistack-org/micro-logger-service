package service

//go:generate protoc -I./proto -I. --go-grpc_out=paths=source_relative:./proto --go_out=paths=source_relative:./proto --micro_out=components=micro|grpc,debug=true,paths=source_relative:./proto proto/service.proto
