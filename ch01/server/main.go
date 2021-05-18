package main

import (
	"context"
	"fmt"
	"net"

	"github.com/wanmei002/grpc/server/ecommerce"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) AddProduct(ctx context.Context, in *ecommerce.Product) (*ecommerce.ProductID, error) {
	fmt.Println("server get product info :", in)
	return &ecommerce.ProductID{Value: in.Id}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8093")

	if err != nil {
		fmt.Println("server listen failed;err : ", err)
		return
	}

	s := grpc.NewServer()
	ecommerce.RegisterProductInfoServer(s, &server{})
	fmt.Println("server begin")
	if err := s.Serve(lis); err != nil {
		fmt.Println("server : grpc failed; err:", err)
	}
}
