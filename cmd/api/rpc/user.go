package rpc

import (
	"context"
	"easy-note/kitex_gen/demouser"
	"easy-note/kitex_gen/demouser/userservice"
	"easy-note/pkg/consts"
	"easy-note/pkg/errno"
	"easy-note/pkg/middleware"
	register "easy-note/pkg/registry"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"time"
)

var userClient userservice.Client

func initUserRPC() {
	cli, err := register.NewNacosRegistryCli()
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// CreateUser create user info
func CreateUser(ctx context.Context, req *demouser.CreateUserRequest) error {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// CheckUser check user info
func CheckUser(ctx context.Context, req *demouser.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}