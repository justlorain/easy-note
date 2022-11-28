// Code generated by hertz generator.

package main

import (
	"easy-note/cmd/api/router"
	"easy-note/cmd/api/rpc"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	rpc.Init()
	router.InitJWT()
}

func main() {
	Init()
	h := server.New(
		server.WithHostPorts(":8080"),
		server.WithHandleMethodNotAllowed(true), // coordinate with NoMethod
	)
	register(h)
	h.Spin()
}
