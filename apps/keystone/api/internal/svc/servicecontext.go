package svc

import (
	"github.com/ScottCho/goweb/apps/keystone/api/internal/config"
	"github.com/ScottCho/goweb/apps/keystone/rpc/keystoneclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	KeystoneRpc keystoneclient.Keystone
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		KeystoneRpc: keystoneclient.NewKeystone(zrpc.MustNewClient(c.KeystoneRpc)),
	}
}
