package virtual

import (
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/sdk"
	
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

func Deploy(c *Cluster, svcs []types.VirtualSvc) (err error) {
	for _, svc := range svcs {
		err = validateVirtualSvc(svc)
		if err != nil {
			return
		}
	}

	for _, svc := range svcs {
		err = c.HostSvc(svc)
		if err != nil {
			return
		}
	}

	return
}

func validateVirtualSvc(svc types.VirtualSvc) (err error) {
	if svc.Id == "" {
		err = errors.New("must set id")

		return
	}

	spew.Dump(svc)
	if svc.Name == "" {
		err = errors.New("must set name")

		return
	}

	if svc.Opts != nil {
		return validateVirtualSvcOpts(*svc.Opts)
	}

	return
}

func validateVirtualSvcOpts(opts types.VirtualSvcOpts) (err error) {
	c := 0

	if opts.Mastodon != nil {
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

func (c *Cluster) HostSvc(svc types.VirtualSvc) (err error) {
	f, t, err := handler[svc.Name](c.f, c.rn, &http.Client{Transport: c}, svc)
	if err != nil {
		return
	}

	_, ok := c.hs[svc.Id]
	if ok {
		err = errors.New(fmt.Sprintf("host %v is already taken", svc.Id))

		return
	}

	c.hs[svc.Id] = f

	c.svcs[svc.Id] = sdk.Service{
		Id: &sdk.ServiceId{
			Value: sdk.String(svc.Id),
		},
		IsVirtual: sdk.Bool(true),
		Transport: &t,
		Port:      sdk.Int32(80),
		Url: &sdk.Url{
			Value: sdk.String("http://" + svc.Id),
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
			Value: sdk.String(id),
		},
		IsVirtual: sdk.Bool(true),
		Transport: sdk.String(transport),
		Port:      sdk.Int32(80),
		Url: &sdk.Url{
			Value: sdk.String("http://" + id),
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
