package configuration

import (
	"fmt"

	"github.com/containerssh/auditlog"
	"github.com/containerssh/auth"
	"github.com/containerssh/docker/v2"
	"github.com/containerssh/geoip"
	"github.com/containerssh/health"
	"github.com/containerssh/kubernetes/v2"
	"github.com/containerssh/log"
	"github.com/containerssh/metrics"
	"github.com/containerssh/security"
	"github.com/containerssh/sshproxy"
	sshserver "github.com/containerssh/sshserver/v2"
)

// AppConfig is the root configuration object of ContainerSSH.
//goland:noinspection GoDeprecation
type AppConfig struct {
	// Listen is an alias for ssh.listen. Its usage is deprecated.
	// swagger:ignore
	// deprecated: use SSH.Listen instead
	Listen string `json:"listen,omitempty" yaml:"listen,omitempty" default:""`
	// SSH contains the configuration for the SSH server.
	// swagger:ignore
	SSH sshserver.Config `json:"ssh" yaml:"ssh"`
	// ConfigServer contains the settings for fetching the user-specific configuration.
	// swagger:ignore
	ConfigServer ClientConfig `json:"configserver" yaml:"configserver"`
	// Auth contains the configuration for user authentication.
	// swagger:ignore
	Auth auth.ClientConfig `json:"auth" yaml:"auth"`
	// Log contains the configuration for the logging level.
	// swagger:ignore
	Log log.Config `json:"log" yaml:"log"`
	// Metrics contains the configuration for the metrics server.
	// swagger:ignore
	Metrics metrics.Config `json:"metrics" yaml:"metrics"`
	// GeoIP contains the configuration for the GeoIP lookups.
	// swagger:ignore
	GeoIP geoip.Config `json:"geoip" yaml:"geoip"`
	// Audit contains the configuration for audit logging and log upload.
	// swagger:ignore
	Audit auditlog.Config `json:"audit" yaml:"audit"`
	// Health contains the configuration for the health check service.
	Health health.Config `json:"health" yaml:"health"`

	// Security contains the security restrictions on what can be executed. This option can be changed from the config
	// server.
	Security security.Config `json:"security" yaml:"security"`
	// Backend defines which backend to use. This option can be changed from the config server.
	Backend string `json:"backend" yaml:"backend" default:"docker"`
	// Docker contains the configuration for the docker backend. This option can be changed from the config server.
	Docker docker.Config `json:"docker,omitempty" yaml:"docker"`
	// DockerRun contains the configuration for the deprecated dockerrun backend. This option can be changed from the
	// config server.
	// deprecated: use Docker instead
	DockerRun docker.DockerRunConfig `json:"dockerrun,omitempty" yaml:"dockerrun"`
	// Kubernetes contains the configuration for the kubernetes backend. This option can be changed from the config
	// server.
	Kubernetes kubernetes.Config `json:"kubernetes,omitempty" yaml:"kubernetes"`
	// KubeRun contains the configuration for the deprecated kuberun backend. This option can be changed from the config
	// server.
	// deprecated: use Kubernetes instead
	KubeRun kubernetes.KubeRunConfig `json:"kuberun,omitempty" yaml:"kuberun"`
	// SSHProxy is the configuration for the SSH proxy backend, which forwards requests to a backing SSH server.
	SSHProxy sshproxy.Config `json:"sshproxy,omitempty" yaml:"sshproxy"`
}

// FixCompatibility moves deprecated options to their new places and issues warnings.
func (cfg *AppConfig) FixCompatibility(logger log.Logger) error {
	if cfg.Listen != "" {
		if cfg.SSH.Listen == "" || cfg.SSH.Listen == "0.0.0.0:2222" {
			logger.Warning(log.NewMessage(
				WListenDeprecated,
				"You are using the 'listen' option deprecated in ContainerSSH 0.4. Please use the new 'ssh -> listen' option. See https://containerssh.io/deprecations/listen for details.",
			))
			cfg.SSH.Listen = cfg.Listen
			cfg.Listen = ""
		} else {
			logger.Warning(log.NewMessage(
				WListenDeprecated,
				"You are using the 'listen' option deprecated in ContainerSSH 0.4 as well as the new 'ssh -> listen' option. The new option takes precedence. Please see https://containerssh.io/deprecations/listen for details.",
			))
		}
	}
	return nil
}

// Validate validates the configuration structure and returns an error if it is invalid.
//
// - dynamic enables the validation for dynamically configurable options.
func (cfg *AppConfig) Validate(dynamic bool) error {
	queue := newValidationQueue()
	queue.add("SSH", &cfg.SSH)
	queue.add("config server", &cfg.ConfigServer)
	queue.add("authentication", &cfg.Auth)
	queue.add("logging", &cfg.Log)
	queue.add("metrics", &cfg.Metrics)
	queue.add("GeoIP", &cfg.GeoIP)
	queue.add("audit log", &cfg.Audit)

	if cfg.ConfigServer.URL != "" && !dynamic {
		return queue.Validate()
	}
	queue.add("security configuration", &cfg.Security)
	switch cfg.Backend {
	case "docker":
		queue.add("Docker", &cfg.Docker)
	case "dockerrun":
		queue.add("DockerRun", &cfg.DockerRun)
	case "kubernetes":
		queue.add("Kubernetes", &cfg.Kubernetes)
	case "kuberun":
		queue.add("KubeRun", &cfg.KubeRun)
	case "sshproxy":
		queue.add("SSH proxy", &cfg.SSHProxy)
	case "":
		return fmt.Errorf("no backend configured")
	default:
		return fmt.Errorf("invalid backend: %s", cfg.Backend)
	}
	return queue.Validate()
}

type validatable interface {
	Validate() error
}

func newValidationQueue() *validationQueue {
	return &validationQueue{
		items: map[string]validatable{},
	}
}

type validationQueue struct {
	items map[string]validatable
}

func (v *validationQueue) add(name string, item validatable) {
	v.items[name] = item
}

func (v *validationQueue) Validate() error {
	for name, item := range v.items {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("invalid %s configuration (%w)", name, err)
		}
	}
	return nil
}
