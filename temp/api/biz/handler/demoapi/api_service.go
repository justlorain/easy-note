// Code generated by hertz generator.

package demoapi

import (
	"context"

	demoapi "easy-note/temp/api/biz/model/demoapi"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreateUser .
// @router /v2/user/register [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(demoapi.CreateUserResponse)

	c.JSON(200, resp)
}

// CheckUser .
// @router /v2/user/login [POST]
func CheckUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.CheckUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(demoapi.CheckUserResponse)

	c.JSON(200, resp)
}

// CreateNote .
// @router /v2/note [POST]
func CreateNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.CreateNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(demoapi.CreateNoteResponse)

	c.JSON(200, resp)
}

// QueryNote .
// @router /v2/note/query [GET]
func QueryNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.QueryNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(demoapi.QueryNoteResponse)

	c.JSON(200, resp)
}

// UpdateNote .
// @router /v2/note/:note_id [PUT]
func UpdateNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.UpdateNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(demoapi.UpdateNoteResponse)

	c.JSON(200, resp)
}

// DeleteNote .
// @router /v2/note/:note_id [DELETE]
func DeleteNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.DeleteNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(demoapi.DeleteNoteResponse)

	c.JSON(200, resp)
}
