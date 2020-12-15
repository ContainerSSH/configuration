# Changelog

## 0.9.3: Metrics integration

This release integrates the [metrics library](https://github.com/containerssh/metrics) which is now required as a dependency when creating a HTTP client.

## 0.9.2: Bumped HTTP dependency

This release bumps the [http](https://github.com/containerssh/http) dependency to 0.9.2 and changes the `Url` to `URL` in the setting.

## 0.9.1 Disabling configuration client

This release restores the ContainerSSH 0.3 functionality where passing an empty `url` in the configuration would disable fetching configuration from the config server.

## 0.9.0 Initial Release

This is the initial port from ContainerSSH 0.3. It is able to load ContainerSSH 0.3 configuration files, but deprecates the `listen` option in the root configuration and instead moves it to `ssh` → `listen`.