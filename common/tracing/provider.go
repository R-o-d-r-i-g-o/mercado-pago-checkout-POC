package tracing

import "context"

type Provider interface {
	Shutdown(context.Context) error
}
