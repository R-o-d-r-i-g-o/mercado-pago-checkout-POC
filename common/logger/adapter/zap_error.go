package adapter

import (
	"code-space-backend-api/common/errors"

	"go.uber.org/zap/zapcore"
)

type zapErrorMarshaller struct {
	err error
}

func ZapError(err error) zapcore.ObjectMarshaler {
	return &zapErrorMarshaller{err}
}

func (z *zapErrorMarshaller) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if err, ok := z.err.(errors.Error); ok {
		addStringOmitEmpty(enc, "type", string(err.Type))
		addStringOmitEmpty(enc, "context", err.Context)
		addStringOmitEmpty(enc, "name", err.Name)
		addStringOmitEmpty(enc, "message", err.Message)
		addStringOmitEmpty(enc, "param", err.Param)
		if len(err.Errors) > 0 {
			enc.AddArray("errors", z)
		}

		return nil
	}

	enc.AddString("message", z.err.Error())

	return nil
}

func (z *zapErrorMarshaller) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	if err, ok := z.err.(errors.Error); ok {
		for _, err := range err.Errors {
			enc.AppendObject(&zapErrorMarshaller{err})
		}

		return nil
	}
	return nil
}

func addStringOmitEmpty(enc zapcore.ObjectEncoder, key string, val string) {
	if val == "" {
		return
	}

	enc.AddString(key, val)
}
