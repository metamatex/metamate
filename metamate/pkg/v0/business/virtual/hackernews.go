package virtual

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk/transport/services/hackernews"
	"github.com/metamatex/metamate/hackernews-svc/pkg"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[Hackernews] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, vSvc types.VirtualSvc) (h http.Handler, t string, err error) {
		svc := pkg.NewService(&http.Client{})

		h = hackernews.NewHttpJsonServer(hackernews.HttpJsonServerOpts{Service: svc})

		t = sdk.ServiceTransport.HttpJson

		return
	}
}
