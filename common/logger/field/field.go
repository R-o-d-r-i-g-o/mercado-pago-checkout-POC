package field

import "go.uber.org/zap/zapcore"

type Field struct {
	ZapField zapcore.Field
}

type FieldFunc[T any] func(key string, value T) Field
