package communication

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/metamatex/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/pkg/services/sqlx-svc/gen/v0/sdk"
	"github.com/metamatex/metamatemono/pkg/services/sqlx-svc/gen/v0/sdk/utils/ptr"
)

func GetServer(supportedIdKinds map[string]bool, db sqlx.Ext, rn *graph.RootNode, f generic.Factory, gGetServiceRsp generic.Generic) (func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic)) {
	return func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic) {
		errs := func() (errs []error) {
			en := gReq.Type().Edges.Endpoint.BelongsTo()

			switch en.Data.Method {
			case graph.MethodPost:
				gRsp, errs = handlePost(ctx, db, f, gReq)
			case graph.MethodGet:
				gRsp, errs = handleGet(ctx, supportedIdKinds, rn, db, f, gReq)
			case graph.MethodPut:
				gRsp, errs = handlePut(ctx, rn, db, f, gReq)
			case graph.MethodDelete:
				gRsp, errs = handleDelete(ctx, db, f, gReq)
			case graph.MethodAction:
				switch en.Name() {
				case endpointnames.LookupService:
					gRsp = gGetServiceRsp
				}
			}

			return
		}()
		if len(errs) != 0 {
			gRsp.MustSetGenericSlice([]string{fieldnames.Meta, fieldnames.Errors}, f.MustFromStructs(ToMessageErrors(errs)))
		}

		return
	}
}

func ToMessageErrors(errs []error) (errs0 []sdk.Error) {
	for _, err := range errs {
		errs0 = append(errs0, sdk.Error{
			Message: &sdk.Text{
				Value: ptr.String(err.Error()),
			},
		})
	}

	return
}
