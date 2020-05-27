package spec

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func getIdsFilter(suffix, caseName string) (idFilter *mql.ServiceIdFilter) {
	return &mql.ServiceIdFilter{
		Value: &mql.StringFilter{
			Contains: mql.String(nameSvcCase(suffix, caseName)),
		},
	}
}

func nameSvcId(suffix, caseName, id string) string {
	return nameSvcCase(suffix, caseName) + "_" + id
}

func nameSvcCase(suffix, caseName string) string {
	return suffix + "_" + caseName
}

func TestFilterStringIs(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName, suffix string) {
	name := "TestFilterStringIs"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			postReq := &mql.PostWhateversRequest{
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
						StringField: mql.Bool(true),
					},
				},
				Whatevers: []mql.Whatever{
					{
						Id: &mql.ServiceId{
							Value: mql.String(nameSvcId(suffix, name, "0")),
						},
						StringField: mql.String("a"),
					},
					{
						Id: &mql.ServiceId{
							Value: mql.String(nameSvcId(suffix, name, "1")),
						},
						StringField: mql.String("b"),
					},
				},
			}

			gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
			if err != nil {
				return
			}

			requirePostRsp(t, gPostRsp)

			postRsp := mql.PostWhateversResponse{}
			gPostRsp.MustToStruct(&postRsp)

			getReq := &mql.GetWhateversRequest{
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
				Filter: &mql.WhateverFilter{
					Id: getIdsFilter(suffix, name),
					StringField: &mql.StringFilter{
						Is: mql.String("a"),
					},
				},
				Select: &mql.GetWhateversResponseSelect{
					Meta: GetCollectionMetaSelect(),
					Whatevers: &mql.WhateverSelect{
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

			require.Len(t, getRsp.Whatevers, 1)
			assert.Equal(t, getRsp.Whatevers[0].StringField, postRsp.Whatevers[0].StringField)

			return
		}()
		if err != nil {
			t.Error(err)
		}

	})

}
