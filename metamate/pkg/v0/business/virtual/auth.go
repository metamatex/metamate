// +build !lite

package virtual

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/auth-svc/pkg"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/transport"
	"github.com/metamatex/metamate/gen/v0/sdk/transport/services/auth"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[Auth] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, opts types.VirtualSvcOpts) (h http.Handler, t string, err error) {
		err = validateAuthOpts(opts)
		if err != nil {
			return
		}

		cli := transport.NewHttpJsonClient(transport.HttpJsonClientOpts{
			HttpClient: c,
			Token:      "",
			Addr:       "http://metamate"},
		)

		privateKey, err := privateKeyFromString(opts.Auth.PrivateKey)
		if err != nil {
			return
		}

		svc, err := pkg.NewService(pkg.ServiceOpts{
			Client:     cli,
			PrivateKey: privateKey,
			Salt:       opts.Auth.Salt,
		})
		if err != nil {
			return
		}

		h = auth.NewHttpJsonServer(auth.HttpJsonServerOpts{Service: svc})

		t = sdk.ServiceTransport.HttpJson

		return
	}
}

func validateAuthOpts(opts types.VirtualSvcOpts) (err error) {
	if opts.Auth == nil {
		err = errors.New("opts.Auth needs to be set")

		return
	}

	if opts.Auth.PrivateKey == "" {
		err = errors.New("opts.Auth.PrivateKey needs to be set")

		return
	}

	if opts.Auth.Salt == "" {
		err = errors.New("opts.Auth.Salt needs to be set")

		return
	}

	return
}

func privateKeyFromString(s string) (k *rsa.PrivateKey, err error) {
	block, _ := pem.Decode([]byte(s))
	if block == nil {
		err = errors.New("failed to parse PEM block containing the key")

		return
	}

	k, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	return
}
