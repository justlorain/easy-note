// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service

import (
	"context"
	"github.com/cloudwego/biz-demo/easy_note/cmd/note/dal/db"
	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/demonote"
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
func (s *DelNoteService) DelNote(req *demonote.DeleteNoteRequest) error {
	return db.DeleteNote(s.ctx, req.NoteId, req.UserId)
}
