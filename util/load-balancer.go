package util

import (
	"fmt"
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

func main() {
	var rws RandomWeightedSelector
	rws.AddEndpoint(Endpoint{Weight: 1, URL: "https://web1.example.com"})
	rws.AddEndpoint(Endpoint{Weight: 2, URL: "https://web2.example.com"})

	count1 := 0
	count2 := 0

	for i := 0; i < 100; i++ {
		switch rws.Select().URL {
		case "https://web1.example.com":
			count1++
		case "https://web2.example.com":
			count2++
		}
	}
	fmt.Println("Times web1: ", count1)
	fmt.Println("Times web2: ", count2)
}
