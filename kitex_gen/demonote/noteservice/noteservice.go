// Code generated by Kitex v0.4.3. DO NOT EDIT.

package noteservice

import (
	"context"
	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/demonote"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return noteServiceServiceInfo
}

var noteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "NoteService"
	handlerType := (*demonote.NoteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateNote": kitex.NewMethodInfo(createNoteHandler, newNoteServiceCreateNoteArgs, newNoteServiceCreateNoteResult, false),
		"DeleteNote": kitex.NewMethodInfo(deleteNoteHandler, newNoteServiceDeleteNoteArgs, newNoteServiceDeleteNoteResult, false),
		"UpdateNote": kitex.NewMethodInfo(updateNoteHandler, newNoteServiceUpdateNoteArgs, newNoteServiceUpdateNoteResult, false),
		"QueryNote":  kitex.NewMethodInfo(queryNoteHandler, newNoteServiceQueryNoteArgs, newNoteServiceQueryNoteResult, false),
		"MGetNote":   kitex.NewMethodInfo(mGetNoteHandler, newNoteServiceMGetNoteArgs, newNoteServiceMGetNoteResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "demonote",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.3",
		Extra:           extra,
	}
	return svcInfo
}

func createNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demonote.NoteServiceCreateNoteArgs)
	realResult := result.(*demonote.NoteServiceCreateNoteResult)
	success, err := handler.(demonote.NoteService).CreateNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceCreateNoteArgs() interface{} {
	return demonote.NewNoteServiceCreateNoteArgs()
}

func newNoteServiceCreateNoteResult() interface{} {
	return demonote.NewNoteServiceCreateNoteResult()
}

func deleteNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demonote.NoteServiceDeleteNoteArgs)
	realResult := result.(*demonote.NoteServiceDeleteNoteResult)
	success, err := handler.(demonote.NoteService).DeleteNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceDeleteNoteArgs() interface{} {
	return demonote.NewNoteServiceDeleteNoteArgs()
}

func newNoteServiceDeleteNoteResult() interface{} {
	return demonote.NewNoteServiceDeleteNoteResult()
}

func updateNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demonote.NoteServiceUpdateNoteArgs)
	realResult := result.(*demonote.NoteServiceUpdateNoteResult)
	success, err := handler.(demonote.NoteService).UpdateNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceUpdateNoteArgs() interface{} {
	return demonote.NewNoteServiceUpdateNoteArgs()
}

func newNoteServiceUpdateNoteResult() interface{} {
	return demonote.NewNoteServiceUpdateNoteResult()
}

func queryNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demonote.NoteServiceQueryNoteArgs)
	realResult := result.(*demonote.NoteServiceQueryNoteResult)
	success, err := handler.(demonote.NoteService).QueryNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceQueryNoteArgs() interface{} {
	return demonote.NewNoteServiceQueryNoteArgs()
}

func newNoteServiceQueryNoteResult() interface{} {
	return demonote.NewNoteServiceQueryNoteResult()
}

func mGetNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demonote.NoteServiceMGetNoteArgs)
	realResult := result.(*demonote.NoteServiceMGetNoteResult)
	success, err := handler.(demonote.NoteService).MGetNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceMGetNoteArgs() interface{} {
	return demonote.NewNoteServiceMGetNoteArgs()
}

func newNoteServiceMGetNoteResult() interface{} {
	return demonote.NewNoteServiceMGetNoteResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateNote(ctx context.Context, req *demonote.CreateNoteRequest) (r *demonote.CreateNoteResponse, err error) {
	var _args demonote.NoteServiceCreateNoteArgs
	_args.Req = req
	var _result demonote.NoteServiceCreateNoteResult
	if err = p.c.Call(ctx, "CreateNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteNote(ctx context.Context, req *demonote.DeleteNoteRequest) (r *demonote.DeleteNoteResponse, err error) {
	var _args demonote.NoteServiceDeleteNoteArgs
	_args.Req = req
	var _result demonote.NoteServiceDeleteNoteResult
	if err = p.c.Call(ctx, "DeleteNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateNote(ctx context.Context, req *demonote.UpdateNoteRequest) (r *demonote.UpdateNoteResponse, err error) {
	var _args demonote.NoteServiceUpdateNoteArgs
	_args.Req = req
	var _result demonote.NoteServiceUpdateNoteResult
	if err = p.c.Call(ctx, "UpdateNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryNote(ctx context.Context, req *demonote.QueryNoteRequest) (r *demonote.QueryNoteResponse, err error) {
	var _args demonote.NoteServiceQueryNoteArgs
	_args.Req = req
	var _result demonote.NoteServiceQueryNoteResult
	if err = p.c.Call(ctx, "QueryNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetNote(ctx context.Context, req *demonote.MGetNoteRequest) (r *demonote.MGetNoteResponse, err error) {
	var _args demonote.NoteServiceMGetNoteArgs
	_args.Req = req
	var _result demonote.NoteServiceMGetNoteResult
	if err = p.c.Call(ctx, "MGetNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
