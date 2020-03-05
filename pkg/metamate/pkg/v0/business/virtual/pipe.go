// +build !lite

package virtual

import (
	"context"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
	"github.com/metamatex/metamatemono/pkg/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[Pipe] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, opts types.VirtualSvcOpts) (h http.Handler, t string, err error) {
		h = httpjson.NewServer(httpjson.ServerOpts{
			Root:    rn,
			Factory: f,
			Handler: func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic) {
				switch gReq.Type().Name() {
				case sdk.LookupServiceRequestName:
					return f.MustFromStruct(sdk.LookupServiceResponse{
						Output: &sdk.LookupServiceOutput{
							Service: &sdk.Service{
								Endpoints: &sdk.Endpoints{
									PipeWhatevers: &sdk.PipeWhateversEndpoint{
										Filter: &sdk.PipeWhateversRequestFilter{
											Mode: &sdk.PipeModeFilter{
												Context: &sdk.ContextPipeModeFilter{
													Method: &sdk.EnumFilter{
														Is: ptr.String("post"),
													},
												},
											},
										},
									},
								},
							},
						},
					})
				case sdk.PipeWhateversRequestName:
					req := sdk.PipeWhateversRequest{}
					gReq.MustToStruct(&req)

					rsp := sdk.PipeWhateversResponse{}
					rsp.Context = req.Context

					return f.MustFromStruct(rsp)
				}

				return
			},
			LogErr: nil,
		})

		t = sdk.ServiceTransport.HttpJson

		return
	}
}
