package virtual

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/kubernetes-svc/pkg"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[Kubernetes] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, vSvc types.VirtualSvc) (h http.Handler, t string, err error) {
		svc, err := pkg.NewService()
		if err != nil {
			return
		}

		h = mql.NewKubernetesHttpJsonServer(mql.KubernetesHttpJsonServerOpts{Service: svc})

		t = mql.ServiceTransport.HttpJson

		return
	}
}