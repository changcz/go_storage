package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"strings"
	"time"
)

// Resolver for grpc client
type Resolver struct {
	schema      string
	EtcdAddrs   []string
	DialTimeout int //超时时间

	closeCh      chan struct{}
	watchCh      clientv3.WatchChan
	cli          *clientv3.Client
	keyPrifix    string
	srvAddrsList []resolver.Address

	cc     resolver.ClientConn
	logger *logrus.Logger
}

// NewRegister 基于etcd new 一个注册对象
func NewRegister(etcdAddrs []string, logger *logrus.Logger) *Register {
	return &Register{
		EtcdAddrs:   etcdAddrs,
		DialTimeout: 3,
		logger:      logger,
	}
}

// Register a service  注册一个服务
func (r *Register) Register(srvInfo Server, ttl int64) (chan<- struct{}, error) {
	var err error

	if strings.Split(srvInfo.Addr, ":")[0] == "" {
		return nil, errors.New("invalid ip address")
	}

	// 新建一个链接
	if r.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdAddrs,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	}); err != nil {
		return nil, err
	}

	r.srvInfo = srvInfo
	r.srvTTL = ttl

	if err = r.register(); err != nil {
		return nil, err
	}

	r.closeCh = make(chan struct{})

	go r.keepAlive()

	return r.closeCh, nil
}

// 创建etcd 自带的实例
func (r *Register) register() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.DialTimeout)*time.Second)
	defer cancel()

	leaseResp, err := r.cli.Grant(ctx, r.srvTTL)
	if err != nil {
		return err
	}

	r.leasesID = leaseResp.ID

	if r.keepAliveCh, err = r.cli.KeepAlive(context.Background(), r.leasesID); err != nil {
		return err
	}

	data, err := json.Marshal(r.srvInfo)
	if err != nil {
		return err
	}

	// push 到etcd  节点注册
	_, err = r.cli.Put(context.Background(), BuildRegisterPath(r.srvInfo), string(data), clientv3.WithLease(r.leasesID))

	return err
}

func (r *Register) keepAlive() {
	ticker := time.NewTicker(time.Duration(r.srvTTL) * time.Second)

	for {
		select {
		case <-r.closeCh:
			//检测是否关闭
			if err := r.unregister(); err != nil {
				r.logger.Error("unregister failed, error: ", err)
			}

			//废除租约
			if _, err := r.cli.Revoke(context.Background(), r.leasesID); err != nil {
				r.logger.Error("revoke failed, error: ", err)
			}
		case res := <-r.keepAliveCh:
			if res == nil {
				if err := r.register(); err != nil {
					r.logger.Error("register failed, error: ", err)
				}
			}
		case <-ticker.C:
			// 超时器
			if r.keepAliveCh == nil {
				if err := r.register(); err != nil {
					r.logger.Error("register failed, error: ", err)
				}
			}
		}
	}
}

// unregister 删除节点
func (r *Register) unregister() error {
	_, err := r.cli.Delete(context.Background(), BuildRegisterPath(r.srvInfo))
	return err
}