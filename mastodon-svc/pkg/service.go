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

func (svc Service) GetDeleteStatusesEndpoint() (sdk.DeleteStatusesEndpoint) {
	return sdk.DeleteStatusesEndpoint{}
}

func (svc Service) DeleteStatuses(ctx context.Context, req sdk.DeleteStatusesRequest) (sdk.DeleteStatusesResponse) {
	c := svc.getClient()

	return deleteStatuses(ctx, c, req)
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
									sdk.PersonRelationName.PersonFavorsStatuses,
									sdk.PersonRelationName.PersonAuthorsStatuses,
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

func (svc Service) GetPostStatusesEndpoint() (sdk.PostStatusesEndpoint) {
	return sdk.PostStatusesEndpoint{
		Filter: &sdk.PostStatusesRequestFilter{},
	}
}

func (svc Service) PostStatuses(ctx context.Context, req sdk.PostStatusesRequest) (rsp sdk.PostStatusesResponse) {
	c := svc.getClient()

	return postStatuses(ctx, c, req)
}

func (svc Service) GetPutStatusesEndpoint() (sdk.PutStatusesEndpoint) {
	return sdk.PutStatusesEndpoint{
		Filter: &sdk.PutStatusesRequestFilter{
			Mode: &sdk.PutModeFilter{
				Kind: &sdk.EnumFilter{
					Is: &sdk.PutModeKind.Relation,
				},
				// todo v1 filter for Id.Kind == "me"
				Relation: &sdk.RelationPutModeFilter{
					Relation: &sdk.StringFilter{
						In: []string{
							sdk.PersonRelationName.PersonFavorsStatuses,
							sdk.StatusRelationName.StatusRebloggedByStatuses,
						},
					},
					Operation: &sdk.EnumFilter{
						In: []string{
							sdk.RelationOperation.Add,
							sdk.RelationOperation.Remove,
						},
					},
					// todo v1 filter for Id.Kind == "serviceId"
				},
			},
		},
	}
}

func (svc Service) PutStatuses(ctx context.Context, req sdk.PutStatusesRequest) (rsp sdk.PutStatusesResponse) {
	c := svc.getClient()

	switch *req.Mode.Kind {
	case sdk.PutModeKind.Relation:
		rsp = putStatusesRelation(ctx, c, req)
	default:
	}

	return
}

func (svc Service) GetGetPeopleEndpoint() (sdk.GetPeopleEndpoint) {
	return sdk.GetPeopleEndpoint{
		Filter: &sdk.GetPeopleRequestFilter{
			Or: []sdk.GetPeopleRequestFilter{
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
									sdk.PersonRelationName.PersonBlocksPeople,
									sdk.PersonRelationName.PersonFollowedByPeople,
									sdk.PersonRelationName.PersonFollowsPeople,
									sdk.PersonRelationName.PersonMutesPeople,
									sdk.PersonRelationName.PersonRequestedToBeFollowedByPeople,
									sdk.StatusRelationName.StatusFavoredByPeople,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (svc Service) GetPeople(ctx context.Context, req sdk.GetPeopleRequest) (rsp sdk.GetPeopleResponse) {
	c := svc.getClient()

	switch *req.Mode.Kind {
	case sdk.GetModeKind.Search:
		rsp = getPeopleSearch(ctx, c, req)
	case sdk.GetModeKind.Id:
		rsp = getPersonId(ctx, c, req)
	case sdk.GetModeKind.Relation:
		rsp = getPeopleRelation(ctx, c, req)
	default:
	}

	return
}

func (svc Service) GetPutPeopleEndpoint() (sdk.PutPeopleEndpoint) {
	return sdk.PutPeopleEndpoint{
		Filter: &sdk.PutPeopleRequestFilter{
			Mode: &sdk.PutModeFilter{
				Kind: &sdk.EnumFilter{
					Is: &sdk.PutModeKind.Relation,
				},
				// todo v1 filter for Id.Kind == "me"
				Relation: &sdk.RelationPutModeFilter{
					Relation: &sdk.StringFilter{
						In: []string{
							sdk.PersonRelationName.PersonBlocksPeople,
							sdk.PersonRelationName.PersonFollowsPeople,
							sdk.PersonRelationName.PersonMutesPeople,
						},
					},
					Operation: &sdk.EnumFilter{
						In: []string{
							sdk.RelationOperation.Add,
							sdk.RelationOperation.Remove,
						},
					},
					// todo v1 filter for Id.Kind == "serviceId"
				},
			},
		},
	}
}

func (svc Service) PutPeople(ctx context.Context, req sdk.PutPeopleRequest) (rsp sdk.PutPeopleResponse) {
	c := svc.getClient()

	switch *req.Mode.Kind {
	case sdk.PutModeKind.Relation:
		rsp = putPeopleRelation(ctx, c, req)
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