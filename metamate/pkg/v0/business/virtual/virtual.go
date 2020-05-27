package virtual

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

const (
	Pipe       = "pipe"
	ReqFilter  = "reqFilter"
	Error      = "error"
	Pagination = "pagination"
	Mastodon   = "mastodon"
	Reddit     = "reddit"
	Hackernews = "hackernews"
	Kubernetes = "kubernetes"
)

var handler = map[string]func(f generic.Factory, rn *graph.RootNode, c *http.Client, svc types.VirtualSvc) (http.Handler, error){}
