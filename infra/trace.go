package infra

import (
	"os"

	"github.com/pkg/errors"
	"go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
)

// InitJaegerTracing init a jaeger exporter and register it.
// The agent endpoint comes from the env JAEGER_SERVICE_ADDR.
func InitJaegerTracing(name string) error {
	svcAddr := os.Getenv(JaegerAddrEnv)
	if svcAddr == "" {
		return ErrNoEnv
	}

	// Register the Jaeger exporter to be able to retrieve
	// the collected spans.
	exporter, err := jaeger.NewExporter(jaeger.Options{
		// Endpoint is deprecated
		AgentEndpoint: svcAddr,
		// CollectorEndpoint: "http://localhost:14268/api/traces", // optional
		Process: jaeger.Process{
			ServiceName: name,
		},
	})
	if err != nil {
		return errors.Wrap(err, "init jaeger")
	}

	trace.RegisterExporter(exporter)
	return nil
}
