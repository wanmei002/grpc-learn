package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/wanmei002/grpc-learn/ch03/server/product"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8093", grpc.WithInsecure())
	if err != nil {
		log.Println("client dial failed; err:", err)
		return
	}

	c := product.NewProductClient(conn)
	searchStream, err := c.SearchGoods(context.Background())
	if err != nil {
		log.Println("client search failed; err:", err)
		return
	}

	// 查询订单
	log.Println("search world")
	if err := searchStream.Send(&wrappers.StringValue{Value: "world"}); err != nil {
		log.Println("client send search value failed; value:", "world")
		return
	}

	if err := searchStream.Send(&wrappers.StringValue{Value: "zyn"}); err != nil {
		log.Println("client send search value failed; value:", "zyn")
		return
	}

	res, err := searchStream.CloseAndRecv()
	if err != nil {
		log.Println("client get result failed; err:", err)
		return
	}

	fmt.Println("recv end")
	fmt.Printf("res:%+v\n", res)
}
