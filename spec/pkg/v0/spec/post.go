package spec

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPost(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string) {
	t.Run("TestPost", func(t *testing.T) {
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
						StringField: mql.String("a"),
					},
					{
						StringField: mql.String("b"),
					},
				},
			}

			gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
			if err != nil {
				return
			}

			requirePostRsp(t, gPostRsp)

			postRsp := mql.GetWhateversResponse{}
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
				req mql.PostWhateversRequest
			}{
				{
					req: mql.PostWhateversRequest{},
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
			var req = mql.PostWhateversRequest{
				Mode: &mql.PostMode{
					Kind:       &mql.PostModeKind.Collection,
					Collection: &mql.CollectionPostMode{},
				},
				Select: &mql.PostWhateversResponseSelect{
					Meta: GetResponseMetaSelect(),
				},
				Whatevers: []mql.Whatever{
					{
						Id: &mql.ServiceId{
							Value: mql.String("match"),
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
								Name: mql.String("a" + suffix),
							},
						},
						StringField: mql.String("a"),
					},
					{
						AlternativeIds: []mql.Id{
							{
								Kind: &mql.IdKind.Name,
								Name: mql.String("b" + suffix),
							},
						},
						StringField: mql.String("b"),
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
