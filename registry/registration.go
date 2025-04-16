package registry
 
type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

type ServiceName string

const(
	logService = ServiceName("logService")
)
