package logger

import (
	"context"

	"code-space-backend-api/common/logger/field"
)

type noopLogger struct {
}

func NewNoopLogger() LogEngine {
	return noopLogger{}
}

func (l noopLogger) WithContext(ctx context.Context) LogEngine {
	return l
}
func (noopLogger) Error(err error, fields ...field.Field)      {}
func (noopLogger) Info(message string, fields ...field.Field)  {}
func (noopLogger) Debug(message string, fields ...field.Field) {}
func (noopLogger) SyncLogs()                                   {}
