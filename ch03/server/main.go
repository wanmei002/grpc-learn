package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/wanmei002/grpc-learn/ch03/server/product"
	"google.golang.org/grpc"
)

var (
	AllOrderInfo = []string{"hello zzh", "hello zyn", "zzh zyn", "world"}
	AllOrderList []product.Order
)

type server struct {
}

func (s *server) SearchGoods(stream product.Product_SearchGoodsServer) error {

	log.Println("start search goods")
	var orderList product.OrderList
	for {
		orderKey, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// 完成读取查询信息流
				log.Println("recv over")
				stream.SendAndClose(&orderList)
				return nil
			}
			log.Println("接收信息失败")
			return err
		}
		// 根据关键字 查询订单里是否有这个
		time.Sleep(1e7)
		log.Printf("recv 到的消息:%+v\n", orderKey)
		for _, order := range AllOrderList {
			for _, item := range order.Items {
				if strings.Contains(item, orderKey.Value) {
					orderList.List = append(orderList.List, &order)
					break
				}
			}
		}
	}
}

func Rand() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(4)

}

func main() {
	// 生成订单列表
	for i := 0; i < 15; i++ {
		d := product.Order{
			Id:    strconv.Itoa(i + 1),
			Items: []string{AllOrderInfo[Rand()], AllOrderInfo[Rand()]},
			Desc:  fmt.Sprint("desc", i+1),
		}
		AllOrderList = append(AllOrderList, d)
	}

	ls, err := net.Listen("tcp", ":8093")
	if err != nil {
		log.Println("server listen failed; err:", err)
		return
	}
	g := grpc.NewServer()
	product.RegisterProductServer(g, &server{})
	fmt.Println("start server")
	if err = g.Serve(ls); err != nil {
		log.Println("server grpc serve failed; err:", err)
		return
	}
}
