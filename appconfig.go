package configuration

import (
	"github.com/containerssh/auditlog"
	"github.com/containerssh/auth"
	"github.com/containerssh/docker"
	"github.com/containerssh/geoip"
	"github.com/containerssh/kubernetes"
	"github.com/containerssh/log"
	"github.com/containerssh/metrics"
	"github.com/containerssh/security"
	"github.com/containerssh/sshserver"
)

// AppConfig is the root configuration object of ContainerSSH.
//goland:noinspection GoDeprecation
type AppConfig struct {
	// Listen is an alias for ssh.listen. Its usage is deprecated.
	Listen string `json:"listen,omitempty" yaml:"listen,omitempty" default:""`
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

	// Security contains the security restrictions on what can be executed.
	Security security.Config `json:"security" yaml:"security" comment:"Security configuration"`

	// Backend defines which backend to use. This option can be changed from the config server.
	Backend string `json:"backend" yaml:"backend" default:"dockerrun" comment:"Backend module to use"`
	// Docker contains the configuration for the docker backend.
	Docker docker.Config `json:"docker,omitempty" yaml:"docker" comment:"Docker configuration to use when the Docker backend is used."`
	// DockerRun contains the configuration for the deprecated dockerrun backend.
	DockerRun docker.DockerRunConfig `json:"dockerrun,omitempty" yaml:"dockerrun" comment:"Docker configuration to use when the Docker run backend is used."`
	// Kubernetes contains the configuration for the kubernetes backend.
	Kubernetes kubernetes.Config `json:"kubernetes,omitempty" yaml:"kubernetes" comment:"Kubernetes configuration to use when the kubernetes run backend is used."`
	// KubeRun contains the configuration for the deprecated kuberun backend.
	KubeRun kubernetes.KubeRunConfig `json:"kuberun,omitempty" yaml:"kuberun" comment:"Kubernetes configuration to use when the kuberun run backend is used."`
}

// FixCompatibility moves deprecated options to their new places and issues warnings.
func (cfg *AppConfig) FixCompatibility(logger log.Logger) error {
	if cfg.Listen != "" {
		if cfg.SSH.Listen == "" || cfg.SSH.Listen == "0.0.0.0:2222" {
			logger.Warningf("You are using the 'listen' option deprecated in ContainerSSH 0.4. Please use the new 'ssh -> listen' option. See https://containerssh.io/deprecations/listen for details.")
			cfg.SSH.Listen = cfg.Listen
			cfg.Listen = ""
		} else {
			logger.Warningf("You are using the 'listen' option deprecated in ContainerSSH 0.4 as well as the new 'ssh -> listen' option. The new option takes precedence. Please see https://containerssh.io/deprecations/listen for details.")
		}
	}
	return nil
}
