package main

import (
	"context"
	demonote "easy-note/temp/note/kitex_gen/demonote"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}

// CreateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) CreateNote(ctx context.Context, req *demonote.CreateNoteRequest) (resp *demonote.CreateNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) DeleteNote(ctx context.Context, req *demonote.DeleteNoteRequest) (resp *demonote.DeleteNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) UpdateNote(ctx context.Context, req *demonote.UpdateNoteRequest) (resp *demonote.UpdateNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) QueryNote(ctx context.Context, req *demonote.QueryNoteRequest) (resp *demonote.QueryNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) MGetNote(ctx context.Context, req *demonote.MGetNoteRequest) (resp *demonote.MGetNoteResponse, err error) {
	// TODO: Your code here...
	return
}
