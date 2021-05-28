package main

import (
    "context"
    pb "github.com/wanmei002/grpc-learn/ch06/server/product"
    "google.golang.org/grpc"
    "log"
)

func main(){
    d, err := grpc.Dial(":8093", grpc.WithInsecure())
    if err != nil {
        log.Println("grpc dial failed; err:", err)
        return
    }
    
    defer d.Close()
    
    client := pb.NewOrderServerClient(d)
    ctx, _ := context.WithTimeout(context.Background(), 2e9)
    res, err := client.AddOrder(ctx, &pb.Order{Name: "zzh", Items: []string{"hello", "zzh"}})
    
    if err != nil {
        log.Println("add order failed; err:", err)
        return
    }
    
    log.Printf("res:[%+v]\n", res)
}
