package configuration

// Config configuration type for the api gateway
type Config interface {
	Get() map[string]string
}

// EndpointCollection implementation based on cli parameters
type EndpointCollection struct {
	elements map[string]string
}

// CreateFromCliParams creates new instance based on endpoints and destinations
func CreateFromCliParams(endpoints, destinations []string) *EndpointCollection {

	var endpointsWithDestinations = map[string]string{}

	for i := 0; i < len(endpoints); i++ {
		endpointsWithDestinations[endpoints[i]] = destinations[i]
	}

	return &EndpointCollection{endpointsWithDestinations}

}

// Get get the configuration
func (c *EndpointCollection) Get() map[string]string {
	return c.elements
}
