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

func (svc Service) GetGetStatusesEndpoint() (sdk.GetStatusesEndpoint) {
	return sdk.GetStatusesEndpoint{
		Filter: &sdk.GetStatusesRequestFilter{
			Or: []sdk.GetStatusesRequestFilter{
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
									sdk.SocialAccountRelationName.SocialAccountAuthorsStatuses,
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
	var ss []sdk.Status
	var errs []error

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Id:
		ss, errs = getStatusesId(svc.c, req)
	case sdk.GetModeKind.Relation:
		ss, errs = getStatusesRelation(svc.c, req)
	case sdk.GetModeKind.Search:
		ss, errs = getStatusesSearch(svc.c, req)
	default:
		errs = append(errs, errors.New(fmt.Sprintf("can't handle %v", req)))
	}

	rsp.Meta = &sdk.CollectionMeta{}
	for _, err := range errs {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Value: sdk.String(err.Error()),
			},
		})
	}

	rsp.Statuses = ss

	return
}

func (svc Service) GetGetFeedsEndpoint() (sdk.GetFeedsEndpoint) {
	return sdk.GetFeedsEndpoint{
		Filter: &sdk.GetFeedsRequestFilter{
			Mode: &sdk.GetModeFilter{
				Kind: &sdk.EnumFilter{
					In: []string{sdk.GetModeKind.Collection},
				},
			},
		},
	}
}

func (svc Service) GetFeeds(ctx context.Context, req sdk.GetFeedsRequest) (rsp sdk.GetFeedsResponse) {
	var fs []sdk.Feed
	var errs []error

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Collection:
		fs, errs = getFeedsCollection(svc.c, req)
	default:
		errs = append(errs, errors.New(fmt.Sprintf("can't handle %v", req)))
	}

	rsp.Meta = &sdk.CollectionMeta{}
	for _, err := range errs {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Value: sdk.String(err.Error()),
			},
		})
	}

	rsp.Feeds = fs

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

	rsp.Meta = &sdk.CollectionMeta{}
	for _, err := range errs {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Value: sdk.String(err.Error()),
			},
		})
	}

	rsp.SocialAccounts = as

	return
}
