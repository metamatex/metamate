package main

import (
	"github.com/metamatex/metamatemono/gen/v0/sdk/transport"
	"github.com/metamatex/metamatemono/pkg/services/mastodon-svc/pkg"
	"net/http"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() (err error) {
	svc := pkg.NewService()

	s := transport.NewHttpJsonServer(transport.HttpJsonServerOpts{Service: svc})

	err = http.ListenAndServe(":80", s)
	if err != nil {
		return
	}

	return
}
