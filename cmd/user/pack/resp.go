package pack

import (
	"easy-note/kitex_gen/demouser"
	"easy-note/pkg/errno"
	"errors"
	"time"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *demouser.BaseResp {
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

func baseResp(err errno.ErrNo) *demouser.BaseResp {
	return &demouser.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
