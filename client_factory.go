package configuration

import (
	"github.com/containerssh/http"
	"github.com/containerssh/log"
)

// NewClient creates a new configuration client that can be used to fetch a user-specific configuration.
func NewClient(
	config ClientConfig,
	logger log.Logger,
) (Client, error) {
	httpClient, err := http.NewClient(config.ClientConfiguration, logger)
	if err != nil {
		return nil, err
	}
	return &client{
		httpClient: httpClient,
		logger:     logger,
	}, nil
}
