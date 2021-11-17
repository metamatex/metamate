package embedded

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/metamatex/metamate/reddit-svc/pkg"
	"github.com/metamatex/metamate/reddit-svc/pkg/communication"
	"net/http"
)

func init() {
	handler[Reddit] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, vSvc types.EmbeddedSvc) (h http.Handler, err error) {
		client, err := communication.NewClient(communication.ClientOpts{Client: &http.Client{}, UserAgent: "mql"})
		if err != nil {
		    return
		}

		svc := pkg.NewService(pkg.ServiceOpts{
			Client: client,
		})

		h = mql.NewRedditServer(mql.RedditServerOpts{Service: svc})

		return
	}
}
