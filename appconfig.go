package configuration

import (
	"github.com/containerssh/auditlog"
	"github.com/containerssh/auth"
	"github.com/containerssh/dockerrun"
	"github.com/containerssh/geoip"
	"github.com/containerssh/kuberun"
	"github.com/containerssh/log"
	"github.com/containerssh/metrics"
	"github.com/containerssh/sshserver"
)

// AppConfig is the root configuration object of ContainerSSH.
type AppConfig struct {
	// SSH contains the configuration for the SSH server.
	// swagger:ignore
	SSH sshserver.Config `json:"ssh" yaml:"ssh" comment:"SSH configuration"`
	// ConfigServer contains the settings for fetching the user-specific configuration.
	// swagger:ignore
	ConfigServer ClientConfig `json:"configserver" yaml:"configserver" comment:"Configuration server settings"`
	// Auth contains the configuration for user authentication.
	// swagger:ignore
	Auth auth.ClientConfig `json:"auth" yaml:"auth" comment:"Authentication server configuration"`
	// Log contains the configuration for the logging level.
	// swagger:ignore
	Log log.Config `json:"log" yaml:"log" comment:"Log configuration"`
	// Metrics contains the configuration for the metrics server.
	// swagger:ignore
	Metrics metrics.Config `json:"metrics" yaml:"metrics" comment:"Metrics configuration."`
	// GeoIP contains the configuration for the GeoIP lookups.
	// swagger:ignore
	GeoIP geoip.Config `json:"geoip" yaml:"geoip" comment:"GeoIP database"`
	// Audit contains the configuration for audit logging and log upload.
	// swagger:ignore
	Audit auditlog.Config `json:"audit" yaml:"audit" comment:"Audit configuration"`

	// Backend defines which backend to use. This option can be changed from the config server.
	Backend string `json:"backend" yaml:"backend" default:"dockerrun" comment:"Backend module to use"`
	// DockerRun contains the configuration for the dockerrun backend.
	DockerRun dockerrun.Config `json:"dockerrun" yaml:"dockerrun" comment:"Docker configuration to use when the Docker run backend is used."`
	// KubeRun contains the configuration for the kuberun backend.
	KubeRun kuberun.Config `json:"kuberun" yaml:"kuberun" comment:"Kubernetes configuration to use when the Kubernetes run backend is used."`
}
