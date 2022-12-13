package main

import (
	"context"
	notedemo "easy-note/temp/demonote/kitex_gen/notedemo"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}

// CreateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) CreateNote(ctx context.Context, req *notedemo.CreateNoteRequest) (resp *notedemo.CreateNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) DeleteNote(ctx context.Context, req *notedemo.DeleteNoteRequest) (resp *notedemo.DeleteNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) UpdateNote(ctx context.Context, req *notedemo.UpdateNoteRequest) (resp *notedemo.UpdateNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) QueryNote(ctx context.Context, req *notedemo.QueryNoteRequest) (resp *notedemo.QueryNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) MGetNote(ctx context.Context, req *notedemo.MGetNoteRequest) (resp *notedemo.MGetNoteResponse, err error) {
	// TODO: Your code here...
	return
}
