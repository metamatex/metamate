package embedded

import (
	"context"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"io/ioutil"
	"net/http"
)

type Cluster struct {
	svcs     map[string]mql.Service
	hs       map[string]http.Handler
	f        generic.Factory
	rn       *graph.RootNode
	serveBus func()
	logErr   func(error)
}

func Deploy(c *Cluster, svcs []types.EmbeddedSvc) (err error) {
	for _, svc := range svcs {
		err = validateEmbeddedSvc(svc)
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

func validateEmbeddedSvc(svc types.EmbeddedSvc) (err error) {
	if svc.Id == "" {
		err = errors.New("must set id")

		return
	}

	if svc.Name == "" {
		err = errors.New("must set name")

		return
	}

	return
}

func NewCluster(rn *graph.RootNode, f generic.Factory, logErr func(err error)) (n *Cluster) {
	n = &Cluster{
		svcs:   map[string]mql.Service{},
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

func (c *Cluster) HostSvc(svc types.EmbeddedSvc) (err error) {
	f, err := handler[svc.Name](c.f, c.rn, &http.Client{Transport: c}, svc)
	if err != nil {
		return
	}

	_, ok := c.hs[svc.Id]
	if ok {
		err = errors.New(fmt.Sprintf("host %v is already taken", svc.Id))

		return
	}

	c.hs[svc.Id] = f

	c.svcs[svc.Id] = mql.Service{
		Id: &mql.ServiceId{
			Value: mql.String(svc.Id),
		},
		IsEmbedded: mql.Bool(true),
		Port:       mql.Int32(80),
		Url: &mql.Url{
			Value: mql.String("http://" + svc.Id),
		},
	}

	return
}

func (c *Cluster) Host(id string, h http.Handler) (err error) {
	_, ok := c.hs[id]
	if ok {
		err = errors.New(fmt.Sprintf("host %v is already taken", id))

		return
	}

	c.hs[id] = h

	c.svcs[id] = mql.Service{
		Id: &mql.ServiceId{
			Value: mql.String(id),
		},
		IsEmbedded: mql.Bool(true),
		Port:       mql.Int32(80),
		Url: &mql.Url{
			Value: mql.String("http://" + id),
		},
	}

	return
}

func (c *Cluster) HostHttpJsonFunc(id string, f func(context.Context, generic.Generic) generic.Generic) (err error) {
	h := generic.NewServer(generic.ServerOpts{
		Root:    c.rn,
		Factory: c.f,
		Handler: f,
		LogErr:  c.logErr,
	})

	return c.Host(id, h)
}

func (c *Cluster) HostBus(id string, f func(context.Context, generic.Generic) generic.Generic) (err error) {
	_, ok := c.hs[id]
	if ok {
		err = errors.New(fmt.Sprintf("host %v is already taken", id))

		return
	}

	c.hs[id] = generic.NewServer(generic.ServerOpts{
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
	rsp.StatusCode = 200
	rsp.Body = ioutil.NopCloser(&w.b)

	return
}

func (c *Cluster) serveDiscovery(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic) {
	switch gReq.Type().Name() {
	case mql.LookupServiceBusRequestName:
		return c.f.MustFromStruct(mql.LookupServiceServiceResponse{
			Output: &mql.LookupServiceOutput{
				Service: &mql.Service{
					Endpoints: &mql.ServiceEndpoints{
						LookupService: &mql.LookupServiceEndpoint{},
						GetServices: &mql.GetServicesEndpoint{
							//Filter: &mql.GetServicesRequestFilter{
							//	Mode: &mql.GetModeFilter{
							//		Kind: &mql.EnumFilter{
							//			In: []string{mql.GetModeKind.Collection, mql.GetModeKind.Id},
							//		},
							//	},
							//},
						},
					},
				},
			},
		})
	case mql.TypeNames.GetServicesBusRequest:
		var errs []mql.Error
		var svcs []mql.Service

		var req mql.GetServicesBusRequest
		gReq.MustToStruct(&req)

		switch *req.Mode.Kind {
		case mql.GetModeKind.Id:

			svc, ok := c.svcs[*req.Mode.Id.ServiceId.Value]
			if !ok {
				errs = append(errs, mql.Error{
					Kind: &mql.ErrorKind.IdNotPresent,
					Id:   req.Mode.Id,
				})
			} else {
				svcs = append(svcs, svc)
			}
		case mql.GetModeKind.Collection:
			for _, svc := range c.svcs {
				svcs = append(svcs, svc)
			}
		}

		return c.f.MustFromStruct(mql.GetServicesServiceResponse{
			Errors:   errs,
			Services: svcs,
		})
	default:
		panic("can't handle " + gReq.Type().Name())
	}

	return
}
