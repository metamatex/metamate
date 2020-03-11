package virtual

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

const (
	Sqlx       = "sqlx"
	Pipe       = "pipe"
	ReqFilter  = "reqFilter"
	Auth       = "auth"
	Mastodon   = "mastodon"
	Kubernetes = "kubernetes"
)

var handler = map[string]func(f generic.Factory, rn *graph.RootNode, c *http.Client, opts types.VirtualSvcOpts) (http.Handler, string, error){}
