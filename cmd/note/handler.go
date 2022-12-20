package main

import (
	"context"
	"easy-note/cmd/note/pack"
	"easy-note/cmd/note/service"
	"easy-note/kitex_gen/demonote"
	"easy-note/pkg/consts"
	"easy-note/pkg/errno"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}

// CreateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) CreateNote(ctx context.Context, req *demonote.CreateNoteRequest) (resp *demonote.CreateNoteResponse, err error) {
	resp = new(demonote.CreateNoteResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateNoteService(ctx).CreateNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DeleteNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) DeleteNote(ctx context.Context, req *demonote.DeleteNoteRequest) (resp *demonote.DeleteNoteResponse, err error) {
	resp = new(demonote.DeleteNoteResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewDelNoteService(ctx).DelNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// UpdateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) UpdateNote(ctx context.Context, req *demonote.UpdateNoteRequest) (resp *demonote.UpdateNoteResponse, err error) {
	resp = new(demonote.UpdateNoteResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUpdateNoteService(ctx).UpdateNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// QueryNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) QueryNote(ctx context.Context, req *demonote.QueryNoteRequest) (resp *demonote.QueryNoteResponse, err error) {
	resp = new(demonote.QueryNoteResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	if req.Limit == 0 {
		req.Limit = consts.DefaultLimit
	}

	notes, total, err := service.NewQueryNoteService(ctx).QueryNoteService(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Notes = notes
	resp.Total = total
	return resp, nil
}

// MGetNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) MGetNote(ctx context.Context, req *demonote.MGetNoteRequest) (resp *demonote.MGetNoteResponse, err error) {
	resp = new(demonote.MGetNoteResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	notes, err := service.NewMGetNoteService(ctx).MGetNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Notes = notes
	return resp, nil
}
