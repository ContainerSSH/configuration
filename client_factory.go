package configuration

import (
	"github.com/containerssh/http"
	"github.com/containerssh/log"
	"github.com/containerssh/metrics"
)

// MetricNameConfigBackendFailure is the number of request failures to the configuration backend.
const MetricNameConfigBackendFailure = "containerssh_config_server_failures"

// NewClient creates a new configuration client that can be used to fetch a user-specific configuration.
func NewClient(
	config ClientConfig,
	logger log.Logger,
	metricsCollector metrics.Collector,
) (Client, error) {
	var httpClient http.Client
	var err error
	if config.ClientConfiguration.URL != "" {
		httpClient, err = http.NewClient(config.ClientConfiguration, logger)
		if err != nil {
			return nil, err
		}
	}
	backendFailureMetric, err := metricsCollector.CreateCounter(
		MetricNameConfigBackendFailure,
		"",
		"The number of request failures to the configuration backend.",
	)
	if err != nil {
		return nil, err
	}
	return &client{
		httpClient:           httpClient,
		logger:               logger,
		backendFailureMetric: backendFailureMetric,
	}, nil
}
