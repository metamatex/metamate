package spec

import (
	"context"
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
	"testing"
)

func TestGetModeCollection(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string) {
	name := "TestGetModeCollection"
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
				Mode: &sdk.PostMode{
					Kind: &sdk.PostModeKind.Collection,
					Collection: &sdk.CollectionPostMode{},
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
							Value: ptr.String(name),
						},
						StringField: ptr.String("a"),
					},
					{
						StringField: ptr.String("b"),
					},
				},
			}

			gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
			if err != nil {
				return
			}

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
					Kind: &sdk.GetModeKind.Collection,
					Collection: &sdk.CollectionGetMode{},
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

			return
		}()
		if err != nil {
			return
		}

	})
}
