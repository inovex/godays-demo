package pkg

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"io"
	"log"
)

func InitGlobalTracer() (io.Closer, error) {
	log.Printf("Initializing global tracer")

	cfg, err := (&config.Configuration{}).FromEnv()
	if err != nil {
		return nil, err
	}
	cfg.Headers = (&jaeger.HeadersConfig{}).ApplyDefaults()
	tracer, closer, err := cfg.NewTracer(
		config.Logger(jaegerlog.StdLogger),
		config.Metrics(metrics.NullFactory),
	)
	opentracing.SetGlobalTracer(tracer)

	return closer, err
}
