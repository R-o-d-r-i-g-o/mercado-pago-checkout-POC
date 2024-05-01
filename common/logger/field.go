package logger

import (
	"code-space-backend-api/common/logger/adapter"

	"go.uber.org/zap"
)

var String = adapter.ZapToFieldFunc(zap.String)
var BytesString = adapter.ZapToFieldFunc(zap.ByteString)
var Any = adapter.ZapToFieldFunc(zap.Any)
