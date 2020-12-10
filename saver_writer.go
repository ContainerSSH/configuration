package configuration

import (
	"io"

	"gopkg.in/yaml.v3"
)

// NewWriterSaver creates a config saver that writes the data in YAML format to the specified writer.
func NewWriterSaver(writer io.Writer) (ConfigSaver, error) {
	return &writerSaver{
		writer: writer,
	}, nil
}

type writerSaver struct {
	writer io.Writer
}

func (w *writerSaver) Save(config *AppConfig) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	_, err = w.writer.Write(data)
	if err != nil {
		return err
	}
	return nil
}
