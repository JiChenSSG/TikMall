package obs

import (
	"context"

	"github.com/cloudwego/kitex/server"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var (
	traceProvider *tracesdk.TracerProvider
)

func GetTraceProvider() *tracesdk.TracerProvider {
	return traceProvider
}

func InitTracer(serviceName string) {
	// p := provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName(serviceName),
	// 	provider.WithExportEndpoint(endpoint),
	// 	provider.WithInsecure(),
	// 	provider.WithEnableMetrics(false),
	// )

	// server.RegisterShutdownHook(func() {
	// 	if err := p.Shutdown(context.Background()); err != nil {
	// 		log.Printf("Failed to shutdown OpenTelemetry provider: %v", err)
	// 	} else {
	// 		log.Println("OpenTelemetry provider shutdown")
	// 	}
	// })

	exporter, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithInsecure(),
	)

	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		exporter.Shutdown(context.Background()) //nolint:errcheck
	})

	processor := tracesdk.NewBatchSpanProcessor(exporter)
	res, err := resource.New(
		context.Background(), 
		resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)),
	)
	if err != nil {
		res = resource.Default()
	}

	traceProvider = tracesdk.NewTracerProvider(
		tracesdk.WithSpanProcessor(processor), 
		tracesdk.WithResource(res),
	)
	otel.SetTracerProvider(traceProvider)
}
