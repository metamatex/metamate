package main

import (
	"fmt"
	"github.com/metamatex/metamate/metamate/pkg/v0/boot"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	version = "dev-0.0.0"
	commit  = "dev"
	date    = "dev"
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

	d, err := boot.NewDependencies(c, types.Version{Version: version, Commit: commit, Date: date})
	if err != nil {
		return
	}

	fmt.Printf("version: %v\nvcommit: %v\ndate: %v\n\n", version, commit, date)

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
