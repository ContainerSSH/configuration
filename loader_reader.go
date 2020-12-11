package configuration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/containerssh/log"
	"gopkg.in/yaml.v3"
)

// NewReaderLoader loads YAML files from reader.
func NewReaderLoader(
	reader io.Reader,
	logger log.Logger,
	format Format,
) (Loader, error) {
	if err := format.Validate(); err != nil {
		return nil, err
	}
	return &readerLoader{
		reader: reader,
		logger: logger,
		format: format,
	}, nil
}

type readerLoader struct {
	reader io.Reader
	logger log.Logger
	format Format
}

func (y *readerLoader) Load(_ context.Context, config *AppConfig) (err error) {
	switch y.format {
	case FormatYAML:
		err = y.loadYAML(y.reader, config)
	case FormatJSON:
		err = y.loadJSON(y.reader, config)
	default:
		err = fmt.Errorf("invalid format: %s", y.format)
	}
	if err != nil {
		return err
	}
	return config.FixCompatibility(y.logger)
}

func (y *readerLoader) LoadConnection(
	_ context.Context,
	_ string,
	_ net.TCPAddr,
	_ string,
	_ *AppConfig,
) error {
	return nil
}

func (y *readerLoader) loadYAML(reader io.Reader, config *AppConfig) error {
	decoder := yaml.NewDecoder(reader)
	decoder.KnownFields(true)
	return decoder.Decode(config)
}

func (y *readerLoader) loadJSON(reader io.Reader, config *AppConfig) error {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	return decoder.Decode(config)
}
