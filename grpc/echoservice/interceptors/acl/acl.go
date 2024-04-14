package acl

import (
	"log"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	echo "github.com/appnet-org/appnet/grpc/echoservice/pb"
	"google.golang.org/grpc"
)

var (
	blockedCount int32
	countMutex   sync.Mutex
)

func ACLClient(optFuncs ...CallOption) grpc.UnaryClientInterceptor {
	// Combine the default options with any user-provided options.
	intOpts := reuseOrNewWithCallOptions(defaultOptions, optFuncs)

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// Split the gRPC CallOptions into two sets: gRPC-specific options and null options.
		grpcOpts, nullOptions := filterCallOptions(opts)

		// Combine the internal options with the null options.
		callOpts := reuseOrNewWithCallOptions(intOpts, nullOptions)

		log.Println("Hello from ACLUnaryClientInterceptor")
		if m, ok := req.(*echo.Msg); ok {

			if m.GetBody() == callOpts.content {
				// Xiangfeng: This does not work for client because there is a new connection everytime. However, it work for server
				// callOpts.blockedCount += 1

				// Xiangfeng: that's why we are using a mutex here, but the performance might hurt
				// XZ: I did a quick test, performance seems to be fine
				countMutex.Lock() // Lock the mutex before updating the counter
				blockedCount++
				count := blockedCount // Copy the value to avoid race condition in Printf
				countMutex.Unlock()

				log.Printf("Request blocked by ACL. (%d Request Blocked)", count)
				// log.Printf("Request blocked by ACL. (%d Request Blocked)", callOpts.blockedCount)

				return status.Error(codes.Aborted, "request blocked by ACL.")
			}
		}
		err := invoker(ctx, method, req, reply, cc, grpcOpts...)
		return err
	}
}

func ACLServer(optFuncs ...CallOption) grpc.UnaryServerInterceptor {
	// Combine the default options with any user-provided options.
	intOpts := reuseOrNewWithCallOptions(defaultOptions, optFuncs)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("Running ACLUnaryServerInterceptor")
		if m, ok := req.(*echo.Msg); ok {
			if m.GetBody() == intOpts.content {
				// Xiangfeng: This does not work for client because there is a new connection everytime. However, it work for server
				// callOpts.blockedCount += 1

				// Xiangfeng: that's why we are using a mutex here, but the performance might hurt

				countMutex.Lock() // Lock the mutex before updating the counter
				blockedCount++
				count := blockedCount // Copy the value to avoid race condition in Printf
				countMutex.Unlock()
				log.Printf("Request blocked by ACL. (%d Request Blocked)", count)

				return nil, status.Error(codes.Aborted, "request blocked by ACL.")
			}
		}
		return handler(ctx, req)
	}
}
