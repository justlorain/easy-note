// Code generated by hertz generator.

package handler

import (
	"context"
	"easy-note/cmd/api/model"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreateNote .
// @router /v2/note [POST]
func CreateNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.CreateNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(model.CreateNoteResponse)

	c.JSON(200, resp)
}

// DeleteNote .
// @router /v2/note/:note_id [DELETE]
func DeleteNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.DeleteNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(model.DeleteNoteResponse)

	c.JSON(200, resp)
}

// UpdateNote .
// @router /v2/note/:note_id [PUT]
func UpdateNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.UpdateNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(model.UpdateNoteResponse)

	c.JSON(200, resp)
}

// QueryNote .
// @router /v2/note/query [GET]
func QueryNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.QueryNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(model.QueryNoteResponse)

	c.JSON(200, resp)
}