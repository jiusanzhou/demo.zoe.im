package main

import "net/http"

// LBTransport is the load balancer implementation of Transport.
type LBTransport struct {
	Base http.RoundTripper

	Services []*Service
}

func (lb *LBTransport) base() http.RoundTripper {
	if lb.Base != nil {
		return lb.Base
	}
	return http.DefaultTransport
}

// RoundTrip must always close the body, including on errors,
// but depending on the implementation may do so in a separate
// goroutine even after RoundTrip returns. This means that
// callers wanting to reuse the body for subsequent requests
// must arrange to wait for the Close call before doing so.
//
// The Request's URL and Header fields must be initialized.
func (lb *LBTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	// TODO: should we need to check the result of service is ok ?
	// do this in the service

	// choose a lb client for real connect

	for _, s := range lb.Services {
		if s.Match(req) {
			return s.RoundTrip(req)
		}
	}

	// on service provider use default round tripper
	return lb.base().RoundTrip(req)
}

func (lb *LBTransport) Init() error {
	// TODO: add more option

	for _, s := range lb.Services {
		// TODO: if init return error, remove current service
		_ = s.Init(DefaultTransport(lb.base()))
	}

	return nil
}

// NewLBTransport return a new transport from config
func NewLBTransport() (*LBTransport, error) {

	// TODO: add more options

	// load service and init service

	lb := &LBTransport{}

	return lb, nil
}
