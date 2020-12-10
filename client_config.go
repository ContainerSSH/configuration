package configuration

import (
	"github.com/containerssh/http"
)

//noinspection GoNameStartsWithPackageName
type ClientConfig struct {
	http.ClientConfiguration `yaml:",inline"`
}
