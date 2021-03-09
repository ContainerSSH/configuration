# Message / error codes

| Code | Explanation |
|------|-------------|
| `CONFIG_BACKEND_ERROR` | ContainerSSH has received an invalid response from the configuration server or the network connection broke. ContainerSSH will retry fetching the user-specific configuration until the timeout. If this error persists check the connectivity to the configuration server, or the logs of the configuration server itself to find out of there may be a specific error. |
| `CONFIG_INVALID_STATUS_CODE` | ContainerSSH has received a non-200 response code when calling a per-user backend configuration from the configuration server. |
| `CONFIG_REQUEST` | ContainerSSH is sending a quest to the configuration server to obtain a per-user backend configuration. |
| `CONFIG_RESPONSE` | ContainerSSH has received a per-user backend configuration from the configuration server. |
| `CONFIG_SERVER_AVAILABLE` | The ContainerSSH configuration server is now available at the specified address. |

