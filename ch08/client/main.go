package main

import (
    "context"
    "fmt"
    "github.com/wanmei002/grpc-learn/ch07/client/interceptor"
    "github.com/wanmei002/grpc-learn/ch07/server/chat"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "log"
    "time"
)

func main(){
    d, err := grpc.Dial(":8093", grpc.WithInsecure(),
        grpc.WithStreamInterceptor(interceptor.ClientStreamInterceptor))
    if err != nil {
        log.Println("dial failed; err:", err)
        return
    }
    defer d.Close()
    
    client := chat.NewChatRoomClient(d)
    ctx, cancel := context.WithCancel(context.Background())
    clientStream, err := client.SendMsg(ctx)
    if err != nil {
        log.Println("send msg failed; err:", err)
        return
    }
    log.Println("start client")
    ch := make(chan struct{})
    latestChatTime := new(int64)
    *latestChatTime = time.Now().Unix()
    go recvMsg(clientStream, ctx, ch, latestChatTime)
    go heartBeat(latestChatTime, ch)
    // 发送数据
    go sendMsg(clientStream, ch, latestChatTime)
    <-ch
    cancel()
    log.Println("client rpc over")
    
    
    // 开始接收数据
}

func sendMsg(stream chat.ChatRoom_SendMsgClient, ch chan struct{}, latestChatTime *int64) {
    defer func(){
        ch <- struct{}{}
    }()
    say := ""
    sendMsg := &chat.Msg{UserID:1, UserName:"zyn", Msg:""}
    for {
        fmt.Scanf("zyn: %s\n", &say)
        sendMsg.Msg = say
        *latestChatTime = time.Now().Unix()
        err := stream.Send(sendMsg)
        if err != nil {
            log.Printf("send msg failed;err[%v]; msg[%v];\n", err, sendMsg)
            return
        }
        
    }
}

func heartBeat(latestChatTime *int64, ch chan struct{}) {
    defer func(){
        ch <- struct{}{}
    }()
    for {
        time.Sleep(1e9)
        nowTime := time.Now().Unix()
        if (nowTime - *latestChatTime) > 10 {
            log.Println("long time no chat")
            return
        }
    }
}

func recvMsg(stream chat.ChatRoom_SendMsgClient, ctx context.Context, ch chan struct{}, latestChatTIme *int64) {
    defer func(){
        ch <- struct{}{}
    }()
    
    for {
        select {
        case <-ctx.Done():
            log.Println("client closed")
            stream.CloseSend()
            return
        default:
            // 接收数据给客户端
            *latestChatTIme = time.Now().Unix()
            err := getMsg2Svr(stream)
            if err != nil {
                return
            }
        }
    }
}

func getMsg2Svr(stream chat.ChatRoom_SendMsgClient) error {
    msg, err := stream.Recv()
    if err != nil {
        log.Println("stream recv failed; err:", err)
        if status.Code(err) == codes.Canceled {
            log.Println("server close connect")
        }
        return err
    }
    fmt.Println("          ", msg.Msg, ":", msg.UserName)
    return nil
}
