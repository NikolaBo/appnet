package main

import (
	"github.com/appnet-org/appnet/grpc/echoservice/interceptors/acl"
	"github.com/appnet-org/appnet/grpc/echoservice/interceptors/cache"
	"github.com/appnet-org/appnet/grpc/echoservice/interceptors/fault"
	"github.com/appnet-org/appnet/grpc/echoservice/interceptors/mutate"
	"google.golang.org/grpc"
)

type interceptInit struct{}

func (interceptInit) ClientInterceptors() []grpc.UnaryClientInterceptor {
	return []grpc.UnaryClientInterceptor{cache.CacheClient(), mutate.MutateClient(), fault.FaultClient()}
}

func (interceptInit) ServerInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{acl.ACLServer(), mutate.MutateServer(), fault.FaultServer()}
}

var InterceptInit interceptInit
