package rpc

import (
	"context"
	"easy-note/kitex_gen/demouser"
	"easy-note/kitex_gen/demouser/userservice"
	"easy-note/pkg/consts"
	"easy-note/pkg/errno"
	"easy-note/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"time"
)

var userClient userservice.Client

func initUser() {
	//otel.Init(consts.NoteClientName)
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.NoteClientName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	c, err := userservice.NewClient(
		consts.UserServiceName, // DestService
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),                                                 // otel
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.NoteClientName}), // otel
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// MGetUser multiple get list of user info
func MGetUser(ctx context.Context, req *demouser.MGetUserRequest) (map[int64]*demouser.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := make(map[int64]*demouser.User)
	for _, u := range resp.Users {
		res[u.UserId] = u
	}
	return res, nil
}
