// 本包保存的是 client 拦截器

package interceptor

import (
    "context"
    "google.golang.org/grpc"
    "log"
    "runtime"
)

// 一元拦截器
func ClientUnaryInterceptor (
    ctx context.Context, method string, req, reply interface{},
    cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opt ...grpc.CallOption,
    ) error {
    log.Printf("method[%v];req[%v];reply[%v];cc[%+v];invoker[%T];\n",
                            method, req, reply, cc, invoker)
    // 上面都是前置逻辑操作
    err := invoker(ctx, method, req, reply, cc, opt...)
    // 下面都是后置逻辑操作
    if err != nil {
        log.Println("invoker err:", err)
    }
    log.Println("interceptor after")
    return err
}

// 流拦截器
type clientStream struct {
    grpc.ClientStream
}

func (c *clientStream) RecvMsg(m interface{}) error {
    log.Printf("recv msg : [%T], val:[%v]\n", m , m)
    return c.ClientStream.RecvMsg(m)
}

func (c *clientStream) SendMsg(m interface{}) error {
    log.Printf("send msg:[%T]; val:[%v]\n", m, m)
    return c.ClientStream.SendMsg(m)
}

func NewClientStream(c grpc.ClientStream) grpc.ClientStream {
    return &clientStream{c}
}

// 流拦截器的实现
func ClientStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
    method string, streamer grpc.Streamer, opt ...grpc.CallOption) (grpc.ClientStream, error) {
    pc, _, _, _ := runtime.Caller(0)
    log.Printf("call method[%v];stream desc[%+v]; method[%v]; streamer[%T];\n",
        runtime.FuncForPC(pc).Name(), desc, method, streamer)
    // 创建客户端流
    s, err := streamer(ctx, desc, cc, method, opt...)
    if err != nil {
        log.Println("client stream create failed; err:", err)
        return nil, err
    }
    // 流的发送和接收
    return NewClientStream(s), nil
}








func init() {
    log.SetPrefix("client : ")
}