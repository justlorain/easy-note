package main

import (
	notedemo "easy-note/cmd/api/router/note"
	userdemo "easy-note/cmd/api/router/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers all routers.
func register(r *server.Hertz) {

	userdemo.Register(r)
	notedemo.Register(r)

}
