package main

import (
	demonote "easy-note/temp/demonote/kitex_gen/demonote/noteservice"
	"log"
)

func main() {
	svr := demonote.NewServer(new(NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
