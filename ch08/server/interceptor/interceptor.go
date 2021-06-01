package interceptor

import (
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
    "log"
)

func ServerStreamInterceptor(srv interface{}, ss grpc.ServerStream,
    info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
    
    log.Println("stream interceptor get header data")
    log.Printf("stream server info :[%v]\n", info)
    // 获取 header 信息
    md, _ := metadata.FromIncomingContext(ss.Context())
    if _, ok := md["token"]; !ok {
        log.Println("client not trans token")
        return status.Error(codes.InvalidArgument, "token must")
    }
    if md["token"][0] != "1234" {
        log.Printf("client token error; request token[%v]; must token:[%v];\n", md["token"], 123)
        return status.Error(codes.InvalidArgument, "token error")
    }
    log.Printf("get header info :==%v==\n", md)
    
    err := handler(srv, ss)
    
    if err != nil {
        log.Println("server stream handler err:", err)
    }
    
    return err
}
