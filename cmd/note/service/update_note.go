package service

import (
	"context"
	"easy-note/cmd/note/dal/db"
	"easy-note/kitex_gen/notedemo"
)

type UpdateNoteService struct {
	ctx context.Context
}

// NewUpdateNoteService new UpdateNoteService
func NewUpdateNoteService(ctx context.Context) *UpdateNoteService {
	return &UpdateNoteService{ctx: ctx}
}

// UpdateNote update note info
func (s *UpdateNoteService) UpdateNote(req *notedemo.UpdateNoteRequest) error {
	return db.UpdateNote(s.ctx, req.NoteId, req.UserId, req.Title, req.Content)
}
