package context

import (
	"context"
	"todo/pkg/constant"
	"todo/pkg/logger"
	"todo/pkg/util"
)

func NewContext(ctx context.Context) context.Context {
	if requestID, ok := ctx.Value(constant.ContextRequestId).(string); !ok || requestID == "" {
		requestID = util.NewID()
		ctx = context.WithValue(ctx, constant.ContextRequestId, requestID)
	}

	var log = GetLoggerFromContext(ctx).ApplyContext(ctx)
	ctx = AppendLoggerToContext(ctx, log)
	return ctx
}

func AppendLoggerToContext(ctx context.Context, log *logger.Logger) context.Context {
	return context.WithValue(ctx, constant.ContextLogger, log)
}

func GetLoggerFromContext(ctx context.Context) *logger.Logger {
	if log, ok := ctx.Value(constant.ContextLogger).(*logger.Logger); ok && log != nil {
		return log
	} else {
		return logger.Default()
	}
}
