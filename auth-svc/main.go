package main

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/metamatex/metamatemono/auth-svc/pkg"
	"github.com/metamatex/metamatemono/gen/v0/sdk/transport"
	"github.com/metamatex/metamatemono/gen/v0/sdk/transport/services/auth"
	"net/http"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() (err error) {
	c := transport.NewHttpJsonClient(transport.HttpJsonClientOpts{HttpClient: &http.Client{}, Token: "", Addr: "metamate:80"})

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	svc, err := pkg.NewService(pkg.ServiceOpts{
		Client:     c,
		Salt:       "hi",
		PrivateKey: privateKey,
	})
	if err != nil {
		return
	}

	s := auth.NewHttpJsonServer(auth.HttpJsonServerOpts{Service: svc})

	println("bra")

	err = http.ListenAndServe(":80", s)
	if err != nil {
		return
	}

	return
}
