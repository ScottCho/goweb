package main

import (
	"flag"
	"fmt"

	"github.com/ScottCho/goweb/apps/keystone/rpc/internal/config"
	"github.com/ScottCho/goweb/apps/keystone/rpc/internal/server"
	"github.com/ScottCho/goweb/apps/keystone/rpc/internal/svc"
	"github.com/ScottCho/goweb/apps/keystone/rpc/pb/keystone"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/keystone.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		keystone.RegisterKeystoneServer(grpcServer, server.NewKeystoneServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
