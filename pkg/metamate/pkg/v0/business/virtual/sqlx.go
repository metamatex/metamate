// +build !lite

package virtual

import (
	"errors"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/pkg/metamate/pkg/v0/types"
	"github.com/metamatex/metamatemono/pkg/services/sqlx-svc/pkg/boot"
	sqlxTypes "github.com/metamatex/metamatemono/pkg/services/sqlx-svc/pkg/types"
	"net/http"
)

func init() {
	handler[Sqlx] = func(f generic.Factory, rn *graph.RootNode, cli *http.Client, opts types.VirtualSvcOpts) (h http.Handler, t string, err error) {
		err = validateSqlxOpts(opts)
		if err != nil {
			return
		}

		c := sqlxTypes.Config{
			Log:        opts.Sqlx.Log,
			DriverName: opts.Sqlx.Driver,
			DataSource: opts.Sqlx.Connection,
			TypeNames:  opts.Sqlx.Types,
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
