package pkg

import (
	"context"
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"log"
)

type Service struct{
	opts ServiceOpts
}

type ServiceOpts struct {
	Host string
	ClientId string
	ClientSecret string
}

func NewService(opts ServiceOpts) Service {
	return Service{opts: opts}
}

func (svc Service) Name() (string) {
	return "mastodon-svc"
}

func (svc Service) getClient() (c *mastodon.Client) {
	c = mastodon.NewClient(&mastodon.Config{
		Server:       svc.opts.Host,
		ClientID:     svc.opts.ClientId,
		ClientSecret: svc.opts.ClientSecret,
	})

	err := c.Authenticate(context.Background(), "ph.woerdehoff@gmail.com", "]9T)8VB6Em")
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func (svc Service) GetGetStatusesEndpoint() (sdk.GetStatusesEndpoint) {
	return sdk.GetStatusesEndpoint{
		Filter: &sdk.GetStatusesRequestFilter{
			Mode: &sdk.GetModeFilter{
				Or: []sdk.GetModeFilter{
					{
						Kind: &sdk.EnumFilter{
							Is: &sdk.GetModeKind.Id,
						},
						Id: &sdk.IdFilter{
							Kind: &sdk.EnumFilter{
								Is: &sdk.IdKind.ServiceId,
							},
						},
					},
					{
						Kind: &sdk.EnumFilter{
							Is: &sdk.GetModeKind.Search,
						},
					},
					{
						Kind: &sdk.EnumFilter{
							Is: &sdk.GetModeKind.Relation,
						},
						Relation: &sdk.RelationGetModeFilter{
							Relation: &sdk.StringFilter{
								In: []string{
									sdk.StatusRelationName.StatusWasRepliedToByStatuses,
									sdk.SocialAccountRelationName.SocialAccountFavorsStatuses,
									sdk.SocialAccountRelationName.SocialAccountAuthorsStatuses,
									sdk.FeedRelationName.FeedContainsStatuses,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (svc Service) GetStatuses(ctx context.Context, req sdk.GetStatusesRequest) (rsp sdk.GetStatusesResponse) {
	c := svc.getClient()

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Id:
		rsp = getStatusId(ctx, c, req)
	case sdk.GetModeKind.Relation:
		rsp = getStatusesRelation(ctx, c, req)
	case sdk.GetModeKind.Search:
		rsp = getStatusesSearch(ctx, c, req)
	default:
	}

	return
}

func (svc Service) GetGetSocialAccountsEndpoint() (sdk.GetSocialAccountsEndpoint) {
	return sdk.GetSocialAccountsEndpoint{
		Filter: &sdk.GetSocialAccountsRequestFilter{
			Or: []sdk.GetSocialAccountsRequestFilter{
				{
					Mode: &sdk.GetModeFilter{
						Kind: &sdk.EnumFilter{
							Is: &sdk.GetModeKind.Id,
						},
						Id: &sdk.IdFilter{
							Kind: &sdk.EnumFilter{
								In: []string{
									sdk.IdKind.ServiceId,
									sdk.IdKind.Me,
								},
							},
						},
					},
				},
				{
					Mode: &sdk.GetModeFilter{
						Kind: &sdk.EnumFilter{
							Is: &sdk.GetModeKind.Search,
						},
					},
				},
				{
					Mode: &sdk.GetModeFilter{
						Kind: &sdk.EnumFilter{
							Is: &sdk.GetModeKind.Relation,
						},
						Relation: &sdk.RelationGetModeFilter{
							Relation: &sdk.StringFilter{
								In: []string{
									sdk.SocialAccountRelationName.SocialAccountBlocksSocialAccounts,
									sdk.SocialAccountRelationName.SocialAccountFollowedBySocialAccounts,
									sdk.SocialAccountRelationName.SocialAccountFollowsSocialAccounts,
									sdk.SocialAccountRelationName.SocialAccountMutesSocialAccounts,
									sdk.SocialAccountRelationName.SocialAccountRequestedToBeFollowedBySocialAccounts,
									sdk.StatusRelationName.StatusFavoredBySocialAccounts,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (svc Service) GetSocialAccounts(ctx context.Context, req sdk.GetSocialAccountsRequest) (rsp sdk.GetSocialAccountsResponse) {
	c := svc.getClient()

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Search:
		rsp = getSocialAccountsSearch(ctx, c, req)
	case sdk.GetModeKind.Id:
		rsp = getSocialAccountId(ctx, c, req)
	case sdk.GetModeKind.Relation:
		rsp = getSocialAccountsRelation(ctx, c, req)
	default:
	}

	return
}

func (svc Service) GetGetFeedsEndpoint() (sdk.GetFeedsEndpoint) {
	return sdk.GetFeedsEndpoint{
		Filter: &sdk.GetFeedsRequestFilter{
			Or: []sdk.GetFeedsRequestFilter{
				{
					Mode: &sdk.GetModeFilter{
						Kind: &sdk.EnumFilter{
							Is: &sdk.GetModeKind.Collection,
						},
					},
				},
			},
		},
	}
}

func (svc Service) GetFeeds(ctx context.Context, req sdk.GetFeedsRequest) (rsp sdk.GetFeedsResponse) {
	switch *req.Mode.Kind {
	case sdk.GetModeKind.Collection:
		rsp = getFeedsCollection(ctx, req)
	default:
	}

	return
}