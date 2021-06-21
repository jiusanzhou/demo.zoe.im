package main

import (
	"net/http"
	"strings"
)



type Filter func(r *http.Request) bool

func filterNooper() Filter {
	return func(_ *http.Request) bool {
		return true
	}
}

func filterSchemas(oschms []string) Filter {
	if len(oschms) == 0 {
		return filterNooper()
	}

	// turn to upper
	schms := make([]string, len(oschms))
	for i, s := range oschms {
		schms[i] = strings.ToUpper(s)
	}

	return func(r *http.Request) bool {
		for _, s := range schms {
			if s == r.Method {
				return true
			}
		}
		return false
	}
}

func filterNamed(name string, endpoints []*Endpoint) Filter {
	return func(r *http.Request) bool {
		if name != "" {
			return r.URL.Host == name
		}

		// check if match the any endpoint
		for _, e := range endpoints {
			return e.URL.Host == r.Host
		}

		// at last return false
		return false
	}
}