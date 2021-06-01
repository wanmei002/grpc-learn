package interceptor

import (
    "context"
	"google.golang.org/grpc"
    "google.golang.org/grpc/metadata"
    "log"
    "strconv"
    "time"
)

// 客户端的流控制器
func ClientStreamInterceptor (ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
    method string, streamer grpc.Streamer, opt ...grpc.CallOption) (grpc.ClientStream, error) {
    // 增加元数据
    log.Println("start add header data")
    // 往 header 头里保存数据
    headerData := metadata.Pairs("timestamp", strconv.Itoa(int(time.Now().Unix())), "token", "123")
    ctxH := metadata.NewOutgoingContext(ctx, headerData)
    s, err := streamer(ctxH, desc, cc, method, opt...)
    if err != nil {
        log.Println("interceptor streamer error:", err)
    }
    return s, nil
}
