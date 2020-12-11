package configuration_test

import (
	"context"
	"os"
	"testing"

	"github.com/containerssh/log"
	"github.com/stretchr/testify/assert"

	"github.com/containerssh/configuration"
)

// Test03Compatibility tests if a configuration file for ContainerSSH version 0.3 can be read.
func Test03Compatibility(t *testing.T) {
	logger, err := log.New(
		log.Config{
			Level:  log.LevelDebug,
			Format: "text",
		},
		"config",
		os.Stdout,
	)
	assert.NoError(t, err)

	logger.Infof("FYI: the deprecation notice in this test is intentional")

	testFile, err := os.Open("data/0.3.yaml")
	assert.NoError(t, err)
	reader, err := configuration.NewReaderLoader(
		testFile,
		logger,
		configuration.FormatYAML,
	)
	assert.NoError(t, err)

	config := configuration.AppConfig{}
	err = reader.Load(context.Background(), &config)
	assert.NoError(t, err)

	assert.Equal(t, "0.0.0.0:2222", config.SSH.Listen)
}
