package main

import (
    "github.com/wanmei002/grpc-learn/ch07/server/chat"
    "github.com/wanmei002/grpc-learn/ch07/server/interceptor"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "log"
    "net"
)

type server struct {

}

func (s *server) SendMsg(stream chat.ChatRoom_SendMsgServer) error {
    for {
        msg, err := stream.Recv()
        if err != nil {
            log.Println("server recv failed; err:", err)
            if status.Code(err) == codes.Canceled {
                log.Println("client closed connect")
            }
            return err
        }
        sendMsg := &chat.Msg{UserID: 2,UserName: "zzh",Msg: ""}
        switch msg.Msg {
        case "你好":
            sendMsg.Msg = "你好"
        case "你在哪里啊":
            sendMsg.Msg = "我在家里"
        case "我去找你吧":
            sendMsg.Msg = "好啊"
        default:
            sendMsg.Msg = "我没听清楚，你再说一遍吧"
        }
        
        err = stream.Send(sendMsg)
        if err != nil {
            log.Println("send msg failed; err:", err)
            return err
        }
    }
}


func main(){
    ls, err := net.Listen("tcp", ":8093")
    if err != nil {
        log.Println("")
    }
    gSvr := grpc.NewServer(grpc.StreamInterceptor(interceptor.ServerStreamInterceptor))
    // 把 &server{} 加入到 gSvr.services 属性里
    chat.RegisterChatRoomServer(gSvr, &server{})
    log.Println("start server")
    if err = gSvr.Serve(ls); err != nil {
        log.Println("server listen failed; err:", err)
    }
    
    
    
}
