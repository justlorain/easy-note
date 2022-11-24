package main

import (
	"easy-note/cmd/note/dal"
	"easy-note/cmd/note/rpc"
	demonote "easy-note/kitex_gen/notedemo/noteservice"
	"easy-note/pkg/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"net"
)

func Init() {
	rpc.Init()
	dal.Init()
}

func main() {
	cli, err := newNacosRegistryCli()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.NoteServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()
	svr := demonote.NewServer(new(NoteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.NoteServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}

func newNacosRegistryCli() (naming_client.INamingClient, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(consts.LocalHost, consts.NacosPort),
	}
	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "info",
		Username:            "nacos",
		Password:            "nacos",
	}
	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	return cli, err
}
