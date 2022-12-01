package otel

import (
	"easy-note/pkg/consts"
	"gorm.io/plugin/opentelemetry/provider"
)

func Init(serviceName string) {
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
}
