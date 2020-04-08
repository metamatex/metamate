package boot_test

import (
	"context"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/boot"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/line"
	"github.com/metamatex/metamate/metamate/pkg/v0/config"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

const (
	AuthSvc   = "auth"
	PipeSvc   = "pipe"
	ReqFilter = "request-filter"
	SqlxA     = "sqlx-a"
	SqlxB     = "sqlx-b"
)

func TestBoot(t *testing.T) {
	c := config.DefaultConfig

	c.Log.Internal = types.InternalLogConfig{
		config.SvcRsp: map[string]string{
			sdk.GetStatusesResponseName: "{{ .Ctx.Svc.Url.Value }} : \n{{ .Ctx.GSvcReq.Type.Name }}\n{{ .Ctx.GSvcReq.Sprint }}\n\n{{ .Ctx.GSvcRsp.Type.Name }}\n{{ .Ctx.GSvcRsp.Sprint }}",
			"*": "{{ .Ctx.Svc.Url.Value }} : {{ .Ctx.GSvcRsp.Type.Name }}",
		},
	}

	d, err := boot.NewDependencies(c, types.Version{})
	if err != nil {
		log.Fatal(err)
	}

	d.ResolveLine.(*line.Line).SetLog(false)

	ctx := context.Background()

	suffix := "-" + time.Now().Format("3:04:05PM")
	println(suffix)

	f := func(ctx context.Context, gCliReq generic.Generic) (gCliRsp generic.Generic, err error) {
		gCliReq.MustSetGeneric([]string{fieldnames.Select, fieldnames.Meta}, d.Factory.MustFromStruct(sdk.CollectionMetaSelect{
			Errors: &sdk.ErrorSelect{
				Message: &sdk.TextSelect{
					Value: sdk.Bool(true),
				},
			},
		}))

		gCliRsp = d.ServeFunc(ctx, gCliReq)

		return
	}

	//spec.TestRequestFilter(t, ctx, d.Factory, f)
	//
	//spec.TestPipe(t, ctx, d.Factory, f)

	//spec.TestDiscovery(t, ctx, d.Factory, f)

	//FTestHackernewsGetFeedContainsStatuses(t, ctx, d.Factory, f)

	FTestHackernews(t, ctx, d.Factory, f)

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

func FTestHackernews(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	FTestHackernewsSocialAccount(t, ctx, f, h)
	FTestHackernewsGetFeedContainsStatuses(t, ctx, f, h)
	FTestHackernewsGetStatusesSearch(t, ctx, f, h)
	FTestHackernewsGetSocialAccountAuthorsStatuses(t, ctx, f, h)
	FTestHackernewsGetFeedsCollection(t, ctx, f, h)
}

func FTestHackernewsSocialAccount(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "FTestHackernews"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetSocialAccountsRequest{
				ServiceFilter: &sdk.ServiceFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Is: sdk.String("hackernews"),
						},
					},
				},
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Id,
					Id: &sdk.Id{
						Kind: &sdk.IdKind.ServiceId,
						ServiceId: &sdk.ServiceId{
							Value:       sdk.String("21stio"),
							ServiceName: sdk.String("hackernews"),
						},
					},
				},
				Select: &sdk.GetSocialAccountsResponseSelect{
					SocialAccounts: &sdk.SocialAccountSelect{
						All: sdk.Bool(true),
						Relations: &sdk.SocialAccountRelationsSelect{
							AuthorsStatuses: &sdk.StatusesCollectionSelect{
								Meta: &sdk.CollectionMetaSelect{
									Errors: &sdk.ErrorSelect{
										Message: &sdk.TextSelect{
											Value: sdk.Bool(true),
										},
									},
								},
								Statuses: &sdk.StatusSelect{
									Id: &sdk.ServiceIdSelect{
										Value:       sdk.Bool(true),
										ServiceName: sdk.Bool(true),
									},
								},
							},
						},
					},
				},
				Relations: &sdk.GetSocialAccountsRelations{
					AuthorsStatuses: &sdk.GetStatusesCollection{
						Select: &sdk.StatusesCollectionSelect{
							Meta: &sdk.CollectionMetaSelect{
								Errors: &sdk.ErrorSelect{
									Message: &sdk.TextSelect{
										Value: sdk.Bool(true),
									},
								},
							},
							Statuses: &sdk.StatusSelect{
								Id: &sdk.ServiceIdSelect{
									Value:       sdk.Bool(true),
									ServiceName: sdk.Bool(true),
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

func FTestHackernewsGetFeedContainsStatuses(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "FTestHackernewsGetFeedContainsStatuses"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetStatusesRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Relation,
					Relation: &sdk.RelationGetMode{
						Id: &sdk.ServiceId{
							ServiceName: sdk.String("hackernews"),
							Value: sdk.String("topstories"),
						},
						Relation: &sdk.FeedRelationName.FeedContainsStatuses,
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := sdk.GetStatusesResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.True(t, len(getRsp.Statuses) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetStatusesSearch(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "FTestHackernewsGetStatusesSearch"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetStatusesRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Search,
					Search: &sdk.SearchGetMode{
						Term: sdk.String("graphql"),
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := sdk.GetStatusesResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.True(t, len(getRsp.Statuses) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetSocialAccountAuthorsStatuses(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "FTestHackernewsGetSocialAccountAuthorsStatuses"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetStatusesRequest{
				Mode: &sdk.GetMode{
					Kind: &sdk.GetModeKind.Relation,
					Relation: &sdk.RelationGetMode{
						Id: &sdk.ServiceId{
							ServiceName: sdk.String("hackernews"),
							Value: sdk.String("21stio"),
						},
						Relation: &sdk.SocialAccountRelationName.SocialAccountAuthorsStatuses,
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := sdk.GetStatusesResponse{}
			gGetRsp.MustToStruct(&getRsp)

			assert.True(t, len(getRsp.Statuses) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestHackernewsGetFeedsCollection(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error)) {
	name := "FTestHackernewsGetFeedsCollection"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		err := func() (err error) {
			getReq := sdk.GetFeedsRequest{
				Filter: &sdk.FeedFilter{
					Id: &sdk.ServiceIdFilter{
						Value: &sdk.StringFilter{
							Contains: sdk.String("top"),
						},
					},
				},
			}

			gGetRsp, err := h(ctx, f.MustFromStruct(getReq))
			if err != nil {
				return
			}

			gGetRsp.Print()

			getRsp := sdk.GetFeedsResponse{}
			gGetRsp.MustToStruct(&getRsp)
			assert.True(t, len(getRsp.Feeds) != 0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}
