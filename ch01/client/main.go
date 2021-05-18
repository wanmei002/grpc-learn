package main

import (
	"context"
	"fmt"

	"github.com/wanmei002/grpc/server/ecommerce"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8093", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("client dial failed; err:%v\n", err)
		return
	}

	defer conn.Close()

	c := ecommerce.NewProductInfoClient(conn)

	p := &ecommerce.Product{
		Id:          "1",
		Name:        "c",
		Description: "desc",
	}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	r, err := c.AddProduct(context.Background(), p)

	if err != nil {
		fmt.Println("client addProduct err:", err)
	}

	fmt.Printf("client get product id :%+v\n", r)

}
