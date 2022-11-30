package otel

import (
	"context"
	"easy-note/pkg/consts"
	"gorm.io/plugin/opentelemetry/provider"
)

func Init(serviceName string) {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
}
