package main

import (
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/mql"
	"github.com/metamatex/metamate/hackernews-svc/pkg"
	"log"
	"net/http"
)

func main() {
	s := mql.NewHackernewsServer(mql.HackernewsServerOpts{
		Service: pkg.NewService(&http.Client{}),
	})

	err := http.ListenAndServe("0.0.0.0:80", s)
	if err != nil {
	    log.Fatal(err)
	}
}
