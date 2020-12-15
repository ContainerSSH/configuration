package configuration_test

import (
	"context"
	"net"
	"os"
	"testing"

	"github.com/containerssh/geoip"
	"github.com/containerssh/http"
	"github.com/containerssh/log"
	"github.com/containerssh/metrics"
	"github.com/containerssh/service"
	"github.com/stretchr/testify/assert"

	"github.com/containerssh/configuration"
)

func TestHTTP(t *testing.T) {
	logger, err := log.New(log.Config{
		Level:  log.LevelDebug,
		Format: "text",
	}, "config", os.Stdout)
	assert.NoError(t, err)
	srv, err := configuration.NewServer(
		http.ServerConfiguration{
			Listen: "127.0.0.1:8080",
		},
		&myConfigReqHandler{},
		logger,
	)
	assert.NoError(t, err)
	lifecycle := service.NewLifecycle(srv)

	ready := make(chan struct{})
	lifecycle.OnRunning(
		func(s service.Service, l service.Lifecycle) {
			ready <- struct{}{}
		})
	go func() {
		_ = lifecycle.Run()
	}()
	<-ready

	client, err := configuration.NewClient(
		configuration.ClientConfig{
			ClientConfiguration: http.ClientConfiguration{
				URL: "http://127.0.0.1:8080",
			},
		}, logger, getMetricsCollector(t),
	)
	assert.NoError(t, err)

	connectionID := "0123456789ABCDEF"

	config, err := client.Get(
		context.Background(),
		"foo",
		net.TCPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 2222,
		},
		connectionID,
	)
	assert.NoError(t, err)
	assert.Equal(t, "yourcompany/yourimage", config.DockerRun.Config.ContainerConfig.Image)

	lifecycle.Stop(context.Background())
	err = lifecycle.Wait()
	assert.NoError(t, err)
}

func getMetricsCollector(t *testing.T) metrics.Collector {
	geoIP, err := geoip.New(geoip.Config{
		Provider: "dummy",
	})
	assert.NoError(t, err)
	return metrics.New(geoIP)
}

type myConfigReqHandler struct {
}

func (m *myConfigReqHandler) OnConfig(
	request configuration.ConfigRequest,
) (config configuration.AppConfig, err error) {
	// We recommend using an IDE to discover the possible options here.
	if request.Username == "foo" {
		config.DockerRun.Config.ContainerConfig.Image = "yourcompany/yourimage"
	}
	return config, err
}
