package types

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"net/http"
)

type Dependencies struct {
	RootNode           *graph.RootNode
	Factory            generic.Factory
	ReqHandler         map[string]RequestHandler
	LinkStore          LinkStore
	ResolveLine        Transformer
	SvcReqLog          func(ctx ReqCtx)
	ClientTransportLog func(ctx ReqCtx)
	ServeFunc          ServeFunc
	Router             http.Handler
	Routes             []Route
	InternalLogTemplates InternalLogTemplates
}
