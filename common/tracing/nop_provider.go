package tracing

import "context"

// NopProvider is user in case to bypass Provider interface where the implementation has no need of use
type NopProvider struct{}

func NewNopProvider() Provider {
	return NopProvider{}
}

func (NopProvider) Shutdown(ctx context.Context) error {
	return nil
}
