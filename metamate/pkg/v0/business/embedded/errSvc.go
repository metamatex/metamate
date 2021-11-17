package embedded

import (
	"context"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[Error] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, vSvc types.EmbeddedSvc) (h http.Handler, err error) {
		h = generic.NewServer(generic.ServerOpts{
			Root:    rn,
			Factory: f,
			Handler: func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic) {
				switch gReq.Type().Name() {
				case mql.LookupServiceBusRequestName:
					return f.MustFromStruct(mql.LookupServiceServiceResponse{
						Output: &mql.LookupServiceOutput{
							Service: &mql.Service{
								Endpoints: &mql.ServiceEndpoints{
									GetDummies: &mql.GetDummiesEndpoint{
										Filter: &mql.GetDummiesBusRequestFilter{},
									},
								},
							},
						},
					})
				case mql.GetDummiesBusRequestName:
					return f.MustFromStruct(mql.GetDummiesServiceResponse{
						Errors: []mql.Error{
							{
								Message: mql.String("a"),
							},
						},
						Dummies: []mql.Dummy{
							{
								Id: &mql.ServiceId{
									Value: mql.String("a"),
								},
								StringField: mql.String("a"),
							},
						},
					})
				}

				return
			},
			LogErr: nil,
		})

		return
	}
}
