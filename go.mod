module github.com/containerssh/configuration/v2

go 1.16

require (
	github.com/containerssh/auditlog v1.0.0
	github.com/containerssh/auth v1.0.1
	github.com/containerssh/docker v1.0.0 // indirect
	github.com/containerssh/docker/v2 v2.0.1
	github.com/containerssh/geoip v1.0.0
	github.com/containerssh/health v1.0.1
	github.com/containerssh/http v1.1.0
	github.com/containerssh/kubernetes/v2 v2.0.1
	github.com/containerssh/log v1.1.6
	github.com/containerssh/metrics v1.0.0
	github.com/containerssh/security v1.0.0
	github.com/containerssh/service v1.0.0
	github.com/containerssh/sshproxy v1.0.0
	github.com/containerssh/sshserver v1.0.0
	github.com/containerssh/structutils v1.0.0
	github.com/docker/docker v20.10.18+incompatible
	github.com/google/go-cmp v0.5.5
	github.com/stretchr/testify v1.7.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

// Fixes CVE-2020-9283
replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 => golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/crypto v0.0.0-20200220183623-bac4c82f6975 => golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 => golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
)

// Fixes CVE-2020-14040
replace (
	golang.org/x/text v0.3.0 => golang.org/x/text v0.3.3
	golang.org/x/text v0.3.1 => golang.org/x/text v0.3.3
	golang.org/x/text v0.3.2 => golang.org/x/text v0.3.3
)

// Fixes CVE-2019-11254
replace (
	gopkg.in/yaml.v2 v2.2.0 => gopkg.in/yaml.v2 v2.2.8
	gopkg.in/yaml.v2 v2.2.1 => gopkg.in/yaml.v2 v2.2.8
	gopkg.in/yaml.v2 v2.2.2 => gopkg.in/yaml.v2 v2.2.8
	gopkg.in/yaml.v2 v2.2.3 => gopkg.in/yaml.v2 v2.2.8
	gopkg.in/yaml.v2 v2.2.4 => gopkg.in/yaml.v2 v2.2.8
	gopkg.in/yaml.v2 v2.2.5 => gopkg.in/yaml.v2 v2.2.8
	gopkg.in/yaml.v2 v2.2.6 => gopkg.in/yaml.v2 v2.2.8
	gopkg.in/yaml.v2 v2.2.7 => gopkg.in/yaml.v2 v2.2.8
)
