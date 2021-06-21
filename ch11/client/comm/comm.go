package comm

import (
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "log"
)

var (
    CrtFile = "/usr/local/https-cart/server.crt"
    Dl *grpc.ClientConn
)

func Close() error {
    return Dl.Close()
}


func init(){
    // www.zzh.com 是证书生成的时候 输入的网址名字
    creds, err := credentials.NewClientTLSFromFile(CrtFile, "www.zzh.com")
    if err != nil {
        log.Println("phrase crtfile failed; err:", err)
        panic(fmt.Sprintf("phrase crt file failed; err:%v\n", err))
    }
    
    opts := []grpc.DialOption{
        grpc.WithTransportCredentials(creds),
    }
    
    
    Dl, err = grpc.Dial(":8093", opts...)
}
