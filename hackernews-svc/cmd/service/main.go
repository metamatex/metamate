package main

import (
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk/transport/services/hackernews"
	"github.com/metamatex/metamate/hackernews-svc/pkg"
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

	s := hackernews.NewHttpJsonServer(hackernews.HttpJsonServerOpts{Service: svc})

	err = http.ListenAndServe(":80", s)
	if err != nil {
		return
	}

	return
}
