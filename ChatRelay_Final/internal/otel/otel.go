// internal/otel/otel.go
package otel

import (
	"chatrelay/internal/config"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

type ShutdownFunc func(ctx context.Context) error

func InitTracerProvider(cfg *config.Config) (ShutdownFunc, error) {
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("ChatRelayBot"),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown, nil
}
