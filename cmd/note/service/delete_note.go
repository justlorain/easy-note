package service

import (
	"context"
	"easy-note/cmd/note/dal/db"
	"easy-note/kitex_gen/notedemo"
)

type DelNoteService struct {
	ctx context.Context
}

// NewDelNoteService new DelNoteService
func NewDelNoteService(ctx context.Context) *DelNoteService {
	return &DelNoteService{
		ctx: ctx,
	}
}

// DelNote delete note info
func (s *DelNoteService) DelNote(req *notedemo.DeleteNoteRequest) error {
	return db.DeleteNote(s.ctx, req.NoteId, req.UserId)
}
