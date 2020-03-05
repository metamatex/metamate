package virtual

import (
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/pkg/metamate/pkg/v0/types"
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
