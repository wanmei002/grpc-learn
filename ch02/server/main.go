package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/wanmei002/grpc-learn/ch02/server/product"
	"google.golang.org/grpc"
)

type server struct {
}

var m = []product.Order{}

var lock = &sync.Mutex{}

func (s *server) SearchOrders(searchQuery *wrappers.StringValue,
	stream product.OrderManagement_SearchOrdersServer) error {
	defer lock.Unlock()
	lock.Lock()
	for _, order := range m {
		for _, item := range order.Items {
			if strings.Contains(item, searchQuery.Value) {
				err := stream.Send(&order)
				if err != nil {
					fmt.Println("server stream send failed; err:", err)
					return err
				}
				fmt.Println("fond " + searchQuery.Value)
				break
			}
		}
	}
	return nil
}

var names = []string{"hello zzh", "hello zyn", "zzh zyn", "zyn is all"}

// 产生随机数
func Rand() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	i := r.Intn(3)
	fmt.Println("rand :", i)
	return i

}

func main() {
	for i := 0; i < 15; i++ {
		p := product.Order{
			Id:          strconv.Itoa(i + 1),
			Items:       []string{names[Rand()], names[Rand()]},
			Desc:        "hello " + strconv.Itoa(i),
			Destination: strconv.Itoa(i),
		}
		lock.Lock()
		m = append(m, p)
		lock.Unlock()
	}

	ls, err := net.Listen("tcp", ":8903")
	if err != nil {
		fmt.Println("listen failed; err:", err)
		return
	}

	g := grpc.NewServer()

	product.RegisterOrderManagementServer(g, &server{})

	fmt.Println("start service")

	if err = g.Serve(ls); err != nil {
		fmt.Println("server start failed; err:", err)
	}

}
