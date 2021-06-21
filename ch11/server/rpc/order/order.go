package order

import (
    "context"
    "github.com/wanmei002/grpc-learn/ch11/rpc/pb"
    "sync"
)

type server struct {

}

func NewOrderServer () *server {
    return &server{}
}



var orderDB = map[string]pb.OrderInfo{}
var dbLock sync.Mutex

func (srv *server) AddOrder(ctx context.Context, in *pb.OrderInfo) ( *pb.Res, error) {
    dbLock.Lock()
    defer dbLock.Unlock()
    ret := &pb.Res{Code: 0, Msg: "success"}
    if _, ok := orderDB[in.Name]; ok {
        ret.Msg = "订单已存在"
        return ret, nil
    }
    
    orderDB[in.Name] = *in
    
    return ret, nil
    
}