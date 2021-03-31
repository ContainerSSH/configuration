module github.com/containerssh/configuration

go 1.14

require (
	github.com/aws/aws-sdk-go v1.38.9 // indirect
	github.com/containerssh/auditlog v0.9.9
	github.com/containerssh/auth v0.9.6
	github.com/containerssh/docker v0.9.12
	github.com/containerssh/geoip v0.9.4
	github.com/containerssh/http v0.9.9
	github.com/containerssh/kubernetes v0.9.9
	github.com/containerssh/log v0.9.13
	github.com/containerssh/metrics v0.9.8
	github.com/containerssh/security v0.9.8
	github.com/containerssh/service v0.9.3
	github.com/containerssh/sshproxy v0.9.0
	github.com/containerssh/sshserver v0.9.25
	github.com/containerssh/structutils v0.9.0
	github.com/docker/docker v20.10.5+incompatible
	github.com/google/go-cmp v0.5.5
	github.com/oschwald/geoip2-golang v1.5.0 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210330210036-cd0ac97f97f6 // indirect
	golang.org/x/oauth2 v0.0.0-20210323180902-22b0adad7558 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

// Exclude this package because it got renamed to /moby/ which breaks packages.
exclude github.com/docker/spdystream v0.2.0

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
