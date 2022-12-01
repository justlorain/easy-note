package main

import (
	"easy-note/cmd/note/dal"
	"easy-note/cmd/note/rpc"
	"easy-note/kitex_gen/demonote/noteservice"
	"easy-note/pkg/consts"
	"easy-note/pkg/middleware"
	"easy-note/pkg/otel"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/registry"
	"net"
)

func Init() {
	rpc.Init()
	dal.Init()
	otel.Init(consts.NoteServiceName)
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

	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(consts.NoteServiceName),
	//	provider.WithExportEndpoint(consts.ExportEndpoint),
	//	provider.WithInsecure(),
	//)
	//defer p.Shutdown(context.Background())

	svr := noteservice.NewServer(new(NoteServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.NoteServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
