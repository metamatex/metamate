package generic

import (
	"context"
	"encoding/json"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"io/ioutil"
	"net/http"
)

type Server struct {
	opts ServerOpts
}

type ServerOpts struct {
	Name    string
	Log     func(name string, b []byte, req *http.Request)
	Root    *graph.RootNode
	Factory Factory
	Handler func(ctx context.Context, gRequest Generic) (gResponse Generic)
	LogErr  func(err error)
}

func NewServer(c ServerOpts) Server {
	return Server{
		opts: c,
	}
}

func (s Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := func() (err error) {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return
		}
		if s.opts.Log != nil {
			s.opts.Log(s.opts.Name, b, req)
		}

		m := map[string]interface{}{}
		err = json.Unmarshal(b, &m)
		if err != nil {
			return
		}

		tn, err := s.opts.Root.Types.ByName(req.Header.Get(AsgTypeHeader))
		if err != nil {
			return
		}

		gReq := s.opts.Factory.MustFromStringInterfaceMap(tn, m)

		gRsp := s.opts.Handler(req.Context(), gReq)

		w.Header().Set(ContentTypeHeader, ContentTypeJson)
		w.Header().Set(AsgTypeHeader, gRsp.Type().Name())

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
	}
}
