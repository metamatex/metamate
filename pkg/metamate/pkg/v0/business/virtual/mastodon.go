// +build !lite

package virtual

import (
	"errors"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/gen/v0/sdk/transport"
	"github.com/metamatex/metamatemono/pkg/services/mastodon-svc/pkg"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/pkg/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[Mastodon] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, opts types.VirtualSvcOpts) (h http.Handler, t string, err error) {
		err = validateMastodonOpts(opts)
		if err != nil {
			return
		}

		svc := pkg.NewService(pkg.ServiceOpts{
			Host: opts.Mastodon.Host,
			ClientId: opts.Mastodon.ClientId,
			ClientSecret: opts.Mastodon.ClientSecret,
		})

		h = transport.NewHttpJsonServer(transport.HttpJsonServerOpts{Service: svc})

		t = sdk.ServiceTransport.HttpJson

		return
	}
}

func validateMastodonOpts(opts types.VirtualSvcOpts) (err error) {
	if opts.Mastodon == nil {
		err = errors.New("opts.Mastodon needs to be set")

		return
	}

	if opts.Mastodon.Host == "" {
		err = errors.New("opts.Mastodon.Host needs to be set")

		return
	}

	if opts.Mastodon.ClientId == "" {
		err = errors.New("opts.Mastodon.ClientId needs to be set")

		return
	}

	if opts.Mastodon.ClientSecret == "" {
		err = errors.New("opts.Mastodon.ClientSecret needs to be set")

		return
	}

	return
}
