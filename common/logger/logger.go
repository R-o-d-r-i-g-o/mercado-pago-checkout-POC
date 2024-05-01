package logger

import (
	"context"

	"code-space-backend-api/common/logger/adapter"
	"code-space-backend-api/common/logger/field"
	"code-space-backend-api/common/trace"
	"code-space-backend-api/common/tracing"

	"go.uber.org/zap"
)

type LogEngine interface {
	WithContext(ctx context.Context) LogEngine
	Error(err error, fields ...field.Field)
	Info(message string, fields ...field.Field)
	Debug(message string, fields ...field.Field)
	SyncLogs()
}

type logEngine struct {
	ctx    context.Context
	zapLog *zap.Logger
}

func New(mode string) (LogEngine, error) {
	zapLog, err := getCurrentModeLog(mode)
	if err != nil {
		return nil, ContextError.WithMessage("cant initialize logger: %w", err)
	}
	return &logEngine{zapLog: zapLog, ctx: context.Background()}, nil
}

func getCurrentModeLog(debugMode string) (*zap.Logger, error) {
	if debugMode == "production" {
		return zap.NewProduction(zap.AddCallerSkip(1))
	} else if debugMode == "development" {
		return zap.NewDevelopment(zap.AddCallerSkip(1))
	}

	return nil, ContextError.WithMessage("invalid mode")
}

func (log *logEngine) WithContext(ctx context.Context) LogEngine {
	return &logEngine{
		zapLog: log.zapLog,
		ctx:    ctx,
	}
}

func (log *logEngine) Info(message string, fields ...field.Field) {
	zapFields := append(
		adapter.FieldsToZap(fields),
		zap.String(trace.Key, trace.FromContext(log.ctx)), // TODO: Deprecate and use only tracing package
		zap.String(tracing.Key, tracing.TraceID(log.ctx)),
	)

	log.zapLog.Info(message, zapFields...)
}

func (log *logEngine) Debug(message string, fields ...field.Field) {
	zapFields := append(
		adapter.FieldsToZap(fields),
		zap.String(trace.Key, trace.FromContext(log.ctx)), // TODO: Deprecate and use only tracing package
		zap.String(tracing.Key, tracing.TraceID(log.ctx)),
	)
	log.zapLog.Debug(message, zapFields...)
}

func (log *logEngine) Error(err error, fields ...field.Field) {
	zapFields := append(
		adapter.FieldsToZap(fields),
		zap.String(trace.Key, trace.FromContext(log.ctx)), // TODO: Deprecate and use only tracing package
		zap.Object("error", adapter.ZapError(err)),
		zap.String(tracing.Key, tracing.TraceID(log.ctx)),
	)

	log.zapLog.Error(err.Error(), zapFields...)
}

func (log *logEngine) SyncLogs() {
	log.zapLog.Sync()
}
