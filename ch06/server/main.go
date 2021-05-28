package main

import (
    "context"
    pb "github.com/wanmei002/grpc-learn/ch06/server/product"
    "google.golang.org/grpc"
    "log"
    "net"
    "sync"
    "time"
)

type server struct {

}

var orderDB = map[string]*pb.Order{}
var lock sync.Mutex
// 保存数据

func (s *server) AddOrder(ctx context.Context, order *pb.Order) (*pb.Res, error) {
    log.Println("in add order")
    ch := make(chan struct{}, 1)
    res := &pb.Res{}
    // 启用两个 goroutine 一个用作处理逻辑  一个用作监听是否超时
    go addOrder(order, ch, res)
    svrCtx, cancel := context.WithCancel(context.Background())
    // 开始检测是否取消
    go observeCtx(ctx, svrCtx, res, ch)
    
    <-ch
    cancel()
    log.Println("end")
    return res, nil
    
}

func observeCtx(clientCtx, serverCtx context.Context, res *pb.Res, ch chan struct{}) {
    defer func(){
        ch <- struct{}{}
    }()
    
    for {
        select {
        case <-clientCtx.Done():// 超过了客户端设置的时间
            log.Println("client ctx done")
            return
        case <-serverCtx.Done():// 服务端逻辑执行完毕，停止执行监听的协程
            res.Code = 0
            res.Msg = "success"
            log.Println("server ctx done")
            return
        }
    }
}



// 处理信息
func addOrder(order *pb.Order, ch chan struct{}, res *pb.Res) {
    defer func(){
        res.Code = 0
        res.Msg = "success"
        ch <- struct{}{}
    }()
    
    defer lock.Unlock()
    lock.Lock()
    if _, ok := orderDB[order.Name]; ok {
        return
    }
    log.Println("sleep add order")
    // sleep 1s, 可修改，模拟超时
    time.Sleep(1e9)
    orderDB[order.Name] = order
    return
}

func main(){
    ls, err := net.Listen("tcp", ":8093")
    if err != nil {
        log.Println("listen failed; err:", err)
        return
    }
    g := grpc.NewServer()
    
    pb.RegisterOrderServerServer(g, &server{})
    log.Println("start server")
    if err = g.Serve(ls); err != nil {
        log.Println("g serve failed; err:", err)
        return
    }
    
}
