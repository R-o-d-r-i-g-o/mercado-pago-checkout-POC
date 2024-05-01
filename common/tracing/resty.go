package tracing

import (
	"github.com/go-resty/resty/v2"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// RestyOnBeforeRequest sets the before request middleware to inject trace context into request's headers for trace propagation
func RestyOnBeforeRequest(c *resty.Client) *resty.Client {
	return c.OnBeforeRequest(RestyInjectOtelHeaders)
}

// RestyInjectOtelHeaders inject trace context into request's headers for trace propagation
func RestyInjectOtelHeaders(c *resty.Client, r *resty.Request) error {
	otel.GetTextMapPropagator().Inject(r.Context(), propagation.HeaderCarrier(r.Header))
	return nil
}
