package main

import (
    "context"
    "fmt"
    "github.com/wanmei002/grpc-learn/ch05/server/interceptor"
    pb "github.com/wanmei002/grpc-learn/ch05/server/product"
    "google.golang.org/grpc"
    "log"
    "net"
    "reflect"
    "time"
)

var OrderDB = map[string]*pb.Order{}

var ProductComment = map[string][]*pb.CommentInfo{
    "1234567" : []*pb.CommentInfo{
        &pb.CommentInfo{
            UserName:    "zzh",
            UserComment: "你好棒哦",
        },
        &pb.CommentInfo{
            UserName:    "zyn",
            UserComment: "你最棒了",
        },
        &pb.CommentInfo{
            UserName: "gly",
            UserComment: "你需要努力了",
        },
    },
}

// 实现proto声明的方法
type server struct {

}

func (s *server) AddProduct(ctx context.Context, od *pb.Order) (*pb.RespRes, error) {
    if _, ok := OrderDB[od.Name]; ok {
        return &pb.RespRes{
            Code: 0,
            Msg: "order existing",
            Data: "",
        }, nil
    }
    
    OrderDB[od.Name] = od
    return &pb.RespRes{
        Code: 0,
        Msg: "success",
        Data: "",
    }, nil
}

// 根据 product id 查询用户的对它的评论，如果这个product 有评论了，就返回数据
func (s *server) CommentProduct(orderId *pb.ProductId, stream pb.Product_CommentProductServer) error {
    sendIndex := 0
    for {
        if commentList, ok := ProductComment[orderId.Id]; ok {
            if len(commentList) > sendIndex {
                sendErr := stream.Send(commentList[sendIndex])
                if sendErr != nil {
                
                }
                sendIndex++
                time.Sleep(1e9)
            } else {
                return nil
            }
        } else {
            return nil
        }
    }
    log.Println("commentProduct over")
    return nil
}

// 一元拦截器
func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler) (resp interface{}, err error) {
    // 前置处理逻辑
    typeO := reflect.TypeOf(req)
    log.Printf("typeof name:%v, type of kind:%v;\n", typeO.Name(), typeO.Kind())
    if _, ok := req.(*pb.Order); ok {
        fmt.Println("req belong order type")
    }
    
    m, err := handler(ctx, req)
    if err != nil {
        log.Panicln("handler 处理的返回error:", err)
    }
    log.Printf("hander ret:%+v\n", m)
    return m, err
}



func main(){
    port := ":8093"
    ls, err := net.Listen("tcp", port)
    if err != nil {
        log.Println("listen failed; err:", err)
        return
    }
    g := grpc.NewServer(grpc.UnaryInterceptor(orderUnaryServerInterceptor), grpc.StreamInterceptor(interceptor.ProductServerStreamInterceptor))
    pb.RegisterProductServer(g, &server{})
    log.Println("server start, port :", port)
    if err = g.Serve(ls); err != nil {
        log.Println("grpc serve failed; err:", err)
    }
}
