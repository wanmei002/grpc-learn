package main

import (
    "context"
    "fmt"
    pb "github.com/wanmei002/grpc-learn/ch05/server/product"
    "google.golang.org/grpc"
    "log"
)

func main(){
    d, err := grpc.Dial(":8093", grpc.WithInsecure())
    if err != nil {
        log.Println("grpc dial failed; err:", err)
        return
    }
    
    client := pb.NewProductClient(d)
    res, err := client.AddProduct(context.Background(), &pb.Order{
        Name:  "zzh",
        Items: []string{"hello", "zzh"},
        Desc:  "this is zzh",
    })
    
    if err != nil {
        log.Println("addProduct ret err; err:", err)
    }
    
    log.Printf("addProduct return res:%+v", res)
    fmt.Println(res.Code)
}
