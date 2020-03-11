package spec

import (
	"context"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"testing"
)

func TestDiscovery(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestDiscovery"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetServicesRequest{
				Mode: &sdk.GetMode{
					Kind:       &sdk.GetModeKind.Collection,
					Collection: &sdk.CollectionGetMode{},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestEmptyServiceFilter(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestEmptyServiceFilter"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetServicesRequest{
				Mode: &sdk.GetMode{
					Kind:       &sdk.GetModeKind.Collection,
					Collection: &sdk.CollectionGetMode{},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}
