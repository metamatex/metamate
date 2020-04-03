package pkg

import (
	"context"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var svc = NewService(&http.Client{})
var ctx = context.Background()

func TestService_GetFeedsCollection(t *testing.T) {
	rsp := svc.GetFeeds(ctx, sdk.GetFeedsRequest{
		Mode: &sdk.GetMode{
			Kind: &sdk.GetModeKind.Collection,
		},
	})

	assert.True(t, len(rsp.Feeds) != 0)
	assert.True(t, rsp.Meta == nil || len(rsp.Meta.Errors) == 0)
}

func TestService_GetSocialAccountsId(t *testing.T) {
	reqs := []sdk.GetSocialAccountsRequest{
		{
			Mode: &sdk.GetMode{
				Kind: &sdk.GetModeKind.Id,
				Id: &sdk.Id{
					Kind: &sdk.IdKind.ServiceId,
					ServiceId: &sdk.ServiceId{
						Value: sdk.String("21stio"),
					},
				},
			},
		},
		{
			Mode: &sdk.GetMode{
				Kind: &sdk.GetModeKind.Id,
				Id: &sdk.Id{
					Kind: &sdk.IdKind.Username,
					Username: sdk.String("21stio"),
				},
			},
		},
	}

	for _, req := range reqs {
		rsp := svc.GetSocialAccounts(ctx, req)

		assert.True(t, len(rsp.SocialAccounts) == 1)
		assert.True(t, rsp.Meta == nil || len(rsp.Meta.Errors) == 0)
	}
}

func TestService_GetStatusesId(t *testing.T) {
	reqs := []sdk.GetStatusesRequest{
		{
			Mode: &sdk.GetMode{
				Kind: &sdk.GetModeKind.Id,
				Id: &sdk.Id{
					Kind: &sdk.IdKind.ServiceId,
					ServiceId: &sdk.ServiceId{
						Value: sdk.String("20878891"),
					},
				},
			},
		},
	}

	for _, req := range reqs {
		rsp := svc.GetStatuses(ctx, req)

		utils.Print(rsp)
	}
}

func TestService_GetStatusesRelation(t *testing.T) {
	reqs := []sdk.GetStatusesRequest{
		{
			Mode: &sdk.GetMode{
				Kind: &sdk.GetModeKind.Relation,
				Relation: &sdk.RelationGetMode{
					Id: &sdk.ServiceId{
						Value: sdk.String("21stio"),
					},
					Relation: &sdk.SocialAccountRelationName.SocialAccountAuthorsStatuses,
				},
			},
		},
	}

	for _, req := range reqs {
		rsp := svc.GetStatuses(ctx, req)

		utils.Print(rsp)
	}
}

func TestService_GetStatusesSearch(t *testing.T) {
	reqs := []sdk.GetStatusesRequest{
		{
			Mode: &sdk.GetMode{
				Kind: &sdk.GetModeKind.Search,
				Search: &sdk.SearchGetMode{
					Term: sdk.String("21stio"),
				},
			},
		},
	}

	for _, req := range reqs {
		rsp := svc.GetStatuses(ctx, req)

		utils.Print(rsp)
	}
}


