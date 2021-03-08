package configuration

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/containerssh/http"
	"github.com/containerssh/log"
	"github.com/containerssh/metrics"
)

type client struct {
	httpClient            http.Client
	logger                log.Logger
	backendRequestsMetric metrics.SimpleCounter
	backendFailureMetric  metrics.SimpleCounter
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
	logger := c.logger.
		WithLabel("connectionId", connectionID).
		WithLabel("username", username)
	request, response := c.createRequestResponse(username, remoteAddr, connectionID)
	var lastError error = nil
	var lastLabels []metrics.MetricLabel
loop:
	for {
		lastLabels = []metrics.MetricLabel{}
		if lastError != nil {
			lastLabels = append(
				lastLabels,
				metrics.Label("retry", "1"),
			)
		} else {
			lastLabels = append(
				lastLabels,
				metrics.Label("retry", "0"),
			)
		}
		c.logAttempt(logger, lastLabels)

		lastError = c.configServerRequest(&request, &response)
		if lastError == nil {
			c.logConfigResponse(logger)
			return response.Config, nil
		}
		reason := c.getReason(lastError)
		lastLabels = append(lastLabels, metrics.Label("reason", reason))
		c.logTemporaryFailure(logger, lastError, reason, lastLabels)
		select {
		case <-ctx.Done():
			break loop
		case <-time.After(10 * time.Second):
		}
	}
	return c.logAndReturnPermanentFailure(lastError, lastLabels, logger)
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

func (c *client) logAttempt(logger log.Logger, lastLabels []metrics.MetricLabel) {
	logger.Debug(
		log.NewMessage(
			MConfig,
			"Configuration request",
		),
	)
	c.backendRequestsMetric.Increment(lastLabels...)
}

func (c *client) logAndReturnPermanentFailure(
	lastError error,
	lastLabels []metrics.MetricLabel,
	logger log.Logger,
) (AppConfig, error) {
	err := log.Wrap(
		lastError,
		EConfigBackendError,
		"Configuration request to backend failed, giving up",
	)
	c.backendFailureMetric.Increment(
		append(
			[]metrics.MetricLabel{
				metrics.Label("type", "hard"),
			}, lastLabels...,
		)...,
	)
	logger.Error(err)
	return AppConfig{}, err
}

func (c *client) logTemporaryFailure(
	logger log.Logger,
	lastError error,
	reason string,
	lastLabels []metrics.MetricLabel,
) {
	logger.Debug(
		log.Wrap(
			lastError,
			EConfigBackendError,
			"Configuration request to backend failed, retrying in 10 seconds",
		).
			Label("reason", reason),
	)
	c.backendFailureMetric.Increment(
		append(
			[]metrics.MetricLabel{
				metrics.Label("type", "soft"),
			}, lastLabels...,
		)...,
	)
}

func (c *client) getReason(lastError error) string {
	var typedErr log.Message
	reason := log.EUnknownError
	if errors.As(lastError, &typedErr) {
		reason = typedErr.Code()
	}
	return reason
}

func (c *client) logConfigResponse(
	logger log.Logger,
) {
	logger.Debug(
		log.NewMessage(
			MConfigSuccess,
			"User-specific configuration received",
		),
	)
}

func (c *client) configServerRequest(requestObject interface{}, response interface{}) error {
	statusCode, err := c.httpClient.Post("", requestObject, response)
	if err != nil {
		return err
	}
	if statusCode != 200 {
		return log.UserMessage(
			EInvalidStatus,
			// The message indicates authentication because the config server is
			// called at config-time.
			"Cannot authenticate at this time.",
			"Configuration server responded with an invalid status code: %d",
			statusCode,
		)
	}
	return nil
}
