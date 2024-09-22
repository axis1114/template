package core

import (
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"template/global"
	"time"
)

func InitEtcd() *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		global.Log.Fatal("init etcd client fail", zap.Error(err))
	}
	return client
}
