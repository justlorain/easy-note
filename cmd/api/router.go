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

	user.Register(r)
	note.Register(r)
	r.POST("/v2/user/login", router.JwtMiddleware.LoginHandler)
	r.NoRoute(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 404
		c.String(consts.StatusOK, "no route")
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 405
		c.String(consts.StatusOK, "no method")
	})

}
