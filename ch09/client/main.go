package main

import (
    "context"
    "fmt"
    "github.com/wanmei002/grpc-learn/ch09/client/balance"
    "github.com/wanmei002/grpc-learn/ch09/server/order"
    "google.golang.org/grpc"
    "google.golang.org/grpc/balancer/roundrobin"
    "io"
    "log"
)


func main(){

    di, err := grpc.Dial(fmt.Sprintf("%s:///%s", balance.LoopScheme, balance.LoopServiceName),
        grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
        grpc.WithInsecure())
    if err != nil {
        log.Println("grpc dial failed; err :", err)
        return
    }
    client := order.NewOrderClient(di)
    res, err := client.AddOrder(context.Background(), &order.OrderInfo{Name: "first", Items: []string{"zzh", "zyn"}})
    if err != nil {
        if err == io.EOF {
            log.Println("send over")
        } else {
            log.Println("client add order err:", err)
            return
        }
    }
    log.Printf("add order return msg:%+v\n", res)
    if res.Code != 0 {
        log.Println("add order failed; err :", res.Msg, "; return info:", res)
        return
    }
    log.Println("client end")
}
