package logger

import (
	"context"
	"todo/pkg/constant"
)

func getFromContext(ctx context.Context, key string) (value string, ok bool) {
	value, ok = ctx.Value(key).(string)
	return
}

func (g Logger) ApplyContext(ctx context.Context) *Logger {
	log := &g
	if value, ok := getFromContext(ctx, constant.ContextRequestId); ok {
		log = log.WithRequestID(value)
	}
	return log
}
