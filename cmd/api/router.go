package main

import (
	"context"
	"easy-note/cmd/api/router"
	"easy-note/cmd/api/router/note"
	"easy-note/cmd/api/router/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// register registers all routers.
func register(r *server.Hertz) {

	r.POST("/v1/user/login", router.JwtMiddleware.LoginHandler)
	user.Register(r)
	note.Register(r)
	r.NoRoute(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 404
		c.String(consts.StatusOK, "no route")
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 405
		c.String(consts.StatusOK, "no method")
	})

}
