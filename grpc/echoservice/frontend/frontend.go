package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	echo "github.com/appnet-org/appnet/grpc/echoservice/pb"
	"github.com/appnet-org/appnet/grpc/plugininterceptor"
)

func initHandler() (func(writer http.ResponseWriter, request *http.Request), func()) {

	conn, err := grpc.Dial(
		"server:9000",
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			plugininterceptor.ClientInterceptor("/interceptors.so"),
		),
	)
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	return func(writer http.ResponseWriter, request *http.Request) {
			request_body := strings.Replace(request.URL.String(), "/", "", -1)
			fmt.Printf("Got request with body: %s\n", request_body)

			c := echo.NewEchoServiceClient(conn)

			message := echo.Msg{
				Body: request_body,
			}

			response, err := c.Echo(context.Background(), &message)
			if err != nil {
				fmt.Fprintf(writer, "Echo server returns an error.\n")
				log.Printf("Error when calling echo: %s", err)
			} else {
				fmt.Fprintf(writer, "Echo request finished! Length of the request is %d\n", len(response.Body))
				log.Printf("Response from server: %s", response.Body)
			}
		}, func() {
			conn.Close()
		}
}

func main() {
	handler, cleanUp := initHandler()
	defer cleanUp()

	http.HandleFunc("/", handler)

	fmt.Printf("Starting frontend pod at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
