package logger

import (
	"context"
	"go.uber.org/zap"
)

const RequestID = "request_id"

var zapLog *zap.Logger

func init() {
	var logErr error
	config := zap.NewProductionConfig()
	config.EncoderConfig = zap.NewProductionEncoderConfig()

	zapLog, logErr = config.Build(zap.AddCallerSkip(1))
	if logErr != nil {
		panic(logErr)
	}
}

func Debug(ctx context.Context, message string, fields ...zap.Field) {
	zapLog.Debug(message, append(addRequestId(ctx), fields...)...)
}

func Info(ctx context.Context, message string, fields ...zap.Field) {
	zapLog.Info(message, append(addRequestId(ctx), fields...)...)
}

func Error(ctx context.Context, message string, fields ...zap.Field) {
	zapLog.Error(message, append(addRequestId(ctx), fields...)...)
}

func Fatal(ctx context.Context, message string, fields ...zap.Field) {
	zapLog.Fatal(message, append(addRequestId(ctx), fields...)...)
}

func addRequestId(ctx context.Context) []zap.Field {
	return []zap.Field{zap.String(RequestID, ctx.Value(RequestID).(string))}
}
