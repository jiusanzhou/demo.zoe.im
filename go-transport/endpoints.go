package main

import (
	"net/http"
	"net/url"
)

// Endpoint contains the http endpoint
//
// Which can be use a a balancer
type Endpoint struct {
	rawurl string

	URL *url.URL

	// TODO: add more config
	// timeout: connect_timeout, write_timeout, read_timeout
	// max_fails
	// fail_timeout
	// weight
	// is_backup

	// state of current endpont
	countReq int64
	countErr int64
}

func (ep *Endpoint) Apply(r *http.Request) error {

	if ep.URL.Scheme != "" {
		r.URL.Scheme = ep.URL.Scheme
	}
	if ep.URL.Host != "" {
		r.URL.Host = ep.URL.Host
	}
	if ep.URL.User != nil {
		r.URL.User = ep.URL.User
	}

	return nil
}

// NewEndpoint parse from string
func NewEndpoint(rawurl string) (ep *Endpoint, err error) {
	ep = &Endpoint{
		rawurl: rawurl,
	}

	ep.URL, err = url.Parse(rawurl)
	if err != nil {
		return
	}

	return
}