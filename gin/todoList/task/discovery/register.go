package discovery

import (
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Register struct {
	EtcdAddrs   []string
	DialTimeout int

	closeCh     chan struct{} //是否关闭
	leasesID    clientv3.LeaseID // 租约id
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse // 心跳检验

	srvInfo Server
	srvTTL  int64
	cli     *clientv3.Client // 客户端
	logger  *logrus.Logger
}


// Stop stop register
func (r *Register) Stop() {
	r.closeCh <- struct{}{}
}