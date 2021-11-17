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
	handler[Pagination] = func(f generic.Factory, rn *graph.RootNode, c *http.Client, vSvc types.EmbeddedSvc) (h http.Handler, err error) {
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
					var req = mql.GetDummiesBusRequest{}
					gReq.MustToStruct(&req)

					var ws0 []mql.Dummy
					var pagination *mql.ServicePagination

					switch *req.Mode.Kind {
					case mql.GetModeKind.Collection, mql.GetModeKind.Relation, mql.GetModeKind.Search:
						ws0, pagination = getDummiesCollection(ws, req)
					}

					return f.MustFromStruct(mql.GetDummiesServiceResponse{
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

func getDummiesCollection(ws []mql.Dummy, req mql.GetDummiesBusRequest) (ws0 []mql.Dummy, pagination *mql.ServicePagination) {
	p := req.Page
	if p == nil {
		p = &mql.Page{
			Kind: &mql.PageKind.IndexPage,
			IndexPage: &mql.IndexPage{
				Value: mql.Int32(0),
			},
		}
	}

	pagination = &mql.ServicePagination{}

	pagination.Current = p

	if int(*p.IndexPage.Value) < len(ws) {
		pagination.Next = &mql.Page{
			Kind: &mql.PageKind.IndexPage,
			IndexPage: &mql.IndexPage{
				Value: mql.Int32(*p.IndexPage.Value + 1),
			},
		}
	}

	if *p.IndexPage.Value != 0 {
		pagination.Previous = &mql.Page{
			Kind: &mql.PageKind.IndexPage,
			IndexPage: &mql.IndexPage{
				Value: mql.Int32(*p.IndexPage.Value - 1),
			},
		}
	}

	if int(*p.IndexPage.Value)+1 >= len(ws) {
		return
	}

	ws0 = []mql.Dummy{
		ws[*p.IndexPage.Value],
	}

	return
}
