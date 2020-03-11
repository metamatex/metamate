package boot_test

import (
	"context"
	"fmt"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamate/metamate/pkg/v0/boot"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/line"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/metamatex/metamate/spec/pkg/v0/spec"
	"log"
	"net/http"
	"testing"
	"time"
)

const (
	AuthSvc   = "auth"
	PipeSvc   = "pipe"
	ReqFilter = "request-filter"
	SqlxA     = "sqlx-a"
	SqlxB     = "sqlx-b"
)

func TestBoot(t *testing.T) {
	c := boot.NewBaseConfig()

	d, err := boot.NewDependencies(c, types.Version{})
	if err != nil {
		log.Fatal(err)
	}

	d.ResolveLine.(*line.Line).SetLog(false)

	ctx := context.Background()

	suffix := "-" + time.Now().Format("3:04:05PM")
	println(suffix)

	client := httpjson.NewClient(d.Factory, &http.Client{}, fmt.Sprintf("http://localhost:%v/httpjson", c.Host.HttpPort), "")

	f := func(ctx context.Context, gCliReq generic.Generic) (gCliRsp generic.Generic, err error) {
		gCliRsp, err = client.Send(gCliReq)
		if err != nil {
			return
		}

		return
	}

	go func() {
		err := http.ListenAndServe(":80", d.Router)
		if err != nil {
			panic(err)
		}
	}()

	//spec.TestRequestFilter(t, ctx, d.Factory, f)

	//spec.TestPipe(t, ctx, d.Factory, f)

	spec.TestDiscovery(t, ctx, d.Factory, f)

	//spec.TestEmptyPost(t, ctx, d.Factory, f)
	//
	//spec.TestEmptyGet(t, ctx, d.Factory, f)
	//
	//spec.TestPost(t, ctx, d.Factory, f, SqlxA)
	//
	//spec.TestPostWithNameId(t, ctx, d.Factory, f, SqlxA, suffix)
	//
	//spec.TestGetModeId(t, ctx, d.Factory, f, SqlxA, suffix)
	//
	//spec.TestGetModeCollection(t, ctx, d.Factory, f, SqlxA)
	//
	//spec.TestFilterStringIs(t, ctx, d.Factory, f, SqlxA, suffix)

	// spec.TestGetModeRelation(t, ctx, d.Factory, f, SqlxA, suffix)

	//spec.TestGetModeIdWithNameId(t, ctx, d.Factory, f, suffix, SqlxA)
	//
	//spec.TestGetModeIdWithServiceFilter(t, ctx, d.Factory, f, suffix, SqlxA, SqlxB)
	//
	//spec.TestPostClientAccounts(t, ctx, d.Factory, f, SqlxA)
	//
	//spec.TestAuthenticateClientAccount(t, ctx, d.Factory, f, SqlxA, AuthSvc)
	//
	//spec.TestToken(t, ctx, d.Factory, f, AuthSvc, SqlxA)
	//
	//spec.TestGetModeRelationInter(t, ctx, d.Factory, f, SqlxA, SqlxB, suffix)
	//
	//spec.TestGetModeIdWithSelfReferencingRelation(t, ctx, d.Factory, f, suffix, SqlxA)
	//
	//spec.TestGetModeIdWithRelation(t, ctx, d.Factory, f, suffix, SqlxA)
}
