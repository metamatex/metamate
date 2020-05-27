package boot_test

import (
	"context"
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/boot"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/line"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/virtual"
	"github.com/metamatex/metamate/metamate/pkg/v0/config"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

const (
	AuthSvc     = "auth"
	PipeSvc     = "pipe"
	ReqFilter   = "request-filter"
	SqlxA       = "sqlx-a"
	SqlxB       = "sqlx-b"
	ErrorA      = "error-a"
	ErrorB      = "error-b"
	PaginationA = "pagination-a"
	PaginationB = "pagination-b"
)

func TestBoot(t *testing.T) {
	c := config.DefaultConfig

	c.Log.Internal = types.InternalLogConfig{
		config.SvcRsp: map[string]string{
			//mql.GetPostsResponseName: "{{ .Ctx.Svc.Url.Value }} : \n{{ .Ctx.GSvcReq.Type.Name }}\n{{ .Ctx.GSvcReq.Sprint }}\n\n{{ .Ctx.GSvcRsp.Type.Name }}\n{{ .Ctx.GSvcRsp.Sprint }}",
			//"*": "{{ .Ctx.Svc.Url.Value }} : {{ .Ctx.GSvcRsp.Type.Name }}",
		},
	}

	c.Virtual.Services = append(c.Virtual.Services, []types.VirtualSvc{
		{
			Id:   ErrorA,
			Name: virtual.Error,
		},
		{
			Id:   ErrorB,
			Name: virtual.Error,
		},
		{
			Id:   PaginationA,
			Name: virtual.Pagination,
		},
		{
			Id:   PaginationB,
			Name: virtual.Pagination,
		},
	}...)

	d, err := boot.NewDependencies(c, types.Version{})
	if err != nil {
		log.Fatal(err)
	}

	d.ResolveLine.(*line.Line).SetLog(false)

	ctx := context.Background()

	suffix := "-" + time.Now().Format("3:04:05PM")
	println(suffix)

	f := func(ctx context.Context, gCliReq generic.Generic) (gCliRsp generic.Generic, err error) {
		gCliReq.MustSetGeneric([]string{fieldnames.Select, fieldnames.Errors}, d.Factory.MustFromStruct(mql.ErrorSelect{
			Message: mql.Bool(true),
		}))

		gCliRsp = d.ServeFunc(ctx, gCliReq)

		return
	}

	//spec.TestRequestFilter(t, ctx, d.Factory, f)
	//
	//spec.TestPipe(t, ctx, d.Factory, f)

	//spec.TestDiscovery(t, ctx, d.Factory, f)

	//FTestHackernewsGetPostFeedContainsPosts(t, ctx, d.Factory, f)

	//FTestError(t, ctx, d.Factory, f)

	//FTestHackernews(t, ctx, d.Factory, f)

	//FTestPagination(t, ctx, d.Factory, f)
	//
	//FTestHackernews(t, ctx, d.Factory, f)

	//FTestHackernews(t, ctx, d.Factory, f)

	//FTestHackernewsSocialAccount(t, ctx, d.Factory, f)

	//spec.TestEmptyPost(t, ctx, d.Factory, f)
	//
	//spec.TestEmptyGet(t, ctx, d.Factory, f)
	//
	//spec.TestPost(t, ctx, d.Factory, f, SqlxA)
	//
	//spec.TestPostWithNameId(t, ctx, d.Factory, f, SqlxA, suffix)
	//
	//spec.TestGetModeId(t, ctx, d.Factory, f, SqlxA, suffix)
	//
	//spec.TestGetModeCollection(t, ctx, d.Factory, f, SqlxA)
	//
	//spec.TestFilterStringIs(t, ctx, d.Factory, f, SqlxA, suffix)
	//
	//spec.TestGetModeRelation(t, ctx, d.Factory, f, SqlxA, suffix)
	//
	//spec.TestGetModeIdWithNameId(t, ctx, d.Factory, f, suffix, SqlxA)
	//
	//spec.TestGetModeIdWithServiceFilter(t, ctx, d.Factory, f, suffix, SqlxA, SqlxB)
	//
	//spec.TestPostClientAccounts(t, ctx, d.Factory, f, SqlxA)
	//
	//spec.TestAuthenticateClientAccount(t, ctx, d.Factory, f, SqlxA, AuthSvc)
	//
	//spec.TestToken(t, ctx, d.Factory, f, AuthSvc, SqlxA)
	//
	//spec.TestGetModeRelationInter(t, ctx, d.Factory, f, SqlxA, SqlxB, suffix)
	//
	//spec.TestGetModeIdWithSelfReferencingRelation(t, ctx, d.Factory, f, suffix, SqlxA)
	//
	//spec.TestGetModeIdWithRelation(t, ctx, d.Factory, f, suffix, SqlxA)

	tester := boot.NewTester(false, t, ctx, d.Factory, d.RootNode, f)

	GeneralTest(tester, false)
	RedditTest(tester, false)
	HackerNewsTest(tester, false)
}

func GeneralTest(tester boot.Tester, print bool) {
	tester.NewSubTester(print, "general", func(tester boot.Tester) {
		tester.TestGetCollection(false, mql.TypeNames.Service, nil, nil)
	})
}

func RedditTest(tester boot.Tester, print bool) {
	tester.NewSubTester(print, "reddit", func(tester boot.Tester) {
		svcSvcId := mql.Id{
			Kind: &mql.IdKind.ServiceId,
			ServiceId: &mql.ServiceId{
				ServiceName: mql.String("discovery"),
				Value:       mql.String("reddit"),
			},
		}

		postFeedSvcId := mql.Id{
			Kind: &mql.IdKind.ServiceId,
			ServiceId: &mql.ServiceId{
				ServiceName: mql.String("reddit"),
				Value:       mql.String("graphql"),
			},
		}

		socialAccountSvcId := mql.Id{
			Kind: &mql.IdKind.ServiceId,
			ServiceId: &mql.ServiceId{
				ServiceName: mql.String("reddit"),
				Value:       mql.String("TheMrZZ0"),
			},
		}

		tester.TestGetId(false, mql.TypeNames.Service, svcSvcId, nil)
		tester.TestGetId(false, mql.TypeNames.PostFeed, postFeedSvcId, []string{mql.FieldNames.ContainsPosts})
		tester.TestGetRelation(false, mql.TypeNames.Post, postFeedSvcId, mql.PathNames.PostFeedContainsPosts, nil, 2)
		tester.TestGetRelation(false, mql.TypeNames.Post, socialAccountSvcId, mql.PathNames.SocialAccountAuthorsPosts, nil, 1)
		tester.TestGetId(false, mql.TypeNames.SocialAccount, socialAccountSvcId, []string{mql.FieldNames.AuthorsPosts})
	})
}

func HackerNewsTest(tester boot.Tester, print bool) {
	tester.NewSubTester(print, "hackernews", func(tester boot.Tester) {
		svcSvcId := mql.Id{
			Kind: &mql.IdKind.ServiceId,
			ServiceId: &mql.ServiceId{
				ServiceName: mql.String("discovery"),
				Value:       mql.String("hackernews"),
			},
		}

		svcFilter := mql.ServiceFilter{
			Id: &mql.ServiceIdFilter{
				Value: &mql.StringFilter{
					Is: mql.String("hackernews"),
				},
			},
		}

		socialAccountSvcId := mql.Id{
			Kind: &mql.IdKind.ServiceId,
			ServiceId: &mql.ServiceId{
				ServiceName: mql.String("hackernews"),
				Value:       mql.String("21stio"),
			},
		}

		postFeedSvcId := mql.Id{
			Kind: &mql.IdKind.ServiceId,
			ServiceId: &mql.ServiceId{
				ServiceName: mql.String("hackernews"),
				Value:       mql.String("topstories"),
			},
		}

		tester.TestGetId(false, mql.TypeNames.Service, svcSvcId, nil)
		tester.TestGetRelation(false, mql.TypeNames.Post, socialAccountSvcId, mql.PathNames.SocialAccountAuthorsPosts, nil, 1)
		tester.TestGetId(false, mql.TypeNames.SocialAccount, socialAccountSvcId, []string{mql.FieldNames.AuthorsPosts})
		tester.TestGetRelation(false, mql.TypeNames.Post, postFeedSvcId, mql.PathNames.PostFeedContainsPosts, nil, 1)
		tester.TestGetRelation(false, mql.TypeNames.Post, socialAccountSvcId, mql.PathNames.SocialAccountBookmarksPosts, nil, 1)
		tester.TestGetRelation(false, mql.TypeNames.Post, socialAccountSvcId, mql.PathNames.SocialAccountAuthorsPosts, nil, 1)
		tester.TestGetSearch(false, mql.TypeNames.Post, "books", &svcFilter, nil)
		tester.TestGetCollection(false, mql.TypeNames.PostFeed, &svcFilter, nil)
		//tester.TestGetId(mql.TypeNames.PostFeed, postFeedSvcId, []string{mql.FieldNames.ContainsPosts})
	})
}

func FTestPagination(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestPagination"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		t.Run("returnsPaginations", func(t *testing.T) {
			t.Parallel()

			tcs := []struct {
				mode *mql.GetMode
			}{
				{
					mode: &mql.GetMode{
						Kind:       &mql.GetModeKind.Collection,
						Collection: &mql.CollectionGetMode{},
					},
				},
				{
					mode: &mql.GetMode{
						Kind:     &mql.GetModeKind.Relation,
						Relation: &mql.RelationGetMode{},
					},
				},
				{
					mode: &mql.GetMode{
						Kind:   &mql.GetModeKind.Search,
						Search: &mql.SearchGetMode{},
					},
				},
			}

			for _, tc := range tcs {
				t.Run(*tc.mode.Kind, func(t *testing.T) {
					err := func() (err error) {
						getReq := mql.GetDummiesRequest{
							ServiceFilter: &mql.ServiceFilter{
								Id: &mql.ServiceIdFilter{
									Value: &mql.StringFilter{
										In: []string{PaginationA},
									},
								},
							},
							Mode: tc.mode,
						}

						gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
						if err != nil {
							return
						}

						_, ok := gGetRsp.Generic(fieldnames.Pagination)
						assert.True(t, ok)

						return
					}()
					if err != nil {
						t.Error(err)
					}
				})
			}
		})

		t.Run("addsServiceIds", func(t *testing.T) {
			t.Parallel()

			tcs := []struct {
				mode *mql.GetMode
			}{
				{
					mode: &mql.GetMode{
						Kind:       &mql.GetModeKind.Collection,
						Collection: &mql.CollectionGetMode{},
					},
				},
				{
					mode: &mql.GetMode{
						Kind:     &mql.GetModeKind.Relation,
						Relation: &mql.RelationGetMode{},
					},
				},
				{
					mode: &mql.GetMode{
						Kind:   &mql.GetModeKind.Search,
						Search: &mql.SearchGetMode{},
					},
				},
			}

			for _, tc := range tcs {
				t.Run(*tc.mode.Kind, func(t *testing.T) {
					err := func() (err error) {
						getReq := mql.GetDummiesRequest{
							ServiceFilter: &mql.ServiceFilter{
								Id: &mql.ServiceIdFilter{
									Value: &mql.StringFilter{
										In: []string{PaginationA},
									},
								},
							},
							Mode: tc.mode,
						}

						gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
						if err != nil {
							return
						}

						gPagination, ok := gGetRsp.Generic(fieldnames.Pagination)
						assert.True(t, ok)

						var p mql.Pagination
						gPagination.MustToStruct(&p)

						for _, sp := range append(p.Previous, append(p.Current, p.Next...)...) {
							assert.NotNil(t, sp.Id)
						}

						return
					}()
					if err != nil {
						t.Error(err)
					}
				})
			}
		})

		t.Run("mergesPaginations", func(t *testing.T) {
			t.Parallel()

			tcs := []struct {
				mode *mql.GetMode
			}{
				{
					mode: &mql.GetMode{
						Kind:       &mql.GetModeKind.Collection,
						Collection: &mql.CollectionGetMode{},
					},
				},
				{
					mode: &mql.GetMode{
						Kind:   &mql.GetModeKind.Search,
						Search: &mql.SearchGetMode{},
					},
				},
			}

			for _, tc := range tcs {
				t.Run(*tc.mode.Kind, func(t *testing.T) {
					err := func() (err error) {
						getReq := mql.GetDummiesRequest{
							ServiceFilter: &mql.ServiceFilter{
								Id: &mql.ServiceIdFilter{
									Value: &mql.StringFilter{
										In: []string{PaginationA, PaginationB},
									},
								},
							},
							Mode: tc.mode,
						}

						gGetReq := f.MustFromStruct(getReq)

						gGetRsp, err := h(ctx, gGetReq)
						if err != nil {
							return
						}

						gPagination, ok := gGetRsp.Generic(fieldnames.Pagination)
						assert.True(t, ok)

						var p mql.Pagination
						gPagination.MustToStruct(&p)

						assert.Len(t, p.Previous, 0, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))
						assert.Len(t, p.Current, 2, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))
						assert.Len(t, p.Next, 2, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))

						return
					}()
					if err != nil {
						t.Error(err)
					}
				})
			}
		})

		t.Run("distributesPages", func(t *testing.T) {
			tcs := []struct {
				mode *mql.GetMode
			}{
				{
					mode: &mql.GetMode{
						Kind:       &mql.GetModeKind.Collection,
						Collection: &mql.CollectionGetMode{},
					},
				},
				{
					mode: &mql.GetMode{
						Kind:   &mql.GetModeKind.Search,
						Search: &mql.SearchGetMode{},
					},
				},
			}

			for _, tc := range tcs {
				t.Run(*tc.mode.Kind, func(t *testing.T) {
					err := func() (err error) {
						getReq := mql.GetDummiesRequest{
							ServiceFilter: &mql.ServiceFilter{
								Id: &mql.ServiceIdFilter{
									Value: &mql.StringFilter{
										In: []string{PaginationA, PaginationB},
									},
								},
							},
							Mode: tc.mode,
							Pages: []mql.ServicePage{
								{
									Id: &mql.ServiceId{
										ServiceName: mql.String("discovery"),
										Value:       mql.String(PaginationA),
									},
									Page: &mql.Page{
										Kind: &mql.PageKind.IndexPage,
										IndexPage: &mql.IndexPage{
											Value: mql.Int32(0),
										},
									},
								},
								{
									Id: &mql.ServiceId{
										ServiceName: mql.String("discovery"),
										Value:       mql.String(PaginationB),
									},
									Page: &mql.Page{
										Kind: &mql.PageKind.IndexPage,
										IndexPage: &mql.IndexPage{
											Value: mql.Int32(1),
										},
									},
								},
							},
						}

						gGetReq := f.MustFromStruct(getReq)

						gGetRsp, err := h(ctx, gGetReq)
						if err != nil {
							return
						}

						gPagination, ok := gGetRsp.Generic(fieldnames.Pagination)
						assert.True(t, ok)

						var p mql.Pagination
						gPagination.MustToStruct(&p)

						assert.ElementsMatch(t, p.Current, getReq.Pages, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))

						assert.Len(t, p.Previous, 1, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))
						assert.Len(t, p.Current, 2, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))
						assert.Len(t, p.Next, 2, fmt.Sprintf("%v:\n%v\n%v:\n%v", gGetReq.Type().Name(), gGetReq.Sprint(), gGetRsp.Type().Name(), gGetRsp.Sprint()))

						return
					}()
					if err != nil {
						t.Error(err)
					}
				})
			}
		})
	})
}

func FTestError(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestError"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := mql.GetDummiesRequest{
				ServiceFilter: &mql.ServiceFilter{
					Id: &mql.ServiceIdFilter{
						Value: &mql.StringFilter{
							In: []string{ErrorA, ErrorB},
						},
					},
				},
				Select: &mql.GetDummiesResponseSelect{
					Errors: &mql.ErrorSelect{
						Message: mql.Bool(true),
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}
