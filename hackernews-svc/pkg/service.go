package pkg

import (
	"context"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"
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

func (svc Service) GetGetPostsEndpoint() sdk.GetPostsEndpoint {
	return sdk.GetPostsEndpoint{
		Filter: &sdk.GetPostsRequestFilter{
			Or: []sdk.GetPostsRequestFilter{
				{
					Mode: &sdk.GetModeFilter{
						Kind: &sdk.EnumFilter{
							In: []string{sdk.GetModeKind.Id},
						},
						Id: &sdk.IdFilter{
							Kind: &sdk.EnumFilter{
								Is: &sdk.IdKind.ServiceId,
							},
						},
					},
				},
				{
					Mode: &sdk.GetModeFilter{
						Kind: &sdk.EnumFilter{
							In: []string{sdk.GetModeKind.Relation},
						},
						Id: &sdk.IdFilter{
							Kind: &sdk.EnumFilter{
								Is: &sdk.IdKind.ServiceId,
							},
						},
						Relation: &sdk.RelationGetModeFilter{
							Relation: &sdk.StringFilter{
								In: []string{
									sdk.SocialAccountRelationName.SocialAccountAuthorsPosts,
									sdk.SocialAccountRelationName.SocialAccountBookmarksPosts,
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
	var ps []sdk.Post
	var pagination *sdk.Pagination
	var errs []sdk.Error

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Id:
		ps, errs = firebase.GetPostsId(svc.c, req)
	case sdk.GetModeKind.Relation:
		switch *req.Mode.Relation.Relation {
		case sdk.SocialAccountRelationName.SocialAccountAuthorsPosts:
			ps, errs = firebase.GetSocialAccountAuthorsPosts(svc.c, req)
		case sdk.SocialAccountRelationName.SocialAccountBookmarksPosts:
			ps, pagination, errs = website.GetSocialAccountBookmarksPosts(svc.c, *req.Mode.Relation.Id.Value, nil)
		case sdk.PostFeedRelationName.PostFeedContainsPosts:
			ps, errs = firebase.GetPostFeedContainsPosts(svc.c, *req.Mode.Relation.Id.Value)
		default:
			errs = append(errs, sdk.Error{
				Message: sdk.String(fmt.Sprintf("can't handle relation %v", *req.Mode.Relation.Relation)),
			})
		}
	case sdk.GetModeKind.Search:
		ps, errs = angolia.GetPostsSearch(svc.c, req)
	default:
		errs = append(errs, sdk.Error{
			Message: sdk.String(fmt.Sprintf("can't handle mode %v", req.Mode.Kind)),
		})
	}

	rsp.Posts = ps
	rsp.Pagination = pagination
	rsp.Errors = errs

	return
}

func (svc Service) GetGetPostFeedsEndpoint() sdk.GetPostFeedsEndpoint {
	return sdk.GetPostFeedsEndpoint{
		Filter: &sdk.GetPostFeedsRequestFilter{
			Mode: &sdk.GetModeFilter{
				Kind: &sdk.EnumFilter{
					In: []string{sdk.GetModeKind.Collection},
				},
			},
		},
	}
}

func (svc Service) GetPostFeeds(ctx context.Context, req sdk.GetPostFeedsRequest) (rsp sdk.GetPostFeedsResponse) {
	var fs []sdk.PostFeed
	var errs []sdk.Error

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Collection:
		fs, errs = static.GetPostFeedsCollection()
	default:
		errs = append(errs, sdk.Error{
			Message: sdk.String(fmt.Sprintf("can't handle %v", req)),
		})
	}

	rsp.PostFeeds = fs
	rsp.Errors = errs

	return
}

func (svc Service) GetGetSocialAccountsEndpoint() sdk.GetSocialAccountsEndpoint {
	return sdk.GetSocialAccountsEndpoint{
		Filter: &sdk.GetSocialAccountsRequestFilter{
			Or: []sdk.GetSocialAccountsRequestFilter{
				{
					Mode: &sdk.GetModeFilter{
						Kind: &sdk.EnumFilter{
							In: []string{sdk.GetModeKind.Id},
						},
						Id: &sdk.IdFilter{
							Kind: &sdk.EnumFilter{
								In: []string{
									sdk.IdKind.ServiceId,
									sdk.IdKind.Username,
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
	var as []sdk.SocialAccount
	var errs []sdk.Error

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Id:
		as, errs = firebase.GetSocialAccountId(svc.c, req)
	default:
		errs = append(errs, sdk.Error{
			Message: sdk.String(fmt.Sprintf("can't handle %v", req)),
		})
	}

	rsp.SocialAccounts = as
	rsp.Errors = errs

	return
}
