module github.com/appnet-org/appnet/grpc/echoservice

go 1.22.1

require (
	github.com/appnet-org/appnet/grpc/plugininterceptor v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.24.0
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
)

require (
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
)

replace github.com/appnet-org/appnet/grpc/plugininterceptor => ../plugininterceptor
