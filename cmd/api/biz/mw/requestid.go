package mw

import (
	"github.com/sirupsen/logrus"
)

type RequestIDHook struct{}

func (h *RequestIDHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *RequestIDHook) Fire(e *logrus.Entry) error {
	ctx := e.Context
	if ctx == nil {
		return nil
	}
	requestid := ctx.Value("X-Request-ID")
	if requestid != nil {
		e.Data["request_id"] = requestid
	}
	return nil
}
