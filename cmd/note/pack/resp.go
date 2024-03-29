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

package pack

import (
	"errors"
	"time"

	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/demonote"
	"github.com/cloudwego/biz-demo/easy_note/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *demonote.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *demonote.BaseResp {
	return &demonote.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}
