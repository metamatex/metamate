package virtual

import (
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/metamate/pkg/v0/types"
	"net/http"
)

const (
	Sqlx      = "sqlx"
	Pipe      = "pipe"
	ReqFilter = "reqFilter"
	Auth      = "auth"
	Mastodon  = "mastodon"
)

var handler = map[string]func(f generic.Factory, rn *graph.RootNode, c *http.Client, opts types.VirtualSvcOpts) (http.Handler, string, error){}
