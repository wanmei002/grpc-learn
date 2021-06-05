// 服务发现包 discovery
package discovery

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"google.golang.org/grpc/resolver"
)

var scheme = "order.Order"

// 1. 我们先写一个结构体，继承 Bulider Resolver 这两个接口
// 2. Bulid 这个方法是gRPC框架调用的入口方法，在这个入口方法里我们实现如下逻辑:
//     2.1 根据前缀从 etcd 中获取服务列表，用 resolver.ClientConn.UpdateState 更新服务列表
//     2.2 起一个监听服务 调用 etcd 的观察者模式，观察 服务列表是否有变化, 如果有变化，调用 UpdateState 来更新服务列表

type serverDiscovery struct {
	cli           *clientv3.Client    // 用来连接etcd
	conn          resolver.ClientConn // 用来 调用 UpdateState 这个方法，更新本地服务ip列表
	serviceIpList sync.Map            // 用来存储获得的ip列表，因为监听服务是新的协程在运行，可能会存在对map的同时读写，引起资源冲突
}

// Bulid 先实现这个接口
func (s *serverDiscovery) Bulid(target resolver.Target, cc resolver.ClientConn,
	opts resolver.BuildOption) (resolver.Resolver, error) {
	// 获取在 etcd 保存的前缀
	prefix := fmt.Sprintf("/%s/%s/", target.Scheme, target.Endpoint)
	res, err := s.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		log.Println("Bulid etcd get addr failed; err:", err)
		return nil, err
	}
	for _, kv := range res.Kvs {
		s.store(kv.Key, kv.Value)
	}

	s.updateState()

	// 启动 etcd 观察者模式

	return s, nil

}

func (s *serverDiscovery) watch(prefix string) {
	res := s.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	// res 是一个只读的管道
	for val := range res {
		for _, event := range val.Events {
			switch event.Type {
			case mvccpb.PUT:
				s.store(event.Kv.Key, event.Kv.Value)
				s.updateState()
			case mvccpb.DELETE:
				s.del(event.Kv.Key)
			}
		}
	}

}

func (s *serverDiscovery) Scheme() string {
	return scheme
}

func (s *serverDiscovery) ResolveNow(resolver.ResolveNowOption) {

}

func (s *serverDiscovery) Close() {
	s.cli.Close()
}

func (s *serverDiscovery) store(k, v []byte) {

	s.serviceIpList.Store(string(k), string(v))

}

func (s *serverDiscovery) del(k []byte) {
	s.serviceIpList.Delete(string(k))
}

func (s *serverDiscovery) updateState() {
	var addrList resolver.State
	s.serviceIpList.Range(func(k, v interface{}) bool {
		tA, ok := v.(string)
		if !ok {
			return false
		}
		log.Printf("conn.UpdateState key[%v];val[%v]\n", k, v)
		addrList.Addresses = append(addrList.Addresses, resolver.Address{Addr: tA})
		return true
	})

	s.conn.UpdateState(addrList)
}
