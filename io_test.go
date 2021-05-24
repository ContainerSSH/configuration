package configuration_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/containerssh/docker/v2"
	"github.com/containerssh/http"
	"github.com/containerssh/kubernetes/v2"
	"github.com/containerssh/log"
	"github.com/containerssh/structutils"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"

	"github.com/containerssh/configuration/v2"
)

func TestSaveLoadYAML(t *testing.T) {
	testSaveLoad(t, configuration.FormatYAML)
}

func TestSaveLoadJSON(t *testing.T) {
	testSaveLoad(t, configuration.FormatJSON)
}

func testSaveLoad(t *testing.T, format configuration.Format) {
	// region Setup
	logger := log.NewTestLogger(t)

	config := &configuration.AppConfig{}
	newCfg := &configuration.AppConfig{}
	structutils.Defaults(config)

	config.Auth.URL = "http://localhost:8080"

	buf := &bytes.Buffer{}
	// endregion

	// region Save
	saver, err := configuration.NewWriterSaver(
		buf,
		logger,
		format,
	)
	assert.NoError(t, err)
	err = saver.Save(config)
	assert.Nil(t, err, "failed to load config (%v)", err)
	// endregion

	// region Load
	loader, err := configuration.NewReaderLoader(buf, logger, format)
	assert.Nil(t, err, "failed to create reader (%v)", err)
	err = loader.Load(context.Background(), newCfg)
	assert.Nil(t, err, "failed to load config (%v)", err)
	// endregion

	// region Assert
	config.Listen = ""

	diff := cmp.Diff(
		config,
		newCfg,
		cmp.AllowUnexported(http.ServerConfiguration{}),
		cmp.AllowUnexported(http.ClientConfiguration{}),
		cmp.AllowUnexported(configuration.ClientConfig{}),
		cmp.AllowUnexported(kubernetes.PodConfig{}),
		cmp.AllowUnexported(kubernetes.ConnectionConfig{}),
		cmp.AllowUnexported(docker.ExecutionConfig{}),
		cmp.AllowUnexported(log.SyslogConfig{}),
		cmpopts.EquateEmpty(),
	)
	assert.Empty(t, diff)
	// endregion
}
