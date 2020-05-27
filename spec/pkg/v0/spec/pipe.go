package spec

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"

	"testing"
)

func TestPipe(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	t.Run("TestPipe", func(t *testing.T) {
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

			gRsp.Sprint()

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}
