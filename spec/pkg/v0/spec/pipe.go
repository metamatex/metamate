package spec

import (
	"context"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/gen/v0/mql"
	
	"testing"
)

func TestPipe(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	t.Run("TestPipe", func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			var req = sdk.PostWhateversRequest{
				Mode: &sdk.PostMode{
					Kind:       &sdk.PostModeKind.Collection,
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

			gRsp.Sprint()

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}
