package obs


import (
	"context"

	"github.com/cloudwego/hertz/pkg/route"
	"github.com/jichenssg/tikmall/gateway/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var TracerProvider *tracesdk.TracerProvider

func initTracer() route.CtxCallback {
	exporter, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}

	processor := tracesdk.NewBatchSpanProcessor(exporter)
	res, err := resource.New(context.Background(), resource.WithAttributes(semconv.ServiceNameKey.String(config.GetConf().Server.Name)))
	if err != nil {
		res = resource.Default()
	}
	
	TracerProvider = tracesdk.NewTracerProvider(
		tracesdk.WithSpanProcessor(processor),
		tracesdk.WithResource(res),
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
	)
	otel.SetTracerProvider(TracerProvider)

	return route.CtxCallback(func(ctx context.Context) {
		exporter.Shutdown(ctx) //nolint:errcheck
	})
}