package etcd

import (
	"context"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"time"
)

// todo GetEtcdClientWithAuth
func GetEtcdClientNoAuth(addrs []string) (*clientv3.Client,error) {
	c,err := clientv3.New(clientv3.Config{
		Endpoints:   addrs,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil,err
	}
	// 检查etcd各个节点的状态
	success_count := 0
	for _,t := range addrs {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
		_, err = c.Status(timeoutCtx, t)
		cancel()
		if err == nil {
			success_count++
		}
	}
	if (len(addrs) == 1 && success_count < 1) || (len(addrs) > 1 && success_count < len(addrs)/2+1) {
		return nil,errors.New("etcd 健康节点个数小于总节点数量/2+1，集群不成立")
	}
	return c,nil
}