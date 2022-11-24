package rpc

import (
	"context"
	"easy-note/kitex_gen/userdemo"
	demouser "easy-note/kitex_gen/userdemo/userservice"
	"easy-note/pkg/consts"
	"easy-note/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"time"
)

var userClient demouser.Client

func initUserRpc() {
	cli, err := newNacosRegistryCli()
	if err != nil {
		panic(err)
	}
	c, err := demouser.NewClient(
		consts.UserServiceName,
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
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

// MGetUser multiple get list of user info
func MGetUser(ctx context.Context, req *userdemo.MGetUserRequest) (map[int64]*userdemo.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := make(map[int64]*userdemo.User)
	for _, u := range resp.Users {
		res[u.UserId] = u
	}
	return res, nil
}
