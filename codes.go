package configuration

const (
	// MConfig is a message indicating a configuration server request.
	MConfig = "CONFIG_REQUEST"
	// EConfigBackendError is a message indicating that the configuration server responded with an invalid response.
	EConfigBackendError = "CONFIG_BACKEND_ERROR"
	// MConfigSuccess is a message indicating that the config server has responded successfully.
	MConfigSuccess = "CONFIG_RESPONSE"
	// EInvalidStatus indicates that the config server responded with a non-200 status code.
	EInvalidStatus = "CONFIG_INVALID_STATUS_CODE"
)
