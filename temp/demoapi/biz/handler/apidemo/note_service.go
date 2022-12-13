// Code generated by hertz generator.

package apidemo

import (
	"context"

	apidemo "easy-note/temp/demoapi/biz/model/apidemo"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreateNote .
// @router /v2/note [POST]
func CreateNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apidemo.CreateNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(apidemo.CreateNoteResponse)

	c.JSON(200, resp)
}

// QueryNote .
// @router /v2/note/query [GET]
func QueryNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apidemo.QueryNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(apidemo.QueryNoteResponse)

	c.JSON(200, resp)
}

// UpdateNote .
// @router /v2/note/:note_id [PUT]
func UpdateNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apidemo.UpdateNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(apidemo.UpdateNoteResponse)

	c.JSON(200, resp)
}

// DeleteNote .
// @router /v2/note/:note_id [DELETE]
func DeleteNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apidemo.DeleteNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(apidemo.DeleteNoteResponse)

	c.JSON(200, resp)
}
