package httpjson

import (
	"context"
	"encoding/json"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"net/http"

	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
)

type Server struct {
	opts   ServerOpts
}

type ServerOpts struct {
	Root    *graph.RootNode
	Factory generic.Factory
	Handler func(ctx context.Context, gRequest generic.Generic) (gResponse generic.Generic)
	LogErr  func(err error)
}

func NewServer(c ServerOpts) (Server) {
	return Server{
		opts: c,
	}
}

func (s Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := func() (err error) {
		m := map[string]interface{}{}
		err = json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			return
		}

		tn, err := s.opts.Root.Types.ByName(req.Header.Get(METAMATE_TYPE_HEADER))
		if err != nil {
			return
		}

		gReq := s.opts.Factory.MustFromStringInterfaceMap(tn, m)

		gRsp := s.opts.Handler(req.Context(), gReq)

		w.Header().Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_JSON)
		w.Header().Set(METAMATE_TYPE_HEADER, gRsp.Type().Name())

		err = json.NewEncoder(w).Encode(gRsp.ToStringInterfaceMap())
		if err != nil {
			return
		}

		return
	}()
	if err != nil {
		if s.opts.LogErr != nil {
			s.opts.LogErr(err)
		}

		w.WriteHeader(200)
	}
}
