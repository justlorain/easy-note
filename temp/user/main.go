package main

import (
	demouser "easy-note/temp/user/kitex_gen/demouser/userservice"
	"log"
)

func main() {
	svr := demouser.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
