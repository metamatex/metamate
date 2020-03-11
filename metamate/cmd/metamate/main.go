package main

import (
	"fmt"
	"github.com/metamatex/metamatemono/metamate/pkg/v0/boot"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go func() {
		err := run()
		if err != nil {
			panic(err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

func run() (err error) {
	c := boot.NewProdConfig()

	d, err := boot.NewDependencies(c)
	if err != nil {
		return
	}

	for _, r := range d.Routes {
		for _, m := range r.Methods {
			fmt.Printf("%v: %v:%v%v\n", m, c.Host.Bind, c.Host.HttpPort, r.Path)
		}
	}

	err = http.ListenAndServe(fmt.Sprintf(":%v", c.Host.HttpPort), d.Router)
	if err != nil {
		return
	}

	return
}
