package interceptor

import (
    "google.golang.org/grpc"
    "log"
)

// 服务端流拦截器
type serverStreamInterceptor struct {
    grpc.ServerStream
}

func (si *serverStreamInterceptor) RecvMsg(m interface{}) error {
    log.Printf("=====[server stream interceptor recv msg] " +
        "rece msg type[%T]=====\n", m)
    return si.ServerStream.RecvMsg(m)
}

func (si *serverStreamInterceptor) SendMsg(m interface{}) error {
    log.Printf("===== server stream interceptor send msg " +
        "send msg type[%T]=====\n", m)
    return si.ServerStream.SendMsg(m)
}

func NewServerStreamInterceptor(s grpc.ServerStream) grpc.ServerStream {
    return &serverStreamInterceptor{s}
}

func ProductServerStreamInterceptor(srv interface{}, ss grpc.ServerStream,
    info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
    log.Printf("server stream interceptor , get data type[%T]; value : [%v]\n", srv, srv)
    
    err := handler(srv, NewServerStreamInterceptor(ss))
    
    if err != nil {
        log.Println("server stream handler err:", err)
    }
    
    return err
}

