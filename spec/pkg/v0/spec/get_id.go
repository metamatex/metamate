package spec

import (
	"context"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetModeId(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName, suffix string) {
	name := "TestGetModeId"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			postReq := mql.PostWhateversRequest{
				ServiceFilter: &mql.ServiceFilter{
					Id: &mql.ServiceIdFilter{
						Value: &mql.StringFilter{
							Is: mql.String(svcName),
						},
					},
				},
				Select: &mql.PostWhateversResponseSelect{
					Meta: GetResponseMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							Value: mql.Bool(true),
						},
						StringField: mql.Bool(true),
					},
				},
				Whatevers: []mql.Whatever{
					{
						Id: &mql.ServiceId{
							Value: mql.String(nameSvcId(suffix, name, "0")),
						},
						StringField: mql.String("hi"),
					},
				},
			}

			gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
			if err != nil {
				return
			}

			postRsp := mql.PostWhateversResponse{}
			gPostRsp.MustToStruct(&postRsp)

			requirePostRsp(t, gPostRsp)

			getReq := mql.GetWhateversRequest{
				ServiceFilter: &mql.ServiceFilter{
					Id: &mql.ServiceIdFilter{
						Value: &mql.StringFilter{
							Is: mql.String(svcName),
						},
					},
				},
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Id,
					Id: &mql.Id{
						Kind: &mql.IdKind.ServiceId,
						ServiceId: &mql.ServiceId{
							Value: postRsp.Whatevers[0].Id.Value,
						},
					},
				},
				Select: &mql.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							Value: mql.Bool(true),
						},
						StringField: mql.Bool(true),
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			getRsp := mql.GetWhateversResponse{}
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
			getReq := mql.GetWhateversRequest{
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Id,
					Id: &mql.Id{
						Kind: &mql.IdKind.ServiceId,
						ServiceId: &mql.ServiceId{
							Value: mql.String(""),
						},
					},
				},
				Select: &mql.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							Value: mql.Bool(true),
						},
						StringField: mql.Bool(true),
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			getRsp := mql.GetWhateversResponse{}
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
			postReq := mql.PostWhateversRequest{
				ServiceFilter: &mql.ServiceFilter{
					Id: &mql.ServiceIdFilter{
						Value: &mql.StringFilter{
							Is: mql.String(svcNameA),
						},
					},
				},
				Select: &mql.PostWhateversResponseSelect{
					Meta: GetResponseMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							Value: mql.Bool(true),
						},
					},
				},
				Whatevers: []mql.Whatever{
					{
						Id: &mql.ServiceId{
							Value: mql.String(nameSvcId(suffix, name, "0")),
						},
					},
				},
			}

			gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
			if err != nil {
				return
			}

			postRsp := mql.PostWhateversResponse{}
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
			postReq := mql.PostWhateversRequest{
				ServiceFilter: &mql.ServiceFilter{
					Id: &mql.ServiceIdFilter{
						Value: &mql.StringFilter{
							Is: mql.String(svcName),
						},
					},
				},
				Select: &mql.PostWhateversResponseSelect{
					Meta: GetResponseMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							Value: mql.Bool(true),
						},
						AlternativeIds: &mql.IdSelect{
							Kind: mql.Bool(true),
							Name: mql.Bool(true),
						},
						StringField: mql.Bool(true),
					},
				},
				Whatevers: []mql.Whatever{
					{
						AlternativeIds: []mql.Id{
							{
								Kind: &mql.IdKind.Name,
								Name: mql.String(nameSvcId(suffix, name, "0")),
							},
						},
						StringField: mql.String("a"),
					},
				},
			}

			gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
			if err != nil {
				return
			}

			postRsp := mql.PostWhateversResponse{}
			gPostRsp.MustToStruct(&postRsp)

			requirePostRsp(t, gPostRsp)

			getReq := mql.GetWhateversRequest{
				ServiceFilter: &mql.ServiceFilter{
					Id: &mql.ServiceIdFilter{
						Value: &mql.StringFilter{
							Is: mql.String(svcName),
						},
					},
				},
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Id,
					Id: &mql.Id{
						Kind: &mql.IdKind.Name,
						Name: postRsp.Whatevers[0].AlternativeIds[0].Name,
					},
				},
				Select: &mql.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							Value: mql.Bool(true),
						},
						AlternativeIds: &mql.IdSelect{
							Kind: mql.Bool(true),
							Name: mql.Bool(true),
						},
						StringField: mql.Bool(true),
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			requireGetRsp(t, gGetRsp)

			getRsp := mql.GetWhateversResponse{}
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
					Kind: &mql.GetModeKind.Id,
					Id: &mql.Id{
						Kind:      &mql.IdKind.ServiceId,
						ServiceId: whatevers[0].Id,
					},
				},
				ServiceFilter: &mql.ServiceFilter{
					Id: &mql.ServiceIdFilter{
						Value: &mql.StringFilter{
							Is: mql.String(svcName),
						},
					},
				},
				Select: &mql.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							Value: mql.Bool(true),
						},
						AlternativeIds: &mql.IdSelect{
							Kind: mql.Bool(true),
							Name: mql.Bool(true),
							Email: &mql.EmailSelect{
								Value: mql.Bool(true),
							},
						},
						Relations: &mql.WhateverRelationsSelect{
							KnowsWhatevers: &mql.WhateversCollectionSelect{
								Meta: GetCollectionMetaSelect(),
								Whatevers: &mql.WhateverSelect{
									Id: &mql.ServiceIdSelect{
										Value: mql.Bool(true),
									},
								},
							},
						},
					},
				},
				Relations: &mql.GetWhateversRelations{
					KnowsWhatevers: &mql.GetWhateversCollection{
						Relations: &mql.GetWhateversRelations{
							KnewByWhatevers: &mql.GetWhateversCollection{},
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
			whatevers := []mql.Whatever{
				{
					Id: &mql.ServiceId{
						Value: mql.String(nameSvcId(suffix, name, "a")),
					},
				},
			}

			blueWhatevers := []mql.BlueWhatever{
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
			}

			postWhateversRsp := requirePostWhatevers(t, ctx, f, h, svcName, whatevers)

			postBlueWhateversRsp := requirePostBlueWhatevers(t, ctx, f, h, svcName, blueWhatevers)

			err = requirePutBlueWhatevers(t, ctx, f, h, postWhateversRsp.Whatevers[0].Id, []mql.ServiceId{*postBlueWhateversRsp.BlueWhatevers[0].Id, *postBlueWhateversRsp.BlueWhatevers[1].Id}, mql.WhateverRelationName.WhateverKnowsBlueWhatevers, mql.RelationOperation.Add)
			if err != nil {
				return
			}

			getReq := mql.GetWhateversRequest{
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Id,
					Id: &mql.Id{
						Kind:      &mql.IdKind.ServiceId,
						ServiceId: whatevers[0].Id,
					},
				},
				ServiceFilter: &mql.ServiceFilter{
					Id: &mql.ServiceIdFilter{
						Value: &mql.StringFilter{
							Is: mql.String(svcName),
						},
					},
				},
				Select: &mql.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &mql.WhateverSelect{
						Id: &mql.ServiceIdSelect{
							Value: mql.Bool(true),
						},
						AlternativeIds: &mql.IdSelect{
							Kind: mql.Bool(true),
							Name: mql.Bool(true),
							Email: &mql.EmailSelect{
								Value: mql.Bool(true),
							},
						},
						Relations: &mql.WhateverRelationsSelect{
							KnowsBlueWhatevers: &mql.BlueWhateversCollectionSelect{
								Meta: GetCollectionMetaSelect(),
								BlueWhatevers: &mql.BlueWhateverSelect{
									Id: &mql.ServiceIdSelect{
										Value: mql.Bool(true),
									},
								},
							},
						},
					},
				},
				Relations: &mql.GetWhateversRelations{
					KnowsBlueWhatevers: &mql.GetBlueWhateversCollection{
						Relations: &mql.GetBlueWhateversRelations{
							KnewByWhatevers: &mql.GetWhateversCollection{},
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
