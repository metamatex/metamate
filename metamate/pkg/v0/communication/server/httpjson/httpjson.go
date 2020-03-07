package httpjson

import (
	"context"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamatemono/metamate/pkg/v0/types"
	"net/http"
)

func GetRoutes(rn *graph.RootNode, f generic.Factory, handler func(ctx context.Context, gCliReq generic.Generic) (gCliRsp generic.Generic), logErr func(err error)) (rs []types.Route) {
	rs = append(rs, types.Route{Method: http.MethodPost, Path: "/httpjson", Handler: httpjson.NewServer(httpjson.ServerOpts{
		Root:    rn,
		Factory: f,
		Handler: handler,
		LogErr:  logErr,
	})})

	return
}
