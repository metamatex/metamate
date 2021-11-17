package embedded

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/mastodon-svc/pkg"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[Mastodon] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, vSvc types.EmbeddedSvc) (h http.Handler, err error) {
		svc := pkg.NewService(pkg.ServiceOpts{})

		h = mql.NewMastodonServer(mql.MastodonServerOpts{Service: svc})

		return
	}
}