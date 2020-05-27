package virtual

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/metamatex/metamate/reddit-svc/pkg"
	types0 "github.com/metamatex/metamate/reddit-svc/pkg/types"
	"net/http"
)

func init() {
	handler[Reddit] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, vSvc types.VirtualSvc) (h http.Handler, err error) {
		svc, err := pkg.NewService(pkg.ServiceOpts{
			Client: &http.Client{},
			Credentials: types0.Credentials{
				ClientId:     vSvc.Opts.Reddit.ClientId,
				ClientSecret: vSvc.Opts.Reddit.ClientSecret,
				Username:     vSvc.Opts.Reddit.Username,
				Password:     vSvc.Opts.Reddit.Password,
			},
			UserAgent: vSvc.Opts.Reddit.UserAgent,
		})
		if err != nil {
			return
		}

		h = mql.NewMastodonServer(mql.MastodonServerOpts{Service: svc})

		return
	}
}
