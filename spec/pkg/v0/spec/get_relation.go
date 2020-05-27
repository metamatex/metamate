package spec

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"

	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetModeRelation(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName, suffix string) {
	name := "TestGetModeRelation"
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

			postRsp := requirePostWhatevers(t, ctx, f, h, svcName, whatevers)

			err = requirePutWhatevers(t, ctx, f, h, postRsp.Whatevers[0].Id, []mql.ServiceId{*postRsp.Whatevers[1].Id, *postRsp.Whatevers[2].Id}, mql.WhateverRelationName.WhateverKnowsWhatevers, mql.RelationOperation.Add)
			if err != nil {
				return
			}

			getReq := mql.GetWhateversRequest{
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Relation,
					Relation: &mql.RelationGetMode{
						Relation: mql.String(mql.WhateverRelationName.WhateverKnowsWhatevers),
						Id:       postRsp.Whatevers[0].Id,
					},
				},
				Select: &mql.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							ServiceName: mql.Bool(true),
							Value:       mql.Bool(true),
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

			getReq = mql.GetWhateversRequest{
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Relation,
					Relation: &mql.RelationGetMode{
						Relation: &mql.WhateverRelationName.WhateverKnewByWhatevers,
						Id:       postRsp.Whatevers[1].Id,
					},
				},
				Select: &mql.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							ServiceName: mql.Bool(true),
							Value:       mql.Bool(true),
						},
					},
				},
			}

			gGetRsp, err = h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			getRsp := mql.GetWhateversResponse{}
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

			postRspA := requirePostWhatevers(t, ctx, f, h, storageSvcAName, whatevers[:1])

			postRspB := requirePostWhatevers(t, ctx, f, h, storageSvcBName, whatevers[1:])

			err = requirePutWhatevers(t, ctx, f, h, postRspA.Whatevers[0].Id, []mql.ServiceId{*postRspB.Whatevers[0].Id, *postRspB.Whatevers[1].Id}, mql.WhateverRelationName.WhateverKnowsWhatevers, mql.RelationOperation.Add)
			if err != nil {
				return
			}

			getReq := mql.GetWhateversRequest{
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Relation,
					Relation: &mql.RelationGetMode{
						Relation: mql.String(mql.WhateverRelationName.WhateverKnowsWhatevers),
						Id:       postRspA.Whatevers[0].Id,
					},
				},
				Select: &mql.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							ServiceName: mql.Bool(true),
							Value:       mql.Bool(true),
						},
					},
				},
				Relations: &mql.GetWhateversRelations{
					KnowsWhatevers: &mql.GetWhateversCollection{},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			requireSvcIdValues(t, gGetRsp.MustGenericSlice(fieldnames.Whatevers), []string{*whatevers[1].Id.Value, *whatevers[2].Id.Value})

			getReq = mql.GetWhateversRequest{
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Relation,
					Relation: &mql.RelationGetMode{
						Relation: &mql.WhateverRelationName.WhateverKnewByWhatevers,
						Id:       postRspB.Whatevers[0].Id,
					},
				},
				Select: &mql.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							ServiceName: mql.Bool(true),
							Value:       mql.Bool(true),
						},
					},
				},
			}

			gGetRsp, err = h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			getRsp := mql.GetWhateversResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.Equal(t, whatevers[0].Id.Value, getRsp.Whatevers[0].Id.Value)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}
