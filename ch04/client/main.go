package main

import (
    "context"
    "fmt"
    "github.com/wanmei002/grpc-learn/04/server/product"
    "google.golang.org/grpc"
    "log"
)

func main(){
    d, err := grpc.Dial(":8093", grpc.WithInsecure())
    if err != nil {
        log.Println("client dial failed; err:", err)
        return
    }
    defer d.Close()
    
    client := product.NewProductClient(d)
    
    stream, err := client.AddProduct(context.Background())
    if err != nil {
        log.Println("client add product failed; err:", err)
        return
    }
    // 来一个协程专门接收数据
    ctx, closeFunc := context.WithCancel(context.Background())
    go RecvData(stream, closeFunc)
    for i:=0; i<100; i++ {
        o := &product.Order{
            Name: fmt.Sprint("zzh-",i),
            Desc: fmt.Sprint("desc-",i),
            Items: make([]string, 1),
        }
        log.Println("client send data:", o)
        err := stream.Send(o)
        if err != nil {
            log.Println("clinet send failed;err:", err)
        }
    }
    
    err = stream.CloseSend()
    <-ctx.Done()
    fmt.Println("over")
    
}

func RecvData(stream product.Product_AddProductClient, cancelFunc context.CancelFunc) {
    defer cancelFunc()
    for {
        res, err := stream.Recv()
        if err != nil {
            log.Println("client recv failed; err:", err)
            return
        }
        log.Printf("client recv %+v\n", res)
    }
    
    
}
