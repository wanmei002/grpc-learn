package etcd

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// 需要注册服务到 etcd 并设置租约，定期续租，如果服务停止了则租约过期，
// 就会从 etcd 服务中消失

type RegisterEtcdServer struct {
	etcdCli *clientv3.Client
	leaseId clientv3.LeaseID
	ctx     context.Context
}

func RegisterServer(k, v string, expire int64) (*RegisterEtcdServer, error) {
	etcdClient, err := clientv3.New(
		clientv3.Config{
			Endpoints:   []string{":2379"},
			DialTimeout: 5 * time.Second,
		},
	)

	if err != nil {
		log.Println("new etcd client failed; err:", err)
		return nil, err
	}

	svr := &RegisterEtcdServer{
		etcdCli: etcdClient,
		ctx:     context.Background(),
	}

	// 开始运行租约等逻辑
	err = svr.createLease(expire)
	if err != nil {
		log.Println("创建租约失败 err:", err)
		return nil, err
	}

	err = svr.BindLease(k, v)
	if err != nil {
		log.Println("etcd put failed; err:", err)
		return nil, err
	}

	// 定时续期
	err = svr.keepAlive()
	if err != nil {
		log.Println("定时续期失败l err:", err)
		return nil, err
	}

	return svr, nil
}

// 1. 创建 etcd 客户端
// 2. 创建租约
// 3. k v 租约绑定
// 4. 定期续期

func (s *RegisterEtcdServer) createLease(expire int64) error {
	res, err := s.etcdCli.Grant(s.ctx, expire)
	if err != nil {
		log.Println("create grant failed; err : ", err)
		return err
	}
	s.leaseId = res.ID
	return nil
}

func (s *RegisterEtcdServer) BindLease(k, v string) error {
	res, err := s.etcdCli.Put(s.ctx, k, v, clientv3.WithLease(s.leaseId))
	if err != nil {
		log.Println("etcd put failed; err:", err)
		return err
	}

	log.Printf("etcd put succ : %+v\n", res)
	return nil
}

func (s *RegisterEtcdServer) keepAlive() error {
	leaseResCh, err := s.etcdCli.KeepAlive(s.ctx, s.leaseId)
	if err != nil {
		log.Println("etcd keep live failed; err :", err)
		return err
	}

	go s.watch(leaseResCh)
	return nil
}

func (s *RegisterEtcdServer) watch(leaseCh <-chan *clientv3.LeaseKeepAliveResponse) {
	for k := range leaseCh {
		log.Printf("续约成功; val:%+v\n", k)
	}

	log.Println("租约续期关闭")
}

func (s *RegisterEtcdServer) Close() error {
	// 撤销租约
	res, err := s.etcdCli.Revoke(s.ctx, s.leaseId)
	if err != nil {
		log.Println("撤销租约失败")
		return err
	}

	log.Printf("撤销租约返回的结果: %+v\n", res)

	return s.etcdCli.Close()
}
