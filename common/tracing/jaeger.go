package tracing

import (
	"fmt"
	"log"

	"code-space-backend-api/common/uuid"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func newJaegerProvider() (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithAgentEndpoint())
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(ServiceName()),
			attribute.String("environment", Environment()),
			attribute.String("ID", ServiceID()),
		)),
	)
	return tp, nil
}

func SetupJaegerProvider(name, stage string) Provider {
	serviceName = name
	environment = stage
	serviceID = uuid.New()

	tp, err := newJaegerProvider()
	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	setTracer(tp.Tracer(fmt.Sprint(ServiceName(), "-tracer")))
	return tp
}
