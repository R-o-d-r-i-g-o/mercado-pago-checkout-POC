package config

import (
	"code-space-backend-api/common/metrics"
)

type Metrics struct {
	Gin    *metrics.Gin
	Server *metrics.StandaloneServer
}

func provideMetrics(c *Container) {
	c.Metrics = new(Metrics)
	c.Metrics.Gin = metrics.NewMetricsGin()
}
