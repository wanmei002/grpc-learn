package main

import (
	"context"
	"log"
	"net"
	"os"
	"sync"

	"github.com/wanmei002/grpc-learn/ch10/server/etcd"
	"github.com/wanmei002/grpc-learn/ch10/server/order"
	"google.golang.org/grpc"
)

type server struct {
	ip string
}

var orderDB sync.Map

func (s *server) AddOrder(ctx context.Context, orderInfo *order.OrderInfo) (*order.Res, error) {
	log.Println("ip:", s.ip)
	_, ok := orderDB.Load(orderInfo.Name)
	if ok {
		return &order.Res{Msg: "订单已存在"}, nil
	}

	orderDB.Store(orderInfo.Name, orderInfo)

	return &order.Res{Msg: "添加成功"}, nil
}

var serviceName = "/order.Order/"

func main() {
	input := os.Args
	if len(input) < 2 {
		log.Println("请输入监听的端口")
		return
	}

	ls, err := net.Listen("tcp", input[1])
	if err != nil {
		log.Println("listen failed; err:", err)
		return
	}
	defer ls.Close()
	gSvr := grpc.NewServer()

	order.RegisterOrderServer(gSvr, &server{ip:input[1]})

	// etcd 推送监听的端口
	err = etcd.Put(serviceName+input[1], input[1])
	if err != nil {
		log.Println("etcd put failed; err:", err)
		return
	}
	log.Println("server start; port:", input[1])

	if err = gSvr.Serve(ls); err != nil {
		log.Println("grpc serve failed; err:", err)
		return
	}
}
