package mw

import (
	"context"
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
	traceID := e.Data["trace_id"]
	context.WithValue(ctx, "X-Request-ID", traceID)
	e.Data["request_id"] = traceID
	return nil
}
