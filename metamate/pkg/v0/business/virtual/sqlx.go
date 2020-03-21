// +build !lite

package virtual

import (
	"errors"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/metamatex/metamate/sqlx-svc/pkg/boot"
	sqlxTypes "github.com/metamatex/metamate/sqlx-svc/pkg/types"
	"net/http"
)

func init() {
	handler[Sqlx] = func(f generic.Factory, rn *graph.RootNode, cli *http.Client, vSvc types.VirtualSvc) (h http.Handler, t string, err error) {
		err = validateSqlxOpts(*vSvc.Opts)
		if err != nil {
			return
		}

		c := sqlxTypes.Config{
			Log:        vSvc.Opts.Sqlx.Log,
			DriverName: vSvc.Opts.Sqlx.Driver,
			DataSource: vSvc.Opts.Sqlx.Connection,
			TypeNames:  vSvc.Opts.Sqlx.Types,
		}

		d, err := boot.NewDependencies(rn, f, c)
		if err != nil {
			return
		}

		h = httpjson.NewServer(httpjson.ServerOpts{
			Root:    rn,
			Factory: f,
			LogErr:  nil,
			Handler: d.ServeFunc,
		})

		t = sdk.ServiceTransport.HttpJson

		return
	}
}

func validateSqlxOpts(opts types.VirtualSvcOpts) (err error) {
	if opts.Sqlx == nil {
		err = errors.New("opts.Sqlx needs to be set")

		return
	}

	if opts.Sqlx.Driver == "" {
		err = errors.New("opts.Sqlx.Driver needs to be set")

		return
	}

	if opts.Sqlx.Connection == "" {
		err = errors.New("opts.Sqlx.Connection needs to be set")

		return
	}

	if len(opts.Sqlx.Types) == 0 {
		err = errors.New("opts.Sqlx.Types needs to be set")

		return
	}

	return
}
