package obs

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTrack(serviceName string, endpoint string) {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint(endpoint),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
	)

	server.RegisterShutdownHook(func() {
		if err := p.Shutdown(context.Background()); err != nil {
			log.Printf("Failed to shutdown OpenTelemetry provider: %v", err)
		} else {
			log.Println("OpenTelemetry provider shutdown")
		}
	})

}
