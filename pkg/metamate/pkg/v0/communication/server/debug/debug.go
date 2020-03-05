package debug

import (
	"github.com/metamatex/metamatemono/pkg/metamate/pkg/v0/types"
	"net/http"
	"net/http/pprof"
)

func GetRoutes() (rs []types.Route) {
	rs = append(rs, types.Route{Method: http.MethodGet, Path: "/debug/pprof/*", HandlerFunc: pprof.Index})
	rs = append(rs, types.Route{Method: http.MethodGet, Path: "/debug/pprof/cmdline", HandlerFunc: pprof.Cmdline})
	rs = append(rs, types.Route{Method: http.MethodGet, Path: "/debug/pprof/profile", HandlerFunc: pprof.Profile})
	rs = append(rs, types.Route{Method: http.MethodGet, Path: "/debug/pprof/symbol", HandlerFunc: pprof.Symbol} )
	rs = append(rs, types.Route{Method: http.MethodGet, Path: "/debug/pprof/trace", HandlerFunc: pprof.Trace})

	return
}
