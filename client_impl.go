package configuration

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/containerssh/http"
	"github.com/containerssh/log"
	"github.com/containerssh/metrics"
)

type client struct {
	httpClient           http.Client
	logger               log.Logger
	backendFailureMetric metrics.SimpleCounter
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
	request, response := c.createRequestResponse(username, remoteAddr, connectionID)
	var lastError error = nil
	var lastLabels []metrics.MetricLabel
loop:
	for {
		lastLabels = []metrics.MetricLabel{}
		statusCode, err := c.httpClient.Post("", &request, &response)
		lastError = err
		if err != nil {
			clientError := &http.ClientError{}
			if errors.As(err, clientError) {
				lastLabels = []metrics.MetricLabel{
					metrics.Label("type", "soft"),
					metrics.Label("reason", string(clientError.Reason)),
				}
			}
			c.logger.Warningf("HTTP query to config server failed, retrying in 10 seconds (%v)", err)
		} else if statusCode != 200 {
			lastLabels = []metrics.MetricLabel{
				metrics.Label("type", "soft"),
				metrics.Label("reason", "invalid_status_code"),
				metrics.Label("status_code", strconv.Itoa(statusCode)),
			}
			lastError = fmt.Errorf("invalid response status %d from config server", statusCode)
			c.logger.Warningf("invalid response status %d from config server, retrying in 10 seconds", statusCode)
		}
		if lastError == nil {
			break loop
		} else {
			c.backendFailureMetric.Increment(
				lastLabels...,
			)
		}
		select {
		case <-ctx.Done():
			break loop
		case <-time.After(10 * time.Second):
		}
	}
	if lastError != nil {
		c.logger.Errorf("failed to query config server, giving up. (%v)", lastError)
		c.backendFailureMetric.Increment(
			append(
				[]metrics.MetricLabel{
					metrics.Label("type", "hard"),
				}, lastLabels...,
			)...,
		)
	}
	return response.Config, lastError
}

func (c *client) createRequestResponse(username string, remoteAddr net.TCPAddr, connectionID string) (
	ConfigRequest,
	ConfigResponseBody,
) {
	request := ConfigRequest{
		Username:     username,
		RemoteAddr:   remoteAddr.IP.String(),
		ConnectionID: connectionID,
		SessionID:    connectionID,
	}
	response := ConfigResponseBody{}
	return request, response
}
