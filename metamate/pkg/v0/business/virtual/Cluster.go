package virtual

import (
	"context"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/utils/ptr"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"io/ioutil"
	"net/http"
)

type Cluster struct {
	svcs     map[string]sdk.Service
	hs       map[string]http.Handler
	f        generic.Factory
	rn       *graph.RootNode
	serveBus func()
	logErr   func(error)
}

func Deploy(c *Cluster, os []types.VirtualSvcOpts) (err error) {
	for _, opts := range os {
		err = validateVirtualSvcOpts(opts)
		if err != nil {
			return
		}
	}

	for _, opts := range os {
		if opts.Auth != nil {
			err = c.HostName(Auth, opts)
			if err != nil {
				return
			}
		}

		if opts.Mastodon != nil {
			err = c.HostName(Mastodon, opts)
			if err != nil {
				return
			}
		}

		if opts.Sqlx != nil {
			err = c.HostName(Sqlx, opts)
			if err != nil {
				return
			}
		}
	}

	return
}

func validateVirtualSvcOpts(opts types.VirtualSvcOpts) (err error) {
	if opts.Name == "" {
		err = errors.New("must set name")

		return
	}

	c := 0

	if opts.Auth != nil {
		c++
	}

	if opts.Mastodon != nil {
		c++
	}

	if opts.Sqlx != nil {
		c++
	}

	if c != 1 {
		err = errors.New("must exactly one opts")

		return
	}

	return
}

func NewCluster(rn *graph.RootNode, f generic.Factory, logErr func(err error)) (n *Cluster) {
	n = &Cluster{
		svcs:   map[string]sdk.Service{},
		hs:     map[string]http.Handler{},
		f:      f,
		rn:     rn,
		logErr: logErr,
	}

	err := n.HostHttpJsonFunc("discovery", n.serveDiscovery)
	if err != nil {
		return
	}

	return
}

func (c *Cluster) HostName(id string, opts types.VirtualSvcOpts) (err error) {
	f, t, err := handler[id](c.f, c.rn, &http.Client{Transport: c}, opts)
	if err != nil {
		return
	}

	_, ok := c.hs[opts.Name]
	if ok {
		err = errors.New(fmt.Sprintf("host %v is already taken", opts.Name))

		return
	}

	c.hs[opts.Name] = f

	c.svcs[opts.Name] = sdk.Service{
		Id: &sdk.ServiceId{
			Value: ptr.String(opts.Name),
		},
		IsVirtual: ptr.Bool(true),
		Transport: &t,
		Port:      ptr.Int32(80),
		Url: &sdk.Url{
			Value: ptr.String("http://" + opts.Name),
		},
	}

	return
}

func (c *Cluster) Host(id, transport string, h http.Handler) (err error) {
	_, ok := c.hs[id]
	if ok {
		err = errors.New(fmt.Sprintf("host %v is already taken", id))

		return
	}

	c.hs[id] = h

	c.svcs[id] = sdk.Service{
		Id: &sdk.ServiceId{
			Value: ptr.String(id),
		},
		IsVirtual: ptr.Bool(true),
		Transport: ptr.String(transport),
		Port:      ptr.Int32(80),
		Url: &sdk.Url{
			Value: ptr.String("http://" + id),
		},
	}

	return
}

func (c *Cluster) HostHttpJsonFunc(id string, f func(context.Context, generic.Generic) generic.Generic) (err error) {
	h := httpjson.NewServer(httpjson.ServerOpts{
		Root:    c.rn,
		Factory: c.f,
		Handler: f,
		LogErr:  c.logErr,
	})

	return c.Host(id, sdk.ServiceTransport.HttpJson, h)
}

func (c *Cluster) HostBus(id string, f func(context.Context, generic.Generic) generic.Generic) (err error) {
	_, ok := c.hs[id]
	if ok {
		err = errors.New(fmt.Sprintf("host %v is already taken", id))

		return
	}

	c.hs[id] = httpjson.NewServer(httpjson.ServerOpts{
		Root:    c.rn,
		Factory: c.f,
		Handler: f,
		LogErr:  c.logErr,
	})

	return
}

func (c *Cluster) RoundTrip(req *http.Request) (rsp *http.Response, err error) {
	h, ok := c.hs[req.Host]
	if !ok {
		err = errors.New(fmt.Sprintf("can't resolve host %v", req.Host))

		return
	}

	w := newResponseWriter()

	h.ServeHTTP(w, req)

	rsp = &http.Response{}
	rsp.Header = w.header
	rsp.Body = ioutil.NopCloser(&w.b)

	return
}

func (c *Cluster) serveDiscovery(ctx context.Context, gRequest generic.Generic) (gResponse generic.Generic) {
	switch gRequest.Type().Name() {
	case sdk.LookupServiceRequestName:
		return c.f.MustFromStruct(sdk.LookupServiceResponse{
			Output: &sdk.LookupServiceOutput{
				Service: &sdk.Service{
					Endpoints: &sdk.Endpoints{
						LookupService: &sdk.LookupServiceEndpoint{},
						GetServices:   &sdk.GetServicesEndpoint{},
					},
				},
			},
		})
	case sdk.GetServicesRequestName:
		var svcs []sdk.Service

		for _, svc := range c.svcs {
			svcs = append(svcs, svc)
		}

		return c.f.MustFromStruct(sdk.GetServicesResponse{
			Services: svcs,
		})
	}

	return
}
