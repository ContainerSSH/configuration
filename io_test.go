package configuration_test

import (
	"bufio"
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/containerssh/log"
	"github.com/containerssh/structutils"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"

	"github.com/containerssh/configuration"
)

func TestSaveLoadYAML(t *testing.T) {
	testSaveLoad(t, configuration.FormatYAML)
}

func TestSaveLoadJSON(t *testing.T) {
	testSaveLoad(t, configuration.FormatJSON)
}

func testSaveLoad(t *testing.T, format configuration.Format) {
	logger, err := log.New(
		log.Config{
			Level:  log.LevelDebug,
			Format: "text",
		},
		"config",
		os.Stdout,
	)
	assert.NoError(t, err)

	config := &configuration.AppConfig{}
	newCfg := &configuration.AppConfig{}
	structutils.Defaults(config)

	buf := &bytes.Buffer{}
	writer := bufio.NewWriter(buf)
	reader := bufio.NewReader(buf)

	saver, err := configuration.NewWriterSaver(
		writer,
		logger,
		format,
	)
	assert.NoError(t, err)

	loader, err := configuration.NewReaderLoader(reader, logger, format)
	assert.Nil(t, err, "failed to create reader (%v)", err)

	err = saver.Save(config)
	assert.Nil(t, err, "failed to load config (%v)", err)
	assert.NoError(t, writer.Flush())

	err = loader.Load(context.Background(), newCfg)
	assert.Nil(t, err, "failed to load config (%v)", err)

	// The Listen configuration is removed on purpose after load because it's a deprecated field.
	config.Listen = ""

	diff := cmp.Diff(config, newCfg, cmpopts.EquateEmpty())
	assert.Empty(t, diff)
}
