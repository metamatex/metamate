package main

import (
	"context"
	"github.com/metamatex/asg/pkg/v0/asg"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/spec/pkg/v0/spec"
	"github.com/metamatex/metamatemono/pkg/services/sqlx-svc/pkg/boot"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	rn, err := asg.New()
	if err != nil {
		return
	}

	f := generic.NewFactory(rn)

	ctx := context.Background()

	c := boot.NewTestConfig()

	c.Log = true

	d, err := boot.NewDependencies(rn, f, c)
	if err != nil {
		return
	}

	suffix := "-" + time.Now().Format("3:04:05PM")

	//spec.TestPost(t, ctx, d.Factory, d.ServeFunc, "")
	//
	//spec.TestPostWithNameId(t, ctx, d.Factory, d.ServeFunc, "", suffix)
	//
	//spec.TestGetModeId(t, ctx, d.Factory, d.ServeFunc, "", suffix)
	//
	//spec.TestGetModeIdWithNameId(t, ctx, d.Factory, d.ServeFunc, suffix, "")
	//
	//spec.TestGetModeCollection(t, ctx, d.Factory, d.ServeFunc, "")
	//
	//spec.TestFilterStringIs(t, ctx, d.Factory, d.ServeFunc, "", suffix)
	//
	////spec.TestGetModeIdWithZeroId(t, ctx, d.Factory, d.ServeFunc, suffix)
	//
	//spec.TestPostWithNameId(t, ctx, d.Factory, d.ServeFunc, "", suffix)

	//spec.TestPutRelationMode(t, ctx, d.Factory, d.ServeFunc, "", suffix)

	spec.TestGetModeRelation(t, ctx, d.Factory, d.ServeFunc, "", suffix)
}
