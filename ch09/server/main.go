package main

import (
    "context"
    pb "github.com/wanmei002/grpc-learn/ch09/server/order"
    "google.golang.org/grpc"
    "log"
    "net"
    "sync"
)

type server struct {
    addr string
}

var (
	OrderDB = map[string]pb.OrderInfo{}
    orderLock sync.Mutex
	addrList = []string{":8093", ":8094"}
)


func (s *server) AddOrder(ctx context.Context, order *pb.OrderInfo) (*pb.Res, error) {
    defer orderLock.Unlock()
    orderLock.Lock()
    log.Println("server addr :", s.addr)
    if _, ok := OrderDB[order.Name]; ok {
        return &pb.Res{Code: 0, Msg: "订单已存在"}, nil
    }
    
    OrderDB[order.Name] = *order
    return &pb.Res{Code: 0, Msg: "添加成功"}, nil
    
}

func StartServer(addr string, group *sync.WaitGroup) {
    defer group.Done()
    log.Println("server addr :", addr)
    ls, err := net.Listen("tcp", addr)
    if err != nil {
        log.Println("listen failed; err:", err, "; addr:", addr)
        return
    }
    
    gSvr := grpc.NewServer()
    pb.RegisterOrderServer(gSvr, &server{addr: addr})
    
    log.Println("start server")
    if err = gSvr.Serve(ls); err != nil {
        log.Println("gSvr failed;err:", err)
        return
    }
    
}



func main(){
    log.Println("loop register server")
    var wg = &sync.WaitGroup{}
    wg.Add(len(addrList))
    for _, addr := range addrList {
        go StartServer(addr, wg)
    }
    
    wg.Wait()
    log.Println("over")
}
