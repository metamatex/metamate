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

	FTestPagination(t, ctx, d.Factory, f)

	FTestHackernews(t, ctx, d.Factory, f)

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

func FTestHackernews(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	t.Run("GetSocialAccounts", func(t *testing.T) {
		t.Run("id", func(t *testing.T) {
			err := func() (err error) {
				getReq := mql.GetSocialAccountsRequest{
					Mode: &mql.GetMode{
						Kind: &mql.GetModeKind.Id,
						Id: &mql.Id{
							Kind: &mql.IdKind.ServiceId,
							ServiceId: &mql.ServiceId{
								Value:       mql.String("21stio"),
								ServiceName: mql.String("hackernews"),
							},
						},
					},
				}

				gGetReq := f.MustFromStruct(getReq)

				gGetRsp, err := h(ctx, gGetReq)
				if err != nil {
					return
				}

				var getRsp mql.GetSocialAccountsResponse
				gGetRsp.MustToStruct(&getRsp)

				assert.Len(t, getRsp.SocialAccounts, 1)

				return
			}()
			if err != nil {
				t.Error(err)
			}
		})
	})

	//FTestHackernewsSocialAccount(t, ctx, f, h)
	//FTestHackernewsGetPostFeedContainsPosts(t, ctx, f, h)
	//FTestHackernewsGetPostsSearch(t, ctx, f, h)
	//FTestHackernewsGetSocialAccountBookmarksPosts(t, ctx, f, h)
	//FTestHackernewsGetSocialAccountAuthorsPosts(t, ctx, f, h)
	//FTestHackernewsGetPostFeedsCollection(t, ctx, f, h)
}

func FTestHackernewsSocialAccount(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernews"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := mql.GetSocialAccountsRequest{
				ServiceFilter: &mql.ServiceFilter{
					Id: &mql.ServiceIdFilter{
						Value: &mql.StringFilter{
							Is: mql.String("hackernews"),
						},
					},
				},
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Id,
					Id: &mql.Id{
						Kind: &mql.IdKind.ServiceId,
						ServiceId: &mql.ServiceId{
							Value:       mql.String("21stio"),
							ServiceName: mql.String("hackernews"),
						},
					},
				},
				Select: &mql.GetSocialAccountsResponseSelect{
					SocialAccounts: &mql.SocialAccountSelect{
						All: mql.Bool(true),
						Relations: &mql.SocialAccountRelationsSelect{
							AuthorsPosts: &mql.PostsCollectionSelect{
								Errors: &mql.ErrorSelect{
									Message: mql.Bool(true),
								},
								Posts: &mql.PostSelect{
									Id: &mql.ServiceIdSelect{
										Value:       mql.Bool(true),
										ServiceName: mql.Bool(true),
									},
								},
							},
						},
					},
				},
				Relations: &mql.GetSocialAccountsRelations{
					AuthorsPosts: &mql.GetPostsCollection{
						Select: &mql.PostsCollectionSelect{
							Errors: &mql.ErrorSelect{
								Message: mql.Bool(true),
							},
							Posts: &mql.PostSelect{
								Id: &mql.ServiceIdSelect{
									Value:       mql.Bool(true),
									ServiceName: mql.Bool(true),
								},
							},
						},
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

func FTestHackernewsGetPostFeedContainsPosts(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernewsGetPostFeedContainsPosts"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := mql.GetPostsRequest{
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Relation,
					Relation: &mql.RelationGetMode{
						Id: &mql.ServiceId{
							ServiceName: mql.String("hackernews"),
							Value:       mql.String("topstories"),
						},
						Relation: &mql.PostFeedRelationName.PostFeedContainsPosts,
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := mql.GetPostsResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.True(t, len(getRsp.Posts) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetPostsSearch(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernewsGetPostsSearch"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := mql.GetPostsRequest{
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Search,
					Search: &mql.SearchGetMode{
						Term: mql.String("book recommendations"),
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := mql.GetPostsResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.True(t, len(getRsp.Posts) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetSocialAccountAuthorsPosts(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernewsGetSocialAccountAuthorsPosts"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := mql.GetPostsRequest{
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Relation,
					Relation: &mql.RelationGetMode{
						Id: &mql.ServiceId{
							ServiceName: mql.String("hackernews"),
							Value:       mql.String("21stio"),
						},
						Relation: &mql.SocialAccountRelationName.SocialAccountAuthorsPosts,
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := mql.GetPostsResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.True(t, len(getRsp.Posts) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetPostFeedsCollection(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernewsGetPostFeedsCollection"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := mql.GetPostFeedsRequest{
				Filter: &mql.PostFeedFilter{
					Id: &mql.ServiceIdFilter{
						Value: &mql.StringFilter{
							Contains: mql.String("top"),
						},
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := mql.GetPostFeedsResponse{}
			gGetRsp.MustToStruct(&getRsp)
			assert.True(t, len(getRsp.PostFeeds) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetSocialAccountBookmarksPosts(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "TestHackernewsGetSocialAccountBookmarksPosts"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := mql.GetPostsRequest{
				Mode: &mql.GetMode{
					Kind: &mql.GetModeKind.Relation,
					Relation: &mql.RelationGetMode{
						Id: &mql.ServiceId{
							ServiceName: mql.String("hackernews"),
							Value:       mql.String("peter_d_sherman"),
						},
						Relation: &mql.SocialAccountRelationName.SocialAccountBookmarksPosts,
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			//gGetRsp.Print()

			getRsp := mql.GetPostsResponse{}
			gGetRsp.MustToStruct(&getRsp)

			mql.Print(getRsp.Pagination)

			assert.True(t, len(getRsp.Posts) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}
