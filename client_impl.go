package configuration

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/containerssh/http"
	"github.com/containerssh/log"
)

type client struct {
	httpClient http.Client
	logger     log.Logger
}

func (c *client) Get(
	ctx context.Context,
	username string,
	remoteAddr net.TCPAddr,
	connectionID string,
) (AppConfig, error) {
	if c.httpClient == nil {
		return AppConfig{}, nil
	}
	request := ConfigRequest{
		Username:     username,
		RemoteAddr:   remoteAddr.IP.String(),
		ConnectionID: connectionID,
		SessionID:    connectionID,
	}
	response := ConfigResponseBody{}
	var lastError error = nil
loop:
	for {
		statusCode, err := c.httpClient.Post("", &request, &response)
		lastError = err
		if err != nil {
			c.logger.Warningf("HTTP query to config server failed, retrying in 10 seconds (%v)", err)
		} else if statusCode != 200 {
			lastError = fmt.Errorf("invalid response status %d from config server", statusCode)
			c.logger.Warningf("invalid response status %d from config server, retrying in 10 seconds", statusCode)
		}
		if lastError == nil {
			break loop
		}
		select {
		case <-ctx.Done():
			break loop
		case <-time.After(10 * time.Second):
		}
	}
	if lastError != nil {
		c.logger.Errorf("failed to query config server, giving up. (%v)", lastError)
	}
	return response.Config, lastError
}
