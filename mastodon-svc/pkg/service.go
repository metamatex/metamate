package pkg

import (
	"context"
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamate/gen/v0/mql"
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

func (svc Service) GetGetPostsEndpoint() mql.GetPostsEndpoint {
	return mql.GetPostsEndpoint{
		Filter: &mql.GetPostsRequestFilter{
			Mode: &mql.GetModeFilter{
				Or: []mql.GetModeFilter{
					{
						Kind: &mql.EnumFilter{
							Is: &mql.GetModeKind.Id,
						},
						Id: &mql.IdFilter{
							Kind: &mql.EnumFilter{
								Is: &mql.IdKind.ServiceId,
							},
						},
					},
					{
						Kind: &mql.EnumFilter{
							Is: &mql.GetModeKind.Search,
						},
					},
					{
						Kind: &mql.EnumFilter{
							Is: &mql.GetModeKind.Relation,
						},
						Relation: &mql.RelationGetModeFilter{
							Relation: &mql.StringFilter{
								In: []string{
									mql.PostRelationName.PostWasRepliedToByPosts,
									mql.SocialAccountRelationName.SocialAccountFavorsPosts,
									mql.SocialAccountRelationName.SocialAccountAuthorsPosts,
									mql.PostFeedRelationName.PostFeedContainsPosts,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (svc Service) GetPosts(ctx context.Context, req mql.GetPostsRequest) (rsp mql.GetPostsResponse) {
	c, err := svc.getClient()
	if err != nil {
		rsp = mql.GetPostsResponse{
			Errors: []mql.Error{
				{
					Message: mql.String(err.Error()),
				},
			},
		}

		return
	}

	switch *req.Mode.Kind {
	case mql.GetModeKind.Id:
		rsp = getPostId(ctx, c, req)
	case mql.GetModeKind.Relation:
		rsp = getPostsRelation(ctx, c, req)
	case mql.GetModeKind.Search:
		rsp = getPostsSearch(ctx, c, req)
	default:
	}

	return
}

func (svc Service) GetGetSocialAccountsEndpoint() mql.GetSocialAccountsEndpoint {
	return mql.GetSocialAccountsEndpoint{
		Filter: &mql.GetSocialAccountsRequestFilter{
			Or: []mql.GetSocialAccountsRequestFilter{
				{
					Mode: &mql.GetModeFilter{
						Kind: &mql.EnumFilter{
							Is: &mql.GetModeKind.Id,
						},
						Id: &mql.IdFilter{
							Kind: &mql.EnumFilter{
								In: []string{
									mql.IdKind.ServiceId,
									mql.IdKind.Me,
								},
							},
						},
					},
				},
				{
					Mode: &mql.GetModeFilter{
						Kind: &mql.EnumFilter{
							Is: &mql.GetModeKind.Search,
						},
					},
				},
				{
					Mode: &mql.GetModeFilter{
						Kind: &mql.EnumFilter{
							Is: &mql.GetModeKind.Relation,
						},
						Relation: &mql.RelationGetModeFilter{
							Relation: &mql.StringFilter{
								In: []string{
									mql.SocialAccountRelationName.SocialAccountBlocksSocialAccounts,
									mql.SocialAccountRelationName.SocialAccountFollowedBySocialAccounts,
									mql.SocialAccountRelationName.SocialAccountFollowsSocialAccounts,
									mql.SocialAccountRelationName.SocialAccountMutesSocialAccounts,
									mql.SocialAccountRelationName.SocialAccountRequestedToBeFollowedBySocialAccounts,
									mql.PostRelationName.PostFavoredBySocialAccounts,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (svc Service) GetSocialAccounts(ctx context.Context, req mql.GetSocialAccountsRequest) (rsp mql.GetSocialAccountsResponse) {
	c, err := svc.getClient()
	if err != nil {
		rsp = mql.GetSocialAccountsResponse{
			Errors: []mql.Error{
				{
					Message: mql.String(err.Error()),
				},
			},
		}

		return
	}

	switch *req.Mode.Kind {
	case mql.GetModeKind.Search:
		rsp = getSocialAccountsSearch(ctx, c, req)
	case mql.GetModeKind.Id:
		rsp = getSocialAccountId(ctx, c, req)
	case mql.GetModeKind.Relation:
		rsp = getSocialAccountsRelation(ctx, c, req)
	default:
	}

	return
}

func (svc Service) GetGetPostFeedsEndpoint() mql.GetPostFeedsEndpoint {
	return mql.GetPostFeedsEndpoint{
		Filter: &mql.GetPostFeedsRequestFilter{
			Or: []mql.GetPostFeedsRequestFilter{
				{
					Mode: &mql.GetModeFilter{
						Kind: &mql.EnumFilter{
							Is: &mql.GetModeKind.Collection,
						},
					},
				},
			},
		},
	}
}

func (svc Service) GetPostFeeds(ctx context.Context, req mql.GetPostFeedsRequest) (rsp mql.GetPostFeedsResponse) {
	switch *req.Mode.Kind {
	case mql.GetModeKind.Collection:
		rsp = getPostFeedsCollection(ctx, req)
	default:
	}

	return
}
