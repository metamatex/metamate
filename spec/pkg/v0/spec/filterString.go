package spec

import (
	"context"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/utils/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func getIdsFilter(suffix, caseName string) (idFilter *sdk.ServiceIdFilter) {
	return &sdk.ServiceIdFilter{
		Value: &sdk.StringFilter{
			Contains: ptr.String(nameSvcCase(suffix, caseName)),
		},
	}
}

func nameSvcId(suffix, caseName, id string) (string) {
	return nameSvcCase(suffix, caseName) + "_" + id
}

func nameSvcCase(suffix, caseName string) (string) {
	return suffix + "_" + caseName
}

func TestFilterStringIs(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName, suffix string) {
	name := "TestFilterStringIs"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			postReq := &sdk.PostWhateversRequest{
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
						StringField: ptr.Bool(true),
					},
				},
				Whatevers: []sdk.Whatever{
					{
						Id: &sdk.ServiceId{
							Value: ptr.String(nameSvcId(suffix, name, "0")),
						},
						StringField: ptr.String("a"),
					},
					{
						Id: &sdk.ServiceId{
							Value: ptr.String(nameSvcId(suffix, name, "1")),
						},
						StringField: ptr.String("b"),
					},
				},
			}

			gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
			if err != nil {
				return
			}

			requirePostRsp(t, gPostRsp)

			postRsp := sdk.PostWhateversResponse{}
			gPostRsp.MustToStruct(&postRsp)

			getReq := &sdk.GetWhateversRequest{
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
				Filter: &sdk.WhateverFilter{
					Id: getIdsFilter(suffix, name),
					StringField: &sdk.StringFilter{
						Is: ptr.String("a"),
					},
				},
				Select: &sdk.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &sdk.WhateverSelect{
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

			require.Len(t, getRsp.Whatevers, 1)
			assert.Equal(t, getRsp.Whatevers[0].StringField, postRsp.Whatevers[0].StringField)

			return
		}()
		if err != nil {
			t.Error(err)
		}

	})

}
