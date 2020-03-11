// +build !lite

package virtual

import (
	"context"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/utils/ptr"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[ReqFilter] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, opts types.VirtualSvcOpts) (h http.Handler, t string, err error) {
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
									PostWhatevers: &sdk.PostWhateversEndpoint{
										Filter: &sdk.PostWhateversRequestFilter{
											Whatevers: &sdk.WhateverListFilter{
												Some: &sdk.WhateverFilter{
													Id: &sdk.ServiceIdFilter{
														Value: &sdk.StringFilter{
															Is: ptr.String("match"),
														},
													},
												},
											},
										},
									},
								},
							},
						},
					})
				case sdk.PostWhateversRequestName:
					println("RequestFilterService.PostWhatevers called")

					return f.MustFromStruct(sdk.PostWhateversResponse{})
				}

				return
			},
			LogErr: nil,
		})

		t = sdk.ServiceTransport.HttpJson

		return
	}
}
