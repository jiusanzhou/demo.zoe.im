package main

import "net/http"

// cloneRequest makes a duplicate of the request.
func CloneRequest(r *http.Request) *http.Request {
	rc := new(http.Request)
	*rc = *r
	rc.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		rc.Header[k] = append([]string(nil), s...)
	}
	return rc
}