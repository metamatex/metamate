package spec

import (
	"context"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/utils/ptr"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetModeRelation(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName, suffix string) {
	name := "TestGetModeRelation"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func () (err error) {
			whatevers := []sdk.Whatever{
				{
					Id: &sdk.ServiceId{
						Value: ptr.String(nameSvcId(suffix, name, "a")),
					},
				},
				{
					Id: &sdk.ServiceId{
						Value: ptr.String(nameSvcId(suffix, name, "b")),
					},
				},
				{
					Id: &sdk.ServiceId{
						Value: ptr.String(nameSvcId(suffix, name, "c")),
					},
				},
			}

			postRsp := requirePostWhatevers(t, ctx, f, h, svcName, whatevers)

			err = requirePutWhatevers(t, ctx, f, h, postRsp.Whatevers[0].Id, []sdk.ServiceId{*postRsp.Whatevers[1].Id, *postRsp.Whatevers[2].Id}, sdk.WhateverRelationName.WhateverKnowsWhatevers, sdk.RelationOperation.Add)
			if err != nil {
			    return
			}

			getReq := sdk.GetWhateversRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Relation,
					Relation: &sdk.RelationGetMode{
						Relation: ptr.String(sdk.WhateverRelationName.WhateverKnowsWhatevers),
						Id: postRsp.Whatevers[0].Id,
					},
				},
				Select: &sdk.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							ServiceName: ptr.Bool(true),
							Value: ptr.Bool(true),
						},
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			requireSvcIdValues(t, gGetRsp.MustGenericSlice(fieldnames.Whatevers), []string{*whatevers[1].Id.Value, *whatevers[2].Id.Value})

			getReq = sdk.GetWhateversRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Relation,
					Relation: &sdk.RelationGetMode{
						Relation: &sdk.WhateverRelationName.WhateverKnewByWhatevers,
						Id: postRsp.Whatevers[1].Id,
					},
				},
				Select: &sdk.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							ServiceName: ptr.Bool(true),
							Value: ptr.Bool(true),
						},
					},
				},
			}

			gGetRsp, err = h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			getRsp := sdk.GetWhateversResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.Equal(t, whatevers[0].Id.Value, getRsp.Whatevers[0].Id.Value)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestGetModeRelationInter(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), storageSvcAName, storageSvcBName, suffix string) {
	name := "TestGetModeRelationInter"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			whatevers := []sdk.Whatever{
				{
					Id: &sdk.ServiceId{
						Value: ptr.String(nameSvcId(suffix, name, "a")),
					},
				},
				{
					Id: &sdk.ServiceId{
						Value: ptr.String(nameSvcId(suffix, name, "b")),
					},
				},
				{
					Id: &sdk.ServiceId{
						Value: ptr.String(nameSvcId(suffix, name, "c")),
					},
				},
			}

			postRspA := requirePostWhatevers(t, ctx, f, h, storageSvcAName, whatevers[:1])

			postRspB := requirePostWhatevers(t, ctx, f, h, storageSvcBName, whatevers[1:])

			err = requirePutWhatevers(t, ctx, f, h, postRspA.Whatevers[0].Id, []sdk.ServiceId{*postRspB.Whatevers[0].Id, *postRspB.Whatevers[1].Id}, sdk.WhateverRelationName.WhateverKnowsWhatevers, sdk.RelationOperation.Add)
			if err != nil {
			    return
			}

			getReq := sdk.GetWhateversRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Relation,
					Relation: &sdk.RelationGetMode{
						Relation: ptr.String(sdk.WhateverRelationName.WhateverKnowsWhatevers),
						Id: postRspA.Whatevers[0].Id,
					},
				},
				Select: &sdk.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							ServiceName: ptr.Bool(true),
							Value: ptr.Bool(true),
						},
					},
				},
				Relations: &sdk.GetWhateversRelations{
					KnowsWhatevers: &sdk.GetWhateversCollection{},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
			    return
			}

			requireGetRsp(t, gGetRsp)

			requireSvcIdValues(t, gGetRsp.MustGenericSlice(fieldnames.Whatevers), []string{*whatevers[1].Id.Value, *whatevers[2].Id.Value})

			getReq = sdk.GetWhateversRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Relation,
					Relation: &sdk.RelationGetMode{
						Relation: &sdk.WhateverRelationName.WhateverKnewByWhatevers,
						Id: postRspB.Whatevers[0].Id,
					},
				},
				Select: &sdk.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							ServiceName: ptr.Bool(true),
							Value: ptr.Bool(true),
						},
					},
				},
			}

			gGetRsp, err = h(ctx, f.MustFromStruct(getReq))
			if err != nil {
			    return
			}

			requireGetRsp(t, gGetRsp)

			getRsp := sdk.GetWhateversResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.Equal(t, whatevers[0].Id.Value, getRsp.Whatevers[0].Id.Value)

			return
		}()
		if err != nil {
		    t.Error(err)
		}
	})
}
