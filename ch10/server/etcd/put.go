package etcd

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var etcdPort = ":2379"

func Put(key, val string) error {
	etcdCli, err := clientv3.New(
		clientv3.Config{
			Endpoints:   []string{etcdPort},
			DialTimeout: 5 * time.Second,
		},
	)

	if err != nil {
		log.Println("etcd client failed; err:", err)
		return err
	}

	res, err := etcdCli.Put(context.Background(), key, val)
	if err != nil {
		log.Printf("etcd put failed; err:[%v];key:[%v];val:[%v];\n", err, key, val)
		return err
	}

	log.Printf("PUT RES :[%+v]\n", res)
	return nil
}
