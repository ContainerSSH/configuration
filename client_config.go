package configuration

import (
	"github.com/containerssh/http"
)

//noinspection GoNameStartsWithPackageName
type ClientConfig struct {
	http.ClientConfiguration `json:",inline" yaml:",inline"`
}

// Validate validates the client configuration.
func (c *ClientConfig) Validate() error {
	if c.ClientConfiguration.URL == "" {
		return nil
	}
	return c.ClientConfiguration.Validate()
}
