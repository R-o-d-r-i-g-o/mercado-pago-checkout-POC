package adapter

import (
	"code-space-backend-api/common/logger/field"

	"go.uber.org/zap/zapcore"
)

type ZapFieldFunc[T any] func(key string, value T) zapcore.Field

func ZapToFieldFunc[T any](fn ZapFieldFunc[T]) field.FieldFunc[T] {
	return func(key string, value T) field.Field {
		zapField := fn(key, value)
		return field.Field{
			ZapField: zapField,
		}
	}
}

func FieldsToZap(fields []field.Field) []zapcore.Field {
	zapFields := make([]zapcore.Field, 0, len(fields))
	for _, field := range fields {
		zapFields = append(zapFields, field.ZapField)
	}

	return zapFields
}
