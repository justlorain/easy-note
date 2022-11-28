package main

import (
	"easy-note/cmd/note/dal"
	"easy-note/cmd/note/rpc"
	"easy-note/kitex_gen/demonote/noteservice"
	"easy-note/pkg/consts"
	"easy-note/pkg/middleware"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
)

func Init() {
	rpc.Init()
	dal.Init()
}

func main() {
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.NoteServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()
	svr := noteservice.NewServer(new(NoteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.NoteServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
