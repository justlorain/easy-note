package main

import (
	userdemo "easy-note/kitex_gen/userdemo/userservice"
	"log"
)

func main() {
	svr := userdemo.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
