package virtual

import (
	"context"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"

	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[ReqFilter] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, vSvc types.VirtualSvc) (h http.Handler, err error) {
		h = generic.NewServer(generic.ServerOpts{
			Root:    rn,
			Factory: f,
			Handler: func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic) {
				switch gReq.Type().Name() {
				case mql.LookupServiceRequestName:
					return f.MustFromStruct(mql.LookupServiceResponse{
						Output: &mql.LookupServiceOutput{
							Service: &mql.Service{
								Endpoints: &mql.Endpoints{
									GetDummies: &mql.GetDummiesEndpoint{
										Filter: &mql.GetDummiesRequestFilter{},
									},
								},
							},
						},
					})
				case mql.GetDummiesRequestName:
					println("RequestFilterService.GetDummies called")

					return f.MustFromStruct(mql.GetDummiesResponse{})
				}

				return
			},
			LogErr: nil,
		})

		return
	}
}
