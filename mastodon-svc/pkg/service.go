package pkg

import (
	"context"
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamate/gen/v0/sdk"
)

type Service struct {
	opts ServiceOpts
}

type ServiceOpts struct {
	Host         string
	ClientId     string
	ClientSecret string
}

func NewService(opts ServiceOpts) Service {
	return Service{opts: opts}
}

func (svc Service) Name() string {
	return "mastodon-svc"
}

func (svc Service) getClient() (c *mastodon.Client, err error) {
	c = mastodon.NewClient(&mastodon.Config{
		Server:       svc.opts.Host,
		ClientID:     svc.opts.ClientId,
		ClientSecret: svc.opts.ClientSecret,
	})

	err = c.Authenticate(context.Background(), "", "")

	return
}

func (svc Service) GetGetPostsEndpoint() sdk.GetPostsEndpoint {
	return sdk.GetPostsEndpoint{
		Filter: &sdk.GetPostsRequestFilter{
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
									sdk.PostRelationName.PostWasRepliedToByPosts,
									sdk.SocialAccountRelationName.SocialAccountFavorsPosts,
									sdk.SocialAccountRelationName.SocialAccountAuthorsPosts,
									sdk.PostFeedRelationName.PostFeedContainsPosts,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (svc Service) GetPosts(ctx context.Context, req sdk.GetPostsRequest) (rsp sdk.GetPostsResponse) {
	c, err := svc.getClient()
	if err != nil {
		rsp = sdk.GetPostsResponse{
			Errors: []sdk.Error{
				{
					Message: sdk.String(err.Error()),
				},
			},
		}

		return
	}

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Id:
		rsp = getPostId(ctx, c, req)
	case sdk.GetModeKind.Relation:
		rsp = getPostsRelation(ctx, c, req)
	case sdk.GetModeKind.Search:
		rsp = getPostsSearch(ctx, c, req)
	default:
	}

	return
}

func (svc Service) GetGetSocialAccountsEndpoint() sdk.GetSocialAccountsEndpoint {
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
									sdk.PostRelationName.PostFavoredBySocialAccounts,
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
	c, err := svc.getClient()
	if err != nil {
		rsp = sdk.GetSocialAccountsResponse{
			Errors: []sdk.Error{
				{
					Message: sdk.String(err.Error()),
				},
			},
		}

		return
	}

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

func (svc Service) GetGetPostFeedsEndpoint() sdk.GetPostFeedsEndpoint {
	return sdk.GetPostFeedsEndpoint{
		Filter: &sdk.GetPostFeedsRequestFilter{
			Or: []sdk.GetPostFeedsRequestFilter{
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

func (svc Service) GetPostFeeds(ctx context.Context, req sdk.GetPostFeedsRequest) (rsp sdk.GetPostFeedsResponse) {
	switch *req.Mode.Kind {
	case sdk.GetModeKind.Collection:
		rsp = getPostFeedsCollection(ctx, req)
	default:
	}

	return
}
