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
    defer d.Close()
    
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
    
    stream, err := client.CommentProduct(context.Background(), &pb.ProductId{Id: "1234567"})
    if err != nil {
        log.Println("client comment product failed; err:", err)
    }
    log.Println("start recv server stream msg")
    if stream != nil {
        for {
            commentInfo, err := stream.Recv()
            if err != nil {
                log.Println("client stream recv failed; err:", err)
                break
            }
            log.Printf("recv data :[%+v]\n", commentInfo)
        }
        
    }
    
    log.Println("client end")
}
