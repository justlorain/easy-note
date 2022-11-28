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
	"github.com/kitex-contrib/registry-nacos/resolver"
	"time"
)

var userClient userservice.Client

func initUserRPC() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(r),
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
