package trace

import "context"

// Deprecated: due new tracing package
const Key = "trace"

// Deprecated: due new tracing package
func FromContext(ctx context.Context) string {
	if v, ok := ctx.Value(Key).(string); ok {
		return v
	}
	return ""
}

// Deprecated: due new tracing package
func NewContext(ctx context.Context, trace string) context.Context {
	return context.WithValue(ctx, Key, trace)
}
