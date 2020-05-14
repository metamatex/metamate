package pkg

import (
	"context"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/mql"
	"github.com/metamatex/metamate/hackernews-svc/pkg/persistence/angolia"
	"github.com/metamatex/metamate/hackernews-svc/pkg/persistence/firebase"
	"github.com/metamatex/metamate/hackernews-svc/pkg/persistence/static"
	"github.com/metamatex/metamate/hackernews-svc/pkg/persistence/website"
	"net/http"
)

type Service struct {
	c *http.Client
}

func NewService(c *http.Client) (svc Service) {
	return Service{c: c}
}

func (svc Service) Name() string {
	return "hackernews"
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
								Is: &mql.IdKind.ServiceId,
							},
						},
					},
				},
				{
					Mode: &mql.GetModeFilter{
						Kind: &mql.EnumFilter{
							In: []string{mql.GetModeKind.Relation},
						},
						Id: &mql.IdFilter{
							Kind: &mql.EnumFilter{
								Is: &mql.IdKind.ServiceId,
							},
						},
						Relation: &mql.RelationGetModeFilter{
							Relation: &mql.StringFilter{
								In: []string{
									mql.SocialAccountRelationName.SocialAccountAuthorsPosts,
									mql.SocialAccountRelationName.SocialAccountBookmarksPosts,
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
	var ps []mql.Post
	var pagination *mql.Pagination
	var errs []mql.Error

	switch *req.Mode.Kind {
	case mql.GetModeKind.Id:
		ps, errs = firebase.GetPostsId(svc.c, req)
	case mql.GetModeKind.Relation:
		switch *req.Mode.Relation.Relation {
		case mql.SocialAccountRelationName.SocialAccountAuthorsPosts:
			ps, errs = firebase.GetSocialAccountAuthorsPosts(svc.c, req)
		case mql.SocialAccountRelationName.SocialAccountBookmarksPosts:
			ps, pagination, errs = website.GetSocialAccountBookmarksPosts(svc.c, *req.Mode.Relation.Id.Value, nil)
		case mql.PostFeedRelationName.PostFeedContainsPosts:
			ps, errs = firebase.GetPostFeedContainsPosts(svc.c, *req.Mode.Relation.Id.Value)
		default:
			errs = append(errs, mql.Error{
				Message: mql.String(fmt.Sprintf("can't handle relation %v", *req.Mode.Relation.Relation)),
			})
		}
	case mql.GetModeKind.Search:
		ps, errs, pagination = angolia.GetPostsSearch(svc.c, req)
	default:
		errs = append(errs, mql.Error{
			Message: mql.String(fmt.Sprintf("can't handle mode %v", req.Mode.Kind)),
		})
	}

	rsp.Posts = ps
	rsp.Pagination = pagination
	rsp.Errors = errs

	return
}

func (svc Service) GetGetPostFeedsEndpoint() mql.GetPostFeedsEndpoint {
	return mql.GetPostFeedsEndpoint{
		Filter: &mql.GetPostFeedsRequestFilter{
			Mode: &mql.GetModeFilter{
				Kind: &mql.EnumFilter{
					In: []string{mql.GetModeKind.Collection},
				},
			},
		},
	}
}

func (svc Service) GetPostFeeds(ctx context.Context, req mql.GetPostFeedsRequest) (rsp mql.GetPostFeedsResponse) {
	var fs []mql.PostFeed
	var errs []mql.Error

	switch *req.Mode.Kind {
	case mql.GetModeKind.Collection:
		fs, errs = static.GetPostFeedsCollection()
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
								In: []string{
									mql.IdKind.ServiceId,
									mql.IdKind.Username,
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
	var as []mql.SocialAccount
	var errs []mql.Error

	switch *req.Mode.Kind {
	case mql.GetModeKind.Id:
		as, errs = firebase.GetSocialAccountId(svc.c, req)
	default:
		errs = append(errs, mql.Error{
			Message: mql.String(fmt.Sprintf("can't handle %v", req)),
		})
	}

	rsp.SocialAccounts = as
	rsp.Errors = errs

	return
}
