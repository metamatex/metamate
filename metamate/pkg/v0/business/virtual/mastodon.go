// +build !lite

package virtual

import (
	"errors"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/transport/services/mastodon"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/mastodon-svc/pkg"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[Mastodon] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, opts types.VirtualSvcOpts) (h http.Handler, t string, err error) {
		err = validateMastodonOpts(opts)
		if err != nil {
			return
		}

		svc := pkg.NewService(pkg.ServiceOpts{
			Host:         opts.Mastodon.Host,
			ClientId:     opts.Mastodon.ClientId,
			ClientSecret: opts.Mastodon.ClientSecret,
		})

		h = mastodon.NewHttpJsonServer(mastodon.HttpJsonServerOpts{Service: svc})

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
