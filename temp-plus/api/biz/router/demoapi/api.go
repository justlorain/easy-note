// Code generated by hertz generator. DO NOT EDIT.

package Demoapi

import (
	demoapi "easy-note/temp-plus/api/biz/handler/demoapi"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_v2 := root.Group("/v2", _v2Mw()...)
		_v2.POST("/note", append(_noteMw(), demoapi.CreateNote)...)
		_note := _v2.Group("/note", _noteMw()...)
		_note.PUT("/:note_id", append(_updatenoteMw(), demoapi.UpdateNote)...)
		_note.DELETE("/:note_id", append(_deletenoteMw(), demoapi.DeleteNote)...)
		_note.GET("/query", append(_querynoteMw(), demoapi.QueryNote)...)
		{
			_user := _v2.Group("/user", _userMw()...)
			_user.POST("/register", append(_createuserMw(), demoapi.CreateUser)...)
		}
	}
}
