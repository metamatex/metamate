package spec

import (
	"context"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/gen/v0/mql"
	
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPost(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string) {
	t.Run("TestPost", func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			postReq := sdk.PostWhateversRequest{
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Is: sdk.String(svcName),
						},
					},
				},
				Select: &sdk.PostWhateversResponseSelect{
					Meta: GetResponseMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							Value: sdk.Bool(true),
						},
						StringField: sdk.Bool(true),
					},
				},
				Whatevers: []sdk.Whatever{
					{
						StringField: sdk.String("a"),
					},
					{
						StringField: sdk.String("b"),
					},
				},
			}

			gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
			if err != nil {
			    return
			}

			requirePostRsp(t, gPostRsp)

			postRsp := sdk.GetWhateversResponse{}
			gPostRsp.MustToStruct(&postRsp)

			for i, _ := range postReq.Whatevers {
				assert.Equal(t, *postReq.Whatevers[i].StringField, *postRsp.Whatevers[i].StringField)
				assert.NotEqual(t, "", *postRsp.Whatevers[i].Id.Value)
			}

			return
		}()
		if err != nil {
		    t.Error(err)
		}
	})
}

func TestEmptyPost(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	t.Run("TestEmptyPost", func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			table := []struct {
				req sdk.PostWhateversRequest
			}{
				{
					req: sdk.PostWhateversRequest{},
				},
			}

			for _, c := range table {
				gRsp, err := h(ctx, f.MustFromStruct(c.req))
				if err != nil {
					t.Error(err)
				}

				assert.NotNil(t, gRsp)
			}

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestRequestFilter(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	t.Run("TestRequestFilter", func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			var req = sdk.PostWhateversRequest{
				Mode: &sdk.PostMode{
					Kind: &sdk.PostModeKind.Collection,
					Collection: &sdk.CollectionPostMode{},
				},
				Select: &sdk.PostWhateversResponseSelect{
					Meta: GetResponseMetaSelect(),
				},
				Whatevers: []sdk.Whatever{
					{
						Id: &sdk.ServiceId{
							Value: sdk.String("match"),
						},
					},
				},
			}

			gRsp, err := h(ctx, f.MustFromStruct(req))
			if err != nil {
				t.Error(err)
			}

			gRsp.Print()

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestPostWithNameId(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName, suffix string) {
	t.Run("TestPostWithNameId", func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			postReq := sdk.PostWhateversRequest{
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Is: sdk.String(svcName),
						},
					},
				},
				Select: &sdk.PostWhateversResponseSelect{
					Meta: GetResponseMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
						Id: &sdk.ServiceIdSelect{
							Value: sdk.Bool(true),
						},
						AlternativeIds: &sdk.IdSelect{
							Kind: sdk.Bool(true),
							Name: sdk.Bool(true),
						},
						StringField: sdk.Bool(true),
					},
				},
				Whatevers: []sdk.Whatever{
					{
						AlternativeIds: []sdk.Id{
							{
								Kind: &sdk.IdKind.Name,
								Name: sdk.String("a" + suffix),
							},
						},
						StringField: sdk.String("a"),
					},
					{
						AlternativeIds: []sdk.Id{
							{
								Kind: &sdk.IdKind.Name,
								Name: sdk.String("b" + suffix),
							},
						},
						StringField: sdk.String("b"),
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

			for i, _ := range postReq.Whatevers {
				assert.Equal(t, *postReq.Whatevers[i].StringField, *postRsp.Whatevers[i].StringField)
				assert.Equal(t, *postReq.Whatevers[i].AlternativeIds[0].Name, *postRsp.Whatevers[i].AlternativeIds[0].Name)
				assert.NotEqual(t, "", *postRsp.Whatevers[i].Id.Value)
			}

			return
		}()
		if err != nil {
		    t.Error(err)
		}
	})
}