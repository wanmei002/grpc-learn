package main

import (
	"context"
	"fmt"
	"io"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/wanmei002/grpc-learn/ch02/server/product"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8903", grpc.WithInsecure())
	if err != nil {
		fmt.Println("client dial failed; err :", err)
		return
	}

	defer conn.Close()

	c := product.NewOrderManagementClient(conn)

	searchStream, err := c.SearchOrders(context.Background(), &wrappers.StringValue{Value: "zzh"})
	if err != nil {
		fmt.Println("client grpc failed; err :", err)
		return
	}

	for {
		order, err := searchStream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("stream end")
				return
			}
			fmt.Println("clinet recv failed; err :", err)
			return
		}

		fmt.Printf("recv order info : %+v\n", order)
	}
}
