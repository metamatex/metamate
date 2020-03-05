package main

import (
	"github.com/metamatex/asg/pkg/v0/asg"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamatemono/pkg/services/sqlx-svc/pkg/boot"
	"log"
)

func main() {
	err := run()
	if err != nil {
	    log.Fatal(err)
	}
}

func run() (err error) {
	rn, err := asg.New()
	if err != nil {
	    return
	}

	f := generic.NewFactory(rn)

	c := boot.NewTestConfig()

	d, err := boot.NewDependencies(rn, f, c)
	if err != nil {
	    return
	}

	server := httpjson.NewServer(rn, f, d.ServeFunc, "0.0.0.0:80")

	err = server.Listen()
	if err != nil {
	    return
	}

	return
}