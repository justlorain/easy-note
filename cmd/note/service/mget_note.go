package service

import (
	"context"
	"easy-note/cmd/note/dal/db"
	"easy-note/cmd/note/pack"
	"easy-note/cmd/note/rpc"
	"easy-note/kitex_gen/demonote"
	"easy-note/kitex_gen/demouser"
)

type MGetNoteService struct {
	ctx context.Context
}

// NewMGetNoteService new MGetNoteService
func NewMGetNoteService(ctx context.Context) *MGetNoteService {
	return &MGetNoteService{ctx: ctx}
}

// MGetNote multiple get list of note info
func (s *MGetNoteService) MGetNote(req *demonote.MGetNoteRequest) ([]*demonote.Note, error) {
	noteModels, err := db.MGetNotes(s.ctx, req.NoteIds)
	if err != nil {
		return nil, err
	}
	uIds := pack.UserIds(noteModels)
	userMap, err := rpc.MGetUser(s.ctx, &demouser.MGetUserRequest{UserIds: uIds})
	if err != nil {
		return nil, err
	}
	notes := pack.Notes(noteModels)
	for i := 0; i < len(notes); i++ {
		if u, ok := userMap[notes[i].UserId]; ok {
			notes[i].Username = u.Username
			notes[i].UserAvatar = u.Avatar
		}
	}
	return notes, nil
}
