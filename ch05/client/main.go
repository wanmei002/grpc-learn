package main

import (
    "context"
    "fmt"
    "github.com/wanmei002/grpc-learn/ch05/client/interceptor"
    pb "github.com/wanmei002/grpc-learn/ch05/server/product"
    "google.golang.org/grpc"
    "log"
)

func main(){
    // 在这里注册拦截器
    d, err := grpc.Dial(":8093", grpc.WithInsecure(),
        grpc.WithUnaryInterceptor(interceptor.ClientUnaryInterceptor),
        grpc.WithStreamInterceptor(interceptor.ClientStreamInterceptor))
    
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
            log.Println("once return")
            return
        }
        
    }
    
    log.Println("client end")
}
