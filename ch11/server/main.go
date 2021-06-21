package main

import (
    "crypto/tls"
    "github.com/wanmei002/grpc-learn/ch11/rpc/pb"
    "github.com/wanmei002/grpc-learn/ch11/server/rpc/order"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "log"
    "net"
)

var (
    crtFile = "/usr/local/https-cart/server.crt"
    keyFile = "/usr/local/https-cart/server.key"
)

func main(){
    cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
    if err != nil {
        log.Println("tls explain file failed; err:", err)
        return
    }
    
    opts := []grpc.ServerOption{
        grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
    }
    
    
    ls, err := net.Listen("tcp", ":8093")
    if err != nil {
        log.Println("listen failed; err:", err)
        return
    }
    
    defer ls.Close()
    
    s := grpc.NewServer(opts...)
    pb.RegisterOrderServer(s, order.NewOrderServer())
    log.Println("grpc start server port:8093")
    if err = s.Serve(ls); err != nil {
        log.Println("grpc server start failed; err:", err)
        return
    }
}
