package main

import (
    "fmt"
    "github.com/wanmei002/grpc-learn/04/server/product"
    "google.golang.org/grpc"
    "io"
    "log"
    "net"
    "sync"
    "time"
)

type server struct {

}

var (
    // 订单库
	AllOrder = map[string]*product.Order{}
    // 要返回的信息
    RetOrderInfo []*product.Ret
	// 管道 限制起的 goroutine 太多
    ch chan struct{}
	// 订单id
	id int64
	lock = &sync.Mutex{}
	retLock = &sync.Mutex{}
)

func (s *server) AddProduct(stream product.Product_AddProductServer) error {
    i := 0
    log.Println("start AddProduct")
    // 开启一个协程 专门发送数据
    go sendDataToClient(stream)
    for {
        od, err := stream.Recv()
        if err != nil {
            if err == io.EOF {
                log.Println("server stream end")
                
                retLock.Lock()
                for _, v := range RetOrderInfo {
                    err = stream.Send(v)
                    if err != nil {
                        log.Println("server send failed(EOF);err:", err)
                        return err
                    }
                }
                retLock.Unlock()
                
                log.Println("server send over")
                return nil
                
            }
            
            log.Println("server recv failed; err:", err)
            return err
        }
        log.Println("server ch before")
        <-ch
        i++
        log.Println("server go process data", i)
        log.Println("server begin sleep")
        time.Sleep(1e9)
        log.Println("server sleep over")
        // 处理接收的数据的数据
        go processData(od)
    }
}

func sendDataToClient(stream product.Product_AddProductServer) {
    log.Println("server in send to client")
    for {
        retLock.Lock()
        
        if len(RetOrderInfo)>0 {
            err := stream.Send(RetOrderInfo[0])
            if err != nil {
            
            }
            if len(RetOrderInfo) == 1 {
                RetOrderInfo = make([]*product.Ret, 0)
            } else {
                RetOrderInfo = RetOrderInfo[1:]
            }
        }
        
        retLock.Unlock()
    }
}

func processData(od *product.Order) {
    log.Printf("server recv data:%+v\n", od)
    lock.Lock()
    if _, ok := AllOrder[od.Name]; ok {
        return
    }
    AllOrder[od.Name] = od
    lock.Unlock()
    
    id++
    ret := &product.Ret{
        Code: int32(id),
        Msg: "success",
        OrderId: id,
    }
    retLock.Lock()
    RetOrderInfo = append(RetOrderInfo, ret)
    retLock.Unlock()
    ch <- struct{}{}
}


func main() {
    ch = make(chan struct{}, 10)
    for i:=0; i<10; i++ {
        ch <- struct{}{}
    }
    fmt.Println("server ch:", ch)
    ls, err := net.Listen("tcp", ":8093")
    if err != nil {
        log.Println("listen failed; err:", err)
        return
    }
    g := grpc.NewServer()
    product.RegisterProductServer(g, &server{})
    
    log.Println("start server")
    
    if err = g.Serve(ls); err != nil {
        log.Println("server serve failed; err:", err)
    }
}
