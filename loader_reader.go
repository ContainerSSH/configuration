package configuration

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net"

	"gopkg.in/yaml.v3"
)

// NewReaderLoader loads YAML files from reader.
func NewReaderLoader(reader io.Reader) (Loader, error) {
	return &readerLoader{
		reader: reader,
	}, nil
}

type readerLoader struct {
	reader io.Reader
}

func (y *readerLoader) Load(_ context.Context, config *AppConfig) error {
	data, err := ioutil.ReadAll(y.reader)
	if err != nil {
		return err
	}
	return y.loadYAML(data, config)
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

func (y *readerLoader) loadYAML(data []byte, config *AppConfig) error {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)
	return decoder.Decode(config)
}
