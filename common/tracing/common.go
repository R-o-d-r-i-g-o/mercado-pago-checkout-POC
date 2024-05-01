package tracing

const Key = "traceID"

var (
	serviceName string
	environment string
	serviceID   string
)

// ServiceName retrieves then service's name for identification on tracing tool
func ServiceName() string {
	return serviceName
}

// Environment retrieves the current environment that runs the application
func Environment() string {
	return environment
}

// ServiceID retrieves then service's id to identify the container/process
func ServiceID() string {
	return serviceID
}
