package main

import (
    "github.com/wanmei002/grpc-learn/ch07/server/chat"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "io"
    "log"
    "net"
)

type server struct {

}

func (s *server) SendMessage(recvStream chat.ChatSvr_SendMessageServer) error {
    // for 循环接收客户端发送的信息
    for {
        msgInfo, err := recvStream.Recv()
        if err != nil {
            if err == io.EOF {
                log.Println("客户端请求结束发送")
                return nil
            } else if status.Code(err) == codes.Canceled {// 获取错误码状态，看是否是客户端关闭了连接
                log.Println("客户端取消了连接")
                // 返回给客户端 grpc 错误码
                return status.Error(codes.Canceled, "server closed")
            }
            
            
            log.Println("recv failed; err:", err)
            return err
        }
        sendMsg := &chat.MsgInfo{
            ChatRoomID:  msgInfo.ChatRoomID,
            UserID:      123,
            UserName:    "zzh",
            UserHeadImg: "img.jpg",
            Msg:         "",
            Ext:         "",
        }
        switch msgInfo.Msg {
        case "你好":
            sendMsg.Msg = "你好"
            err = recvStream.Send(sendMsg)
        case "你在哪呢":
            sendMsg.Msg = "在家里呢"
            err = recvStream.Send(sendMsg)
        default:
            sendMsg.Msg = "没听清, 你再说一次"
            err = recvStream.Send(sendMsg)
        }
        if err != nil {
            log.Println("send msg failed; err : ", err)
            return err
        }
    }
}

func main(){
    ls, err := net.Listen("tcp", ":8093")
    if err != nil {
        log.Println("listen failed; err:", err)
        return
    }
    g := grpc.NewServer()
    chat.RegisterChatSvrServer(g, &server{})
    
    log.Println("start server; port:8093")
    
    if err = g.Serve(ls); err != nil {
        log.Println("grpc serve failed; err:", err)
        return
    }
    
}
