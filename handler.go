package configuration

// ConfigRequestHandler is a generic interface for simplified configuration request handling.
type ConfigRequestHandler interface {
	// OnConfig handles configuration requests. It should respond with either an error, resulting in a HTTP 500 response
	// code, or an AppConfig struct, which will be passed back to the client.
	OnConfig(request ConfigRequest) (AppConfig, error)
}
