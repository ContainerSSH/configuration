package configuration_test

import (
	"bufio"
	"bytes"
	"context"
	"testing"

	"github.com/containerssh/structutils"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"

	"github.com/containerssh/configuration"
)

func TestSaveLoad(t *testing.T) {
	config := &configuration.AppConfig{}
	newCfg := &configuration.AppConfig{}
	structutils.Defaults(config)

	buf := &bytes.Buffer{}
	writer := bufio.NewWriter(buf)
	reader := bufio.NewReader(buf)


	saver, err := configuration.NewWriterSaver(writer)
	assert.NoError(t, err)

	loader, err := configuration.NewReaderLoader(reader)
	assert.Nil(t, err, "failed to create reader (%v)", err)

	err = saver.Save(config)
	assert.Nil(t, err, "failed to load config (%v)", err)

	err = loader.Load(context.Background(), newCfg)
	assert.Nil(t, err, "failed to load config (%v)", err)

	diff := cmp.Diff(config, newCfg, cmpopts.EquateEmpty())
	assert.Empty(t, diff)
}
