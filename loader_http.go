package configuration

import (
	"context"
	"net"

	"github.com/containerssh/log"
	"github.com/containerssh/structutils"
)

// NewHTTPLoader loads configuration from HTTP servers for specific connections.
//goland:noinspection GoUnusedExportedFunction
func NewHTTPLoader(
	config ClientConfig,
	logger log.Logger,
) (Loader, error) {
	client, err := NewClient(config, logger)
	if err != nil {
		return nil, err
	}
	return &httpLoader{
		client: client,
	}, nil
}

type httpLoader struct {
	client Client
}

func (h *httpLoader) Load(_ context.Context, _ *AppConfig) error {
	return nil
}

func (h *httpLoader) LoadConnection(
	ctx context.Context,
	username string,
	remoteAddr net.TCPAddr,
	connectionID string,
	config *AppConfig,
) error {
	newAppConfig, err := h.client.Get(ctx, username, remoteAddr, connectionID)
	if err != nil {
		return err
	}
	return structutils.Merge(config, &newAppConfig)
}
