package main

import (
	"easy-note/cmd/api/router/note"
	"easy-note/cmd/api/router/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers all routers.
func register(r *server.Hertz) {

	user.Register(r)
	note.Register(r)

}
