package configuration

type Config interface {
	Get() map[string]string
}

type EndpointCollection struct {
	elements map[string]string
}

func CreateFromCliParams(endpoints, destinations []string) *EndpointCollection {

	var endpointsWithDestinations = map[string]string{}

	for i := 0; i < len(endpoints); i++ {
		endpointsWithDestinations[endpoints[i]] = destinations[i]
	}

	return &EndpointCollection{endpointsWithDestinations}

}

func (c *EndpointCollection) Get() map[string]string {
	return c.elements
}
