package main

import (
    "context"
    "github.com/wanmei002/grpc-learn/ch11/client/rpc/order"
    "github.com/wanmei002/grpc-learn/ch11/rpc/pb"
    "log"
)

func main(){
    in := &pb.OrderInfo{
        Name:  "zzh1",
        Items: append([]string(nil), "apple", "banner"),
        Desc:  "有钱",
    }
    res, err := order.AddOrder(in, context.Background())
    defer order.CloseConn()
    if err != nil {
        log.Println("add order failed; err:", err)
        return
    }
    
    log.Printf("res:%+v\n", res)
}
