package main

import (
    "context"
    "fmt"
    "github.com/wanmei002/grpc-learn/ch07/server/chat"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "io"
    "log"
    "os"
    "time"
)

func main(){
    input := os.Args
    say := "你好"
    if len(input) >= 2 {
        say = input[1]
    }
    d, err := grpc.Dial(":8093", grpc.WithInsecure())
    if err != nil {
        log.Println("grpc dial failed; err:", err)
        return
    }
    defer d.Close()
    
    client := chat.NewChatSvrClient(d)
    ctx, cancel := context.WithCancel(context.Background())
    clientStream, err := client.SendMessage(ctx)
    if err != nil {
        log.Println("get client stream failed; err:", err)
        return
    }
    msg := &chat.MsgInfo{
        ChatRoomID:  1,
        UserID:      4567,
        UserName:    "zyn",
        UserHeadImg: "love.jpg",
        Msg:         say,
        Ext:         "",
    }
    fmt.Println("zyn: ", say)
    err = clientStream.Send(msg)
    if err != nil {
        log.Println("client send data failed; err:", err)
        return
    }
    sendMsgTime := new(int64)
    
    *sendMsgTime = time.Now().Unix()
    // 起一个 goroutine 来接收客户端发送的数据
    go getRecvMsg(clientStream)
    ch := make(chan struct{})
    // 起一个 goroutine 来检查是否长时间未通话，未通话关闭连接
    go heartBeat(sendMsgTime, ch)
    // 起一个 goroutine 来发送数据
    go sendMsg(clientStream, ctx, sendMsgTime)
    
    // 取消阻塞
    <- ch
    // 长时间未通话，关闭 grpc 连接
    cancel()
    time.Sleep(1e9)
    log.Println("client end")
    
}

// 客户端发送数据
func sendMsg(clientStream chat.ChatSvr_SendMessageClient, ctx context.Context, sendMsgTime *int64) {
    for {
        select {
        case <-ctx.Done(): // 检查是否关闭了连接
            return
        default:
            err := client2svr(clientStream, sendMsgTime)
            if err != nil {
                log.Println("send msg failed; err:", err)
                return
            }
        }
    }
}

func client2svr(clientStream chat.ChatSvr_SendMessageClient, sendMsgTime *int64) error {
    say := ""
    fmt.Scanf("%s\n", &say)
    if say == "over" {
        clientStream.CloseSend()
        fmt.Printf("    通话结束    \n")
        return nil
    }
    
    msg := &chat.MsgInfo{
        ChatRoomID:  1,
        UserID:      4567,
        UserName:    "zyn",
        UserHeadImg: "love.jpg",
        Msg:         say,
        Ext:         "",
    }
    
    msg.Msg = say
    err := clientStream.Send(msg)
    if err != nil {
        log.Println("send msg failed; err:", err)
        return err
    }
    *sendMsgTime = time.Now().Unix()
    return nil
}
// 检查多长时间未通话了
func heartBeat(t *int64, ch chan struct{}) {
    defer func(){
        ch <- struct{}{}
    }()
    for {
        time.Sleep(time.Second)
        nowTime := time.Now().Unix()
        sub := nowTime - *t
        // 超过 3 秒 关闭连接
        if sub > 3 {
            log.Println("长时间没有通话了[自动关闭连接]")
            return
        }
    }
}


func getRecvMsg(client chat.ChatSvr_SendMessageClient) {
    for {
        msg, err := client.Recv()
        if err != nil {
            if err == io.EOF {
               fmt.Println("END")
                return
            } else if status.Code(err) == codes.Canceled {// 获取 grpc 错误码状态，看是不是对方关闭了连接
                log.Println("服务端已经关闭了请求")
                return
            }
            log.Println("client recv failed; err:", err)
            return
        }
        fmt.Println("        ", msg.Msg, " :", msg.UserName)
    }
}
