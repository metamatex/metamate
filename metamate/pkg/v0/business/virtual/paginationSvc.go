package virtual

import (
	"context"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/sdk"

	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"net/http"
)

func init() {
	handler[Pagination] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, vSvc types.VirtualSvc) (h http.Handler, t string, err error) {
		ws := []sdk.Whatever{
			{
				Id: &sdk.ServiceId{
					Value: sdk.String("0"),
				},
				StringField: sdk.String("a"),
			},
			{
				Id: &sdk.ServiceId{
					Value: sdk.String("1"),
				},
				StringField: sdk.String("b"),
			},
			{
				Id: &sdk.ServiceId{
					Value: sdk.String("2"),
				},
				StringField: sdk.String("c"),
			},
		}

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
									GetWhatevers: &sdk.GetWhateversEndpoint{
										Filter: &sdk.GetWhateversRequestFilter{},
									},
								},
							},
						},
					})
				case sdk.GetWhateversRequestName:
					var req = sdk.GetWhateversRequest{}
					gReq.MustToStruct(&req)

					var ws0 []sdk.Whatever
					var pagination *sdk.Pagination

					switch *req.Mode.Kind {
					case sdk.GetModeKind.Collection, sdk.GetModeKind.Relation, sdk.GetModeKind.Search:
						ws0, pagination = getWhateversCollection(ws, req)
					}

					return f.MustFromStruct(sdk.GetWhateversResponse{
						Whatevers: ws0,
						Pagination: pagination,
					})
				}

				return
			},
			LogErr: nil,
		})

		t = sdk.ServiceTransport.HttpJson

		return
	}
}

func getWhateversCollection(ws []sdk.Whatever, req sdk.GetWhateversRequest) (ws0 []sdk.Whatever, pagination *sdk.Pagination) {
	var page sdk.ServicePage

	if len(req.Pages) == 1 {
		page = req.Pages[0]
	} else {
		page = sdk.ServicePage{
			Page: &sdk.Page{
				Kind: &sdk.PageKind.IndexPage,
				IndexPage: &sdk.IndexPage{
					Value: sdk.Int32(0),
				},
			},
		}
	}

	pagination = &sdk.Pagination{}

	pagination.Current = []sdk.ServicePage{page}

	if int(*page.Page.IndexPage.Value) < len(ws) {
		pagination.Next = []sdk.ServicePage{
			{
				Page: &sdk.Page{
					Kind: &sdk.PageKind.IndexPage,
					IndexPage: &sdk.IndexPage{
						Value: sdk.Int32(*page.Page.IndexPage.Value + 1),
					},
				},
			},
		}
	}

	if *page.Page.IndexPage.Value != 0 {
		pagination.Previous = []sdk.ServicePage{
			{
				Page: &sdk.Page{
					Kind: &sdk.PageKind.IndexPage,
					IndexPage: &sdk.IndexPage{
						Value: sdk.Int32(*page.Page.IndexPage.Value - 1),
					},
				},
			},
		}
	}

	if int(*page.Page.IndexPage.Value) + 1 >= len(ws) {
		return
	}

	ws0 = []sdk.Whatever{
		ws[*page.Page.IndexPage.Value],
	}

	return
}
