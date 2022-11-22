package main

import (
	"easy-note/cmd/note"
	notedemo "easy-note/cmd/note/kitex_gen/notedemo/noteservice"
	"log"
)

func main() {
	svr := notedemo.NewServer(new(note.NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
