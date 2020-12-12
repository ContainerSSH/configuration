package configuration_test

import (
	"context"
	"net"
	"os"
	"testing"

	"github.com/containerssh/http"
	"github.com/containerssh/log"
	"github.com/containerssh/service"
	"gotest.tools/assert"

	"github.com/containerssh/configuration"
)

func TestHTTP(t *testing.T) {
	logger, err := log.New(log.Config{
		Level:  log.LevelDebug,
		Format: "text",
	}, "config", os.Stdout)
	assert.NilError(t, err)
	srv, err := configuration.NewServer(
		http.ServerConfiguration{
			Listen: "127.0.0.1:8080",
		},
		&myConfigReqHandler{},
		logger,
	)
	assert.NilError(t, err)
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

	client, err := configuration.NewClient(configuration.ClientConfig{
		http.ClientConfiguration{
			URL: "http://127.0.0.1:8080",
		},
	}, logger)
	assert.NilError(t, err)

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
	assert.NilError(t, err)
	assert.Equal(t, "yourcompany/yourimage", config.DockerRun.Config.ContainerConfig.Image)

	lifecycle.Stop(context.Background())
	err = lifecycle.Wait()
	assert.NilError(t, err)
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
