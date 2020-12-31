# Changelog

## 0.9.7: Bumped Docker, Kubernetes and Metrics dependencies

This release updates the Docker, Kubernetes, and Metrics dependency to the latest releases.

## 0.9.6: Added Validate()

This release adds a `Validate()` method to the configuration that allows for check the configuration on loading.

## 0.9.5: Security configuration

This release adds support for the new [security library](https://github.com/containerssh/security). It also adds a metric for config server requests and updates several libraries to their latest versions.

## 0.9.4: New docker and kubernetes backends

This release adds support for the new [docker](https://github.com/containerssh/docker) and [kubernetes](https://github.com/containerssh/kubernetes) backends.

## 0.9.3: Metrics integration

This release integrates the [metrics library](https://github.com/containerssh/metrics) which is now required as a dependency when creating a HTTP client.

## 0.9.2: Bumped HTTP dependency

This release bumps the [http](https://github.com/containerssh/http) dependency to 0.9.2 and changes the `Url` to `URL` in the setting.

## 0.9.1 Disabling configuration client

This release restores the ContainerSSH 0.3 functionality where passing an empty `url` in the configuration would disable fetching configuration from the config server.

## 0.9.0 Initial Release

This is the initial port from ContainerSSH 0.3. It is able to load ContainerSSH 0.3 configuration files, but deprecates the `listen` option in the root configuration and instead moves it to `ssh` → `listen`.