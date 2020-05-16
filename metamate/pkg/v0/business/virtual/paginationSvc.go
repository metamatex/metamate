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
	handler[Pagination] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, vSvc types.VirtualSvc) (h http.Handler, err error) {
		ws := []mql.Dummy{
			{
				Id: &mql.ServiceId{
					Value: mql.String("0"),
				},
				StringField: mql.String("a"),
			},
			{
				Id: &mql.ServiceId{
					Value: mql.String("1"),
				},
				StringField: mql.String("b"),
			},
			{
				Id: &mql.ServiceId{
					Value: mql.String("2"),
				},
				StringField: mql.String("c"),
			},
		}

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
					var req = mql.GetDummiesRequest{}
					gReq.MustToStruct(&req)

					var ws0 []mql.Dummy
					var pagination *mql.Pagination

					switch *req.Mode.Kind {
					case mql.GetModeKind.Collection, mql.GetModeKind.Relation, mql.GetModeKind.Search:
						ws0, pagination = getDummiesCollection(ws, req)
					}

					return f.MustFromStruct(mql.GetDummiesResponse{
						Dummies:    ws0,
						Pagination: pagination,
					})
				}

				return
			},
			LogErr: nil,
		})

		return
	}
}

func getDummiesCollection(ws []mql.Dummy, req mql.GetDummiesRequest) (ws0 []mql.Dummy, pagination *mql.Pagination) {
	var page mql.ServicePage

	if len(req.Pages) == 1 {
		page = req.Pages[0]
	} else {
		page = mql.ServicePage{
			Page: &mql.Page{
				Kind: &mql.PageKind.IndexPage,
				IndexPage: &mql.IndexPage{
					Value: mql.Int32(0),
				},
			},
		}
	}

	pagination = &mql.Pagination{}

	pagination.Current = []mql.ServicePage{page}

	if int(*page.Page.IndexPage.Value) < len(ws) {
		pagination.Next = []mql.ServicePage{
			{
				Page: &mql.Page{
					Kind: &mql.PageKind.IndexPage,
					IndexPage: &mql.IndexPage{
						Value: mql.Int32(*page.Page.IndexPage.Value + 1),
					},
				},
			},
		}
	}

	if *page.Page.IndexPage.Value != 0 {
		pagination.Previous = []mql.ServicePage{
			{
				Page: &mql.Page{
					Kind: &mql.PageKind.IndexPage,
					IndexPage: &mql.IndexPage{
						Value: mql.Int32(*page.Page.IndexPage.Value - 1),
					},
				},
			},
		}
	}

	if int(*page.Page.IndexPage.Value)+1 >= len(ws) {
		return
	}

	ws0 = []mql.Dummy{
		ws[*page.Page.IndexPage.Value],
	}

	return
}
