package spec

import (
	"context"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/gen/v0/sdk"
	
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPutRelationMode(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string, suffix string) {
	name := "TestPutRelationMode"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			whatevers := []sdk.Whatever{
				{
					Id: &sdk.ServiceId{
						Value: sdk.String(nameSvcId(suffix, name, "a")),
					},
				},
				{
					Id: &sdk.ServiceId{
						Value: sdk.String(nameSvcId(suffix, name, "b")),
					},
				},
				{
					Id: &sdk.ServiceId{
						Value: sdk.String(nameSvcId(suffix, name, "c")),
					},
				},
			}

			requirePostWhatevers(t, ctx, f, h, svcName, whatevers)

			err = requirePutWhatevers(t, ctx, f, h, whatevers[0].Id, []sdk.ServiceId{*whatevers[1].Id, *whatevers[2].Id}, sdk.WhateverRelationName.WhateverKnowsWhatevers, sdk.RelationOperation.Add)
			if err != nil {
				return
			}

			err = requirePutWhatevers(t, ctx, f, h, whatevers[0].Id, []sdk.ServiceId{*whatevers[1].Id, *whatevers[2].Id}, sdk.WhateverRelationName.WhateverKnowsWhatevers, sdk.RelationOperation.Remove)
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

func requirePutWhatevers(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), id *sdk.ServiceId, ids []sdk.ServiceId, relationName string, operation string) (err error) {
	putReq := sdk.PutWhateversRequest{
		Mode: &sdk.PutMode{
			Kind: &sdk.PutModeKind.Relation,
			Relation: &sdk.RelationPutMode{
				Id: id,
				Ids: ids,
				Relation:  sdk.String(relationName),
				Operation: sdk.String(operation),
			},
		},
		Select: &sdk.PutWhateversResponseSelect{
			Meta:GetResponseMetaSelect(),
		},
	}

	gPutRsp, err := h(ctx, f.MustFromStruct(putReq))
	if err != nil {
	    return
	}

	requirePutRsp(t, gPutRsp)

	return
}

func requirePutBlueWhatevers(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), id *sdk.ServiceId, ids []sdk.ServiceId, relationName string, operation string) (err error) {
	putReq := sdk.PutBlueWhateversRequest{
		Mode: &sdk.PutMode{
			Kind: &sdk.PutModeKind.Relation,
			Relation: &sdk.RelationPutMode{
				Id: id,
				Ids: ids,
				Relation:  sdk.String(relationName),
				Operation: sdk.String(operation),
			},
		},
		Select: &sdk.PutBlueWhateversResponseSelect{
			Meta:GetResponseMetaSelect(),
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