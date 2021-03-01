package eureka

import (
	"math/rand"
)

type Endpoint struct {
	URL    string
	Weight int
}

type RandomWeightedSelector struct {
	max       int
	endpoints []Endpoint
}

func (rws *RandomWeightedSelector) AddEndpoint(endpoint Endpoint) {
	rws.endpoints = append(rws.endpoints, endpoint)
	rws.max += endpoint.Weight
}

func (rws *RandomWeightedSelector) Select() Endpoint {
	r := rand.Intn(rws.max)
	for _, endpoint := range rws.endpoints {
		if r < endpoint.Weight {
			return endpoint
		} else {
			r = r - endpoint.Weight
		}
	}
	// should never get to this point because r is smaller than max
	return Endpoint{}
}
