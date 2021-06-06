package main

import (
	"context"
	"fmt"
	"github.com/wanmei002/grpc-learn/ch10/client/service/discovery"
	"github.com/wanmei002/grpc-learn/ch10/server/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"log"
	"time"
)

func main(){
	for {
		log.Println("start sleep")
		time.Sleep(2*time.Second)
		AddOrder()
		time.Sleep(10*time.Second)
	}
}

func AddOrder() {
	dl, err := grpc.Dial(fmt.Sprintf("%s:///%s", discovery.GetSheme(), discovery.GetSvr()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithInsecure())

	if err != nil {
		log.Println("grpc dial failed; err:", err)
		return
	}

	defer dl.Close()

	client := order.NewOrderClient(dl)


	res, err := client.AddOrder(context.Background(), &order.OrderInfo{Name:"zzh", Items:[]string{"zzh", "zyn"}})

	if err != nil {
		log.Println("add order failed; err : ", err)
		return
	}

	log.Println("res :", res)
	log.Println("end")
}