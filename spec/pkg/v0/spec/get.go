package spec

import (
	"context"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyGet(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	t.Run("TestEmptyGet", func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			table := []struct {
				name          string
				req           sdk.GetWhateversRequest
				nExpectedErrs int
			}{
				{
					name:          "0",
					req:           sdk.GetWhateversRequest{},
					nExpectedErrs: 0,
				},
				{
					name: "1",
					req: sdk.GetWhateversRequest{
						Mode: &sdk.GetMode{},
					},
					nExpectedErrs: 2,
				},
				{
					name: "2",
					req: sdk.GetWhateversRequest{
						Mode: &sdk.GetMode{
							Kind: &sdk.GetModeKind.Collection,
						},
					},
					nExpectedErrs: 1,
				},
				{
					name: "3",
					req: sdk.GetWhateversRequest{
						Mode: &sdk.GetMode{
							Kind:       &sdk.GetModeKind.Collection,
							Collection: &sdk.CollectionGetMode{},
						},
					},
					nExpectedErrs: 0,
				},
				{
					name: "4",
					req: sdk.GetWhateversRequest{
						Mode: &sdk.GetMode{
							Kind:       &sdk.GetModeKind.Collection,
							Collection: &sdk.CollectionGetMode{},
						},
						Select: &sdk.GetWhateversResponseSelect{},
					},
					nExpectedErrs: 0,
				},
			}

			for _, c := range table {
				t.Run(c.name, func(t *testing.T) {
					gRsp, err := h(ctx, f.MustFromStruct(c.req))
					if err != nil {
						t.Error(err)
					}

					if !assert.NotNil(t, gRsp) {
						return
					}

					if !assertNRspMetaErrs(t, gRsp, c.nExpectedErrs) {
						gRsp.Print()
					}
				})
			}

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}
