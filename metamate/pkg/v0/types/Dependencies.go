package types

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"net/http"
)

type Dependencies struct {
	RootNode             *graph.RootNode
	Factory              generic.Factory
	ResolveLine          Transformer
	ServeFunc            ServeFunc
	Router               http.Handler
	Routes               []Route
	InternalLogTemplates InternalLogTemplates
	Run                  []func() (err error)
}
