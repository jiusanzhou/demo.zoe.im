package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)


var te = &LBTransport{
	Services: []*Service{
		&Service{
			Name: "demo",
			Endpoints: []*Endpoint{
				&Endpoint{
					URL: &url.URL{
						Host: "www.google.com",
					},
				},
			},
		},
	},
}

func main() {
	fmt.Println("OK!")

	http.DefaultClient.Transport = te





	resp, err := http.Get("http://demo")
	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("[SUCCESS]")

	io.Copy(os.Stdout, resp.Body)
}
