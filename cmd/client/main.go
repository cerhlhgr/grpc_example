package main

import (
	"context"
	"fmt"
	desc "github.com/cerhlhgr/grpc_example/pkg/pack_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var opts []grpc.DialOption
	loggerFunc := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println(reply)
		return err
	}

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithUnaryInterceptor(loggerFunc))

	conn, err := grpc.NewClient("localhost:8080", opts...)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := desc.NewPackServiceClient(conn)

	feature, err := client.Get(context.Background(), &desc.PackRequest{Id: 1})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Вывод:\n ", feature)
}
