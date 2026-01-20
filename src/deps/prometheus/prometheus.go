package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Prometheus struct {
	*prometheus.Registry
	Middleware
}

// Создаёт кастомный реестр метрик и middleware для использовании в приложении
func New() *Prometheus {
	registry := prometheus.NewRegistry()

	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	middleware := newMiddleware(registry, []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1.0, 1.5, 3, 5})

	return &Prometheus{
		Registry:   registry,
		Middleware: middleware,
	}
}
