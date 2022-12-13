// Code generated by hertz generator.

package apidemo

import (
	"context"

	apidemo "easy-note/temp/demoapi/biz/model/apidemo"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreateUser .
// @router /v2/user/register [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apidemo.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(apidemo.CreateUserResponse)

	c.JSON(200, resp)
}
