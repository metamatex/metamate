package spec

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"

	"github.com/stretchr/testify/require"
	"testing"
)

func TestPutRelationMode(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string, suffix string) {
	name := "TestPutRelationMode"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			whatevers := []mql.Whatever{
				{
					Id: &mql.ServiceId{
						Value: mql.String(nameSvcId(suffix, name, "a")),
					},
				},
				{
					Id: &mql.ServiceId{
						Value: mql.String(nameSvcId(suffix, name, "b")),
					},
				},
				{
					Id: &mql.ServiceId{
						Value: mql.String(nameSvcId(suffix, name, "c")),
					},
				},
			}

			requirePostWhatevers(t, ctx, f, h, svcName, whatevers)

			err = requirePutWhatevers(t, ctx, f, h, whatevers[0].Id, []mql.ServiceId{*whatevers[1].Id, *whatevers[2].Id}, mql.WhateverRelationName.WhateverKnowsWhatevers, mql.RelationOperation.Add)
			if err != nil {
				return
			}

			err = requirePutWhatevers(t, ctx, f, h, whatevers[0].Id, []mql.ServiceId{*whatevers[1].Id, *whatevers[2].Id}, mql.WhateverRelationName.WhateverKnowsWhatevers, mql.RelationOperation.Remove)
			if err != nil {
				return
			}

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func requirePutWhatevers(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), id *mql.ServiceId, ids []mql.ServiceId, relationName string, operation string) (err error) {
	putReq := mql.PutWhateversRequest{
		Mode: &mql.PutMode{
			Kind: &mql.PutModeKind.Relation,
			Relation: &mql.RelationPutMode{
				Id:        id,
				Ids:       ids,
				Relation:  mql.String(relationName),
				Operation: mql.String(operation),
			},
		},
		Select: &mql.PutWhateversResponseSelect{
			Meta: GetResponseMetaSelect(),
		},
	}

	gPutRsp, err := h(ctx, f.MustFromStruct(putReq))
	if err != nil {
		return
	}

	requirePutRsp(t, gPutRsp)

	return
}

func requirePutBlueWhatevers(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), id *mql.ServiceId, ids []mql.ServiceId, relationName string, operation string) (err error) {
	putReq := mql.PutBlueWhateversRequest{
		Mode: &mql.PutMode{
			Kind: &mql.PutModeKind.Relation,
			Relation: &mql.RelationPutMode{
				Id:        id,
				Ids:       ids,
				Relation:  mql.String(relationName),
				Operation: mql.String(operation),
			},
		},
		Select: &mql.PutBlueWhateversResponseSelect{
			Meta: GetResponseMetaSelect(),
		},
	}

	gPutRsp, err := h(ctx, f.MustFromStruct(putReq))
	if err != nil {
		return
	}

	requirePutRsp(t, gPutRsp)

	return
}

func requirePutRsp(t *testing.T, gRsp generic.Generic) {
	require.NotNil(t, gRsp)
	requireNoRspMetaErrs(t, gRsp)
}
