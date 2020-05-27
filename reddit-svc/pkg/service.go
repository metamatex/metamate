package pkg

import (
	"context"
	"fmt"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/reddit-svc/pkg/business"
	"github.com/metamatex/metamate/reddit-svc/pkg/communication"
	"github.com/metamatex/metamate/reddit-svc/pkg/types"
	"net/http"
)

type Service struct {
	c    communication.Client
	opts ServiceOpts
}

type ServiceOpts struct {
	Client      *http.Client
	Credentials types.Credentials
	UserAgent   string
}

func NewService(opts ServiceOpts) (svc Service, err error) {
	c, err := communication.NewClient(communication.ClientOpts{
		Credentials: opts.Credentials,
		Client:      opts.Client,
		UserAgent:   opts.UserAgent,
	})
	if err != nil {
		return
	}

	err = c.Authenticate()
	if err != nil {
		return
	}

	svc = Service{opts: opts, c: c}

	return
}

func (svc Service) Name() string {
	return "reddit"
}

func (svc Service) GetGetPostsEndpoint() mql.GetPostsEndpoint {
	return mql.GetPostsEndpoint{
		Filter: &mql.GetPostsRequestFilter{
			Or: []mql.GetPostsRequestFilter{
				{
					Mode: &mql.GetModeFilter{
						Kind: &mql.EnumFilter{
							In: []string{mql.GetModeKind.Id},
						},
						Id: &mql.IdFilter{
							Kind: &mql.EnumFilter{
								In: []string{mql.IdKind.ServiceId, mql.IdKind.Name},
							},
						},
					},
				},
				{
					Mode: &mql.GetModeFilter{
						Kind: &mql.EnumFilter{
							In: []string{mql.GetModeKind.Relation},
						},
						Relation: &mql.RelationGetModeFilter{
							Id: &mql.IdFilter{
								Kind: &mql.EnumFilter{
									In: []string{mql.IdKind.ServiceId, mql.IdKind.Name},
								},
							},
							Relation: &mql.StringFilter{
								In: []string{
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
	var fs []mql.Post
	var errs []mql.Error
	var pagination mql.Pagination

	switch *req.Mode.Kind {
	case mql.GetModeKind.Relation:
		fs, pagination, errs = business.GetPostRelation(svc.c, req)
	default:
		errs = append(errs, mql.Error{
			Message: mql.String(fmt.Sprintf("can't handle %v", *req.Mode.Kind)),
		})
	}

	rsp.Posts = fs
	rsp.Errors = errs
	rsp.Pagination = &pagination

	return
}

func (svc Service) GetGetPostFeedsEndpoint() mql.GetPostFeedsEndpoint {
	return mql.GetPostFeedsEndpoint{
		Filter: &mql.GetPostFeedsRequestFilter{
			Or: []mql.GetPostFeedsRequestFilter{
				{
					Mode: &mql.GetModeFilter{
						Kind: &mql.EnumFilter{
							In: []string{mql.GetModeKind.Id},
						},
						Id: &mql.IdFilter{
							Kind: &mql.EnumFilter{
								In: []string{mql.IdKind.ServiceId, mql.IdKind.Name},
							},
						},
					},
				},
			},
		},
	}
}

func (svc Service) GetPostFeeds(ctx context.Context, req mql.GetPostFeedsRequest) (rsp mql.GetPostFeedsResponse) {
	var fs []mql.PostFeed
	var errs []mql.Error

	switch *req.Mode.Kind {
	case mql.GetModeKind.Id:
		fs, errs = business.GetPostsFeedId(svc.c, req)
	default:
		errs = append(errs, mql.Error{
			Message: mql.String(fmt.Sprintf("can't handle %v", req)),
		})
	}

	rsp.PostFeeds = fs
	rsp.Errors = errs

	return
}

func (svc Service) GetGetSocialAccountsEndpoint() mql.GetSocialAccountsEndpoint {
	return mql.GetSocialAccountsEndpoint{
		Filter: &mql.GetSocialAccountsRequestFilter{
			Or: []mql.GetSocialAccountsRequestFilter{
				{
					Mode: &mql.GetModeFilter{
						Kind: &mql.EnumFilter{
							In: []string{mql.GetModeKind.Id},
						},
						Id: &mql.IdFilter{
							Kind: &mql.EnumFilter{
								In: []string{mql.IdKind.ServiceId, mql.IdKind.Username},
							},
						},
					},
				},
			},
		},
	}
}

func (svc Service) GetSocialAccounts(ctx context.Context, req mql.GetSocialAccountsRequest) (rsp mql.GetSocialAccountsResponse) {
	var fs []mql.SocialAccount
	var errs []mql.Error

	switch *req.Mode.Kind {
	case mql.GetModeKind.Id:
		fs, errs = business.GetSocialAccountId(svc.c, req)
	default:
		errs = append(errs, mql.Error{
			Message: mql.String(fmt.Sprintf("can't handle %v", req)),
		})
	}

	rsp.SocialAccounts = fs
	rsp.Errors = errs

	return
}
