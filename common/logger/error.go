package logger

import "code-space-backend-api/common/errors"

var ContextError = errors.New().WithType(errors.TypeConfig)
