package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/net/context"

	echo "github.com/appnet-org/appnet/grpc/echoservice/pb"
	"google.golang.org/grpc"
)

type server struct {
	echo.UnimplementedEchoServiceServer
}

type InterceptInit interface {
	ServerInterceptors() []grpc.UnaryServerInterceptor
}

func (s *server) Echo(ctx context.Context, x *echo.Msg) (*echo.Msg, error) {
	log.Printf("got: [%s]", x.GetBody())

	hostname, _ := os.Hostname()
	appendedBody := fmt.Sprintf("You've hit %s\n", hostname)
	msg := &echo.Msg{
		Body: appendedBody,
	}
	return msg, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
	)
	fmt.Printf("Starting server pod at port 9000\n")

	echo.RegisterEchoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
