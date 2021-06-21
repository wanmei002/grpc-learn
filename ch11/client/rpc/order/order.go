package order

import (
    "context"
    "github.com/wanmei002/grpc-learn/ch11/client/comm"
    "github.com/wanmei002/grpc-learn/ch11/rpc/pb"
)

func AddOrder(in *pb.OrderInfo, ctx context.Context) (pb.Res, error) {
    
    client := pb.NewOrderClient(comm.Dl)
    res, err := client.AddOrder(ctx, in)
    if err != nil {
        return pb.Res{}, err
    }
    return *res, err
}

func CloseConn() error {
    return comm.Close()
}