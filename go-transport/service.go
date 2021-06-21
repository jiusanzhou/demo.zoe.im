package main

import (
	"net/http"
)

// Service defines for
type Service struct {
	Name      string      // match connect host
	Schemas   []string    // match the schema
	Endpoints []*Endpoint // endpoints for connect

	matchs []Filter // if matchs
	Base   http.RoundTripper

	current int
}

type ServiceOption func(s *Service)

func DefaultTransport(base http.RoundTripper) ServiceOption {
	return func(s *Service) {
		s.Base = base
	}
}

func (s *Service) Init(opts ...ServiceOption) error {

	for _, o := range opts {
		o(s)
	}

	// make sure we have a noopper always return true
	s.matchs = append(
		s.matchs,
		filterNooper(),
		filterSchemas(s.Schemas),
		filterNamed(s.Name, s.Endpoints),
	)

	return nil
}

func (s *Service) Match(r *http.Request) bool {

	// make sure we have a noopper always return true

	for _, f := range s.matchs {
		if !f(r) {
			return false
		}
	}

	return true
}

func (s *Service) RoundTrip(r *http.Request) (*http.Response, error) {

	ed, err := s.takeEndpoint()
	if err != nil {
		return nil, err
	}

	err = ed.Apply(r)
	if err != nil {
		return nil, err
	}

	// TODO: there are 2 things we need do
	// let's check out the result to mark endpoint failed
	// and retry the request if need

	return s.base().RoundTrip(r)
}

func (s *Service) takeEndpoint() (*Endpoint, error) {
	// TODO: choose a healthy one to apply
	return s.Endpoints[0], nil
}

func (s *Service) base() http.RoundTripper {
	if s.Base != nil {
		return s.Base
	}
	return http.DefaultTransport
}


// TODO:

// string with 2 layout:
// 1. with service name: service_a=http://127.0.0.1:8080,http://127.0.0.1:8081
// 	  connect with service_a proxy to the endpoints
// 2. without service name: http://127.0.0.1:8080,http://127.0.0.1:8081
//    connect with any host of endpoints proxy to endponts