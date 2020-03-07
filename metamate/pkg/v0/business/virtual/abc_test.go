package virtual_test

import (
	"context"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/transport"
	"github.com/metamatex/metamatemono/metamate/pkg/v0/boot"
	"github.com/metamatex/metamatemono/metamate/pkg/v0/business/virtual"
	"github.com/prometheus/common/log"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	err := func() (err error) {
		c := boot.NewBaseConfig()

		d, err := boot.NewDependencies(c)
		if err != nil {
			return
		}

		n := virtual.NewCluster(d.RootNode, d.Factory, nil)

		err = n.HostName("sqlx", virtual.Sqlx)
		if err != nil {
			return
		}

		err = n.HostName("pipe", virtual.Pipe)
		if err != nil {
			return
		}

		err = n.HostName("auth", virtual.Auth)
		if err != nil {
			return
		}

		err = n.HostName("reqFilter", virtual.ReqFilter)
		if err != nil {
			return
		}

		cli := transport.NewHttpJsonClient(transport.HttpJsonClientOpts{
			HttpClient: &http.Client{
				Transport: n,
			},
			Addr: "http://discovery",
		})

		ctx := context.Background()

		rsp, err := cli.GetServices(ctx, sdk.GetServicesRequest{})
		if err != nil {
			return
		}

		d.Factory.MustFromStruct(rsp).Print()

		return
	}()
	if err != nil {
		log.Fatal(err)
	}
}
