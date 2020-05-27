package spec

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"

	"testing"
)

func TestGetModeCollection(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string) {
	name := "TestGetModeCollection"
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
				Mode: &mql.PostMode{
					Kind:       &mql.PostModeKind.Collection,
					Collection: &mql.CollectionPostMode{},
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
							Value: mql.String(name),
						},
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

			getReq := mql.GetWhateversRequest{
				ServiceFilter: &mql.ServiceFilter{
					Id: &mql.ServiceIdFilter{
						Value: &mql.StringFilter{
							Is: mql.String(svcName),
						},
					},
				},
				Mode: &mql.GetMode{
					Kind:       &mql.GetModeKind.Collection,
					Collection: &mql.CollectionGetMode{},
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

			return
		}()
		if err != nil {
			return
		}

	})
}
