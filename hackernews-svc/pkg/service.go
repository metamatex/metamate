package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"
	"net/http"
)

type Service struct {
	c *http.Client
}

func NewService(c *http.Client) (svc Service) {
	return Service{c: c}
}

func (svc Service) Name() (string) {
	return "hackernews"
}

func (svc Service) GetGetPostsEndpoint() (sdk.GetPostsEndpoint) {
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
	var ss []sdk.Post
	var errs []error

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Id:
		ss, errs = getPostsId(svc.c, req)
	case sdk.GetModeKind.Relation:
		ss, errs = getPostsRelation(svc.c, req)
	case sdk.GetModeKind.Search:
		ss, errs = getPostsSearch(svc.c, req)
	default:
		errs = append(errs, errors.New(fmt.Sprintf("can't handle %v", req)))
	}

	for _, err := range errs {
		rsp.Errors = append(rsp.Errors, sdk.Error{
			Message: sdk.String(err.Error()),
		})
	}

	rsp.Posts = ss

	return
}

func (svc Service) GetGetPostFeedsEndpoint() (sdk.GetPostFeedsEndpoint) {
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
	var errs []error

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Collection:
		fs, errs = getPostFeedsCollection(svc.c, req)
	default:
		errs = append(errs, errors.New(fmt.Sprintf("can't handle %v", req)))
	}

	for _, err := range errs {
		rsp.Errors = append(rsp.Errors, sdk.Error{
			Message: sdk.String(err.Error()),
		})
	}

	rsp.PostFeeds = fs

	return
}

func (svc Service) GetGetSocialAccountsEndpoint() (sdk.GetSocialAccountsEndpoint) {
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
	var errs []error

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Id:
		as, errs = getSocialAccountId(svc.c, req)
	default:
		errs = append(errs, errors.New(fmt.Sprintf("can't handle %v", req)))
	}

	for _, err := range errs {
		rsp.Errors = append(rsp.Errors, sdk.Error{
			Message: sdk.String(err.Error()),
		})
	}

	rsp.SocialAccounts = as

	return
}
