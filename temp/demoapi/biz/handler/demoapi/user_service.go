// Code generated by hertz generator.

package demoapi

import (
	"context"
	"easy-note/cmd/api/model/user"
	"easy-note/kitex_gen/demouser"
	"easy-note/pkg/errno"
	"easy-note/temp/demoapi/biz/rpc"

	"github.com/cloudwego/hertz/pkg/app"
)

// CreateUser .
// @router /v2/user/register [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	err = rpc.CreateUser(context.Background(), &demouser.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
