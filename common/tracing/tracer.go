package tracing

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func setTracer(t trace.Tracer) {
	tracer = t
}

func Tracer() trace.Tracer {
	return tracer
}

// Start inits a new trace span
func Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return tracer.Start(ctx, spanName, opts...)
}

// TraceID gets trace id from context
func TraceID(ctx context.Context) string {
	return trace.SpanContextFromContext(ctx).TraceID().String()
}

// SpanID gets span id from context
func SpanID(ctx context.Context) string {
	return trace.SpanContextFromContext(ctx).SpanID().String()
}
