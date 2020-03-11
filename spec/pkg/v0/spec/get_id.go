package spec

import (
	"context"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/utils/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetModeId(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName, suffix string) {
	name := "TestGetModeId"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			postReq := sdk.PostWhateversRequest{
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Is: ptr.String(svcName),
						},
					},
				},
				Select: &sdk.PostWhateversResponseSelect{
					Meta: GetResponseMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							Value: ptr.Bool(true),
						},
						StringField: ptr.Bool(true),
					},
				},
				Whatevers: []sdk.Whatever{
					{
						Id: &sdk.ServiceId{
							Value: ptr.String(nameSvcId(suffix, name, "0")),
						},
						StringField: ptr.String("hi"),
					},
				},
			}

			gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
			if err != nil {
				return
			}

			postRsp := sdk.PostWhateversResponse{}
			gPostRsp.MustToStruct(&postRsp)

			requirePostRsp(t, gPostRsp)

			getReq := sdk.GetWhateversRequest{
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Is: ptr.String(svcName),
						},
					},
				},
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Id,
					Id: &sdk.Id{
						Kind: &sdk.IdKind.ServiceId,
						ServiceId: &sdk.ServiceId{
							Value: postRsp.Whatevers[0].Id.Value,
						},
					},
				},
				Select: &sdk.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							Value: ptr.Bool(true),
						},
						StringField: ptr.Bool(true),
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			getRsp := sdk.GetWhateversResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.Equal(t, *postRsp.Whatevers[0].Id.Value, *getRsp.Whatevers[0].Id.Value)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestGetModeIdWithZeroId(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), suffix string) {
	name := "TestGetModeIdWithZeroId"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetWhateversRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Id,
					Id: &sdk.Id{
						Kind: &sdk.IdKind.ServiceId,
						ServiceId: &sdk.ServiceId{
							Value: ptr.String(""),
						},
					},
				},
				Select: &sdk.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							Value: ptr.Bool(true),
						},
						StringField: ptr.Bool(true),
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			getRsp := sdk.GetWhateversResponse{}
			gGetRsp.MustToStruct(&getRsp)

			return
		}()
		if err != nil {
			return
		}
	})
}

func TestGetModeIdWithServiceFilter(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), suffix, svcNameA, svcNameB string) {
	name := "TestGetModeIdWithServiceFilter"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			postReq := sdk.PostWhateversRequest{
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Is: ptr.String(svcNameA),
						},
					},
				},
				Select: &sdk.PostWhateversResponseSelect{
					Meta: GetResponseMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							Value: ptr.Bool(true),
						},
					},
				},
				Whatevers: []sdk.Whatever{
					{
						Id: &sdk.ServiceId{
							Value: ptr.String(nameSvcId(suffix, name, "0")),
						},
					},
				},
			}

			gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
			if err != nil {
				return
			}

			postRsp := sdk.PostWhateversResponse{}
			gPostRsp.MustToStruct(&postRsp)

			requirePostRsp(t, gPostRsp)

			err = requireSvcHasSvcId(t, h, f, ctx, svcNameA, *postRsp.Whatevers[0].Id.Value)
			if err != nil {
			    return
			}

			err = requireSvcHasNotSvcId(t, h, f, ctx, svcNameB, *postRsp.Whatevers[0].Id.Value)
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

func TestGetModeIdWithNameId(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), suffix, svcName string) {
	name := "TestGetModeIdWithNameId"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			postReq := sdk.PostWhateversRequest{
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Is: ptr.String(svcName),
						},
					},
				},
				Select: &sdk.PostWhateversResponseSelect{
					Meta: GetResponseMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							Value: ptr.Bool(true),
						},
						AlternativeIds: &sdk.IdSelect{
							Kind: ptr.Bool(true),
							Name: ptr.Bool(true),
						},
						StringField: ptr.Bool(true),
					},
				},
				Whatevers: []sdk.Whatever{
					{
						AlternativeIds: []sdk.Id{
							{
								Kind: &sdk.IdKind.Name,
								Name: ptr.String(nameSvcId(suffix, name, "0")),
							},
						},
						StringField: ptr.String("a"),
					},
				},
			}

			gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
			if err != nil {
				return
			}

			postRsp := sdk.PostWhateversResponse{}
			gPostRsp.MustToStruct(&postRsp)

			requirePostRsp(t, gPostRsp)

			getReq := sdk.GetWhateversRequest{
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Is: ptr.String(svcName),
						},
					},
				},
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Id,
					Id: &sdk.Id{
						Kind: &sdk.IdKind.Name,
						Name: postRsp.Whatevers[0].AlternativeIds[0].Name,
					},
				},
				Select: &sdk.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							Value: ptr.Bool(true),
						},
						AlternativeIds: &sdk.IdSelect{
							Kind: ptr.Bool(true),
							Name: ptr.Bool(true),
						},
						StringField: ptr.Bool(true),
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			getRsp := sdk.GetWhateversResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.Equal(t, *postReq.Whatevers[0].StringField, *getRsp.Whatevers[0].StringField)

			require.NotEqual(t, 0, len(postReq.Whatevers[0].AlternativeIds))
			require.NotEqual(t, 0, len(getRsp.Whatevers[0].AlternativeIds))

			assert.Equal(t, *postReq.Whatevers[0].AlternativeIds[0].Name, *getRsp.Whatevers[0].AlternativeIds[0].Name)
			assert.NotEqual(t, "", *getRsp.Whatevers[0].Id.Value)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestGetModeIdWithSelfReferencingRelation(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), suffix, svcName string) {
	name := "TestGetModeIdWithSelfReferencingRelation"
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

			postRsp := requirePostWhatevers(t, ctx, f, h, svcName, whatevers)

			err = requirePutWhatevers(t, ctx, f, h, postRsp.Whatevers[0].Id, []sdk.ServiceId{*postRsp.Whatevers[1].Id, *postRsp.Whatevers[2].Id}, sdk.WhateverRelationName.WhateverKnowsWhatevers, sdk.RelationOperation.Add)
			if err != nil {
				return
			}

			getReq := sdk.GetWhateversRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Id,
					Id: &sdk.Id{
						Kind:      &sdk.IdKind.ServiceId,
						ServiceId: whatevers[0].Id,
					},
				},
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Is: ptr.String(svcName),
						},
					},
				},
				Select: &sdk.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							Value: ptr.Bool(true),
						},
						AlternativeIds: &sdk.IdSelect{
							Kind: ptr.Bool(true),
							Name: ptr.Bool(true),
							Email: &sdk.EmailSelect{
								Value: ptr.Bool(true),
							},
						},
						Relations: &sdk.WhateverRelationsSelect{
							KnowsWhatevers: &sdk.WhateversCollectionSelect{
								Meta: GetCollectionMetaSelect(),
								Whatevers: &sdk.WhateverSelect{
									Id: &sdk.ServiceIdSelect{
										Value: ptr.Bool(true),
									},
								},
							},
						},
					},
				},
				Relations: &sdk.GetWhateversRelations{
					KnowsWhatevers: &sdk.GetWhateversCollection{
						Relations: &sdk.GetWhateversRelations{
							KnewByWhatevers: &sdk.GetWhateversCollection{
							},
						},
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			requireSvcIdValues(t, gGetRsp.MustGenericSlice(fieldnames.Whatevers).Get()[0].MustGenericSlice(fieldnames.Relations, "knowsWhatevers", fieldnames.Whatevers), []string{*whatevers[1].Id.Value, *whatevers[2].Id.Value})

			requireSvcIdValues(t, gGetRsp.MustGenericSlice(fieldnames.Whatevers).Get()[0].MustGenericSlice(fieldnames.Relations, "knowsWhatevers", fieldnames.Whatevers).Get()[0].MustGenericSlice(fieldnames.Relations, "knewByWhatevers", fieldnames.Whatevers), []string{*whatevers[0].Id.Value})
			requireSvcIdValues(t, gGetRsp.MustGenericSlice(fieldnames.Whatevers).Get()[0].MustGenericSlice(fieldnames.Relations, "knowsWhatevers", fieldnames.Whatevers).Get()[1].MustGenericSlice(fieldnames.Relations, "knewByWhatevers", fieldnames.Whatevers), []string{*whatevers[0].Id.Value})

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestGetModeIdWithRelation(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), suffix, svcName string) {
	name := "TestGetModeIdWithRelation"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			whatevers := []sdk.Whatever{
				{
					Id: &sdk.ServiceId{
						Value: ptr.String(nameSvcId(suffix, name, "a")),
					},
				},
			}

			blueWhatevers := []sdk.BlueWhatever{
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
			}

			postWhateversRsp := requirePostWhatevers(t, ctx, f, h, svcName, whatevers)

			postBlueWhateversRsp := requirePostBlueWhatevers(t, ctx, f, h, svcName, blueWhatevers)

			err = requirePutBlueWhatevers(t, ctx, f, h, postWhateversRsp.Whatevers[0].Id, []sdk.ServiceId{*postBlueWhateversRsp.BlueWhatevers[0].Id, *postBlueWhateversRsp.BlueWhatevers[1].Id}, sdk.WhateverRelationName.WhateverKnowsBlueWhatevers, sdk.RelationOperation.Add)
			if err != nil {
				return
			}

			getReq := sdk.GetWhateversRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Id,
					Id: &sdk.Id{
						Kind:      &sdk.IdKind.ServiceId,
						ServiceId: whatevers[0].Id,
					},
				},
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Is: ptr.String(svcName),
						},
					},
				},
				Select: &sdk.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							Value: ptr.Bool(true),
						},
						AlternativeIds: &sdk.IdSelect{
							Kind: ptr.Bool(true),
							Name: ptr.Bool(true),
							Email: &sdk.EmailSelect{
								Value: ptr.Bool(true),
							},
						},
						Relations: &sdk.WhateverRelationsSelect{
							KnowsBlueWhatevers: &sdk.BlueWhateversCollectionSelect{
								Meta: GetCollectionMetaSelect(),
								BlueWhatevers: &sdk.BlueWhateverSelect{
									Id: &sdk.ServiceIdSelect{
										Value: ptr.Bool(true),
									},
								},
							},
						},
					},
				},
				Relations: &sdk.GetWhateversRelations{
					KnowsBlueWhatevers: &sdk.GetBlueWhateversCollection{
						Relations: &sdk.GetBlueWhateversRelations{
							KnewByWhatevers: &sdk.GetWhateversCollection{},
						},
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			requireSvcIdValues(t, gGetRsp.MustGenericSlice(fieldnames.Whatevers).Get()[0].MustGenericSlice(fieldnames.Relations, "knowsBlueWhatevers", fieldnames.BlueWhatevers), []string{*blueWhatevers[0].Id.Value, *blueWhatevers[1].Id.Value})

			requireSvcIdValues(t, gGetRsp.MustGenericSlice(fieldnames.Whatevers).Get()[0].MustGenericSlice(fieldnames.Relations, "knowsBlueWhatevers", fieldnames.BlueWhatevers).Get()[0].MustGenericSlice(fieldnames.Relations, "knewByWhatevers", fieldnames.Whatevers), []string{*whatevers[0].Id.Value})
			requireSvcIdValues(t, gGetRsp.MustGenericSlice(fieldnames.Whatevers).Get()[0].MustGenericSlice(fieldnames.Relations, "knowsBlueWhatevers", fieldnames.BlueWhatevers).Get()[1].MustGenericSlice(fieldnames.Relations, "knewByWhatevers", fieldnames.Whatevers), []string{*whatevers[0].Id.Value})

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}
