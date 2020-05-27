package spec

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyGet(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	t.Run("TestEmptyGet", func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			table := []struct {
				name          string
				req           mql.GetWhateversRequest
				nExpectedErrs int
			}{
				{
					name:          "0",
					req:           mql.GetWhateversRequest{},
					nExpectedErrs: 0,
				},
				{
					name: "1",
					req: mql.GetWhateversRequest{
						Mode: &mql.GetMode{},
					},
					nExpectedErrs: 2,
				},
				{
					name: "2",
					req: mql.GetWhateversRequest{
						Mode: &mql.GetMode{
							Kind: &mql.GetModeKind.Collection,
						},
					},
					nExpectedErrs: 1,
				},
				{
					name: "3",
					req: mql.GetWhateversRequest{
						Mode: &mql.GetMode{
							Kind:       &mql.GetModeKind.Collection,
							Collection: &mql.CollectionGetMode{},
						},
					},
					nExpectedErrs: 0,
				},
				{
					name: "4",
					req: mql.GetWhateversRequest{
						Mode: &mql.GetMode{
							Kind:       &mql.GetModeKind.Collection,
							Collection: &mql.CollectionGetMode{},
						},
						Select: &mql.GetWhateversResponseSelect{},
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
