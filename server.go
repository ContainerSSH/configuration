package configuration

import (
	"github.com/containerssh/http"
	"github.com/containerssh/log"
)

// NewServer returns a complete HTTP server that responds to the configuration requests.
//goland:noinspection GoUnusedExportedFunction
func NewServer(
	configuration http.ServerConfiguration,
	h ConfigRequestHandler,
	logger log.Logger,
) (http.Server, error) {
	handler, err := NewHandler(h, logger)
	if err != nil {
		return nil, err
	}
	return http.NewServer(
		"Config Server",
		configuration,
		handler,
		logger,
	)
}
