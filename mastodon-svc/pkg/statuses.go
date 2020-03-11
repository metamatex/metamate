package pkg

import (
	"context"
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/utils/ptr"
)

func getStatusId(ctx context.Context, c *mastodon.Client, req sdk.GetStatusesRequest) (rsp sdk.GetStatusesResponse) {
	rsp.Meta = &sdk.CollectionMeta{}

	err := func() (err error) {
		var status *mastodon.Status
		switch *req.Mode.Id.Kind {
		case sdk.IdKind.ServiceId:
			status, err = c.GetStatus(ctx, mastodon.ID(*req.Mode.Id.ServiceId.Value))
			if err != nil {
				return
			}
		default:
		}

		rsp.Statuses = []sdk.Status{MapStatusFromMastodonStatus(*status)}

		return
	}()
	if err != nil {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Formatting: &sdk.FormattingKind.Plain,
				Value:      ptr.String(err.Error()),
			},
		})
	}

	return
}

func getStatusesSearch(ctx context.Context, c *mastodon.Client, req sdk.GetStatusesRequest) (rsp sdk.GetStatusesResponse) {
	rsp.Meta = &sdk.CollectionMeta{}

	err := func() (err error) {
		results, err := c.Search(ctx, *req.Mode.Search.Term, false)
		if err != nil {
			return
		}

		rsp.Statuses = MapStatusesFromMastodonStatuses(results.Statuses)

		return
	}()
	if err != nil {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Formatting: &sdk.FormattingKind.Plain,
				Value:      ptr.String(err.Error()),
			},
		})
	}

	return
}

func getStatusesRelation(ctx context.Context, c *mastodon.Client, req sdk.GetStatusesRequest) (rsp sdk.GetStatusesResponse) {
	rsp.Meta = &sdk.CollectionMeta{}

	var statuses []*mastodon.Status

	//var page *sdk.Page
	pg := &mastodon.Pagination{}

	//if len(req.Mode.Relation.Pages) > 0 {
	//	page = req.Mode.Relation.Pages[0].Page
	//}
	//
	//if page != nil &&
	//	page.CursorPage != nil {
	//	if page.CursorPage.Next != nil {
	//		pg.MaxID = mastodon.ID(*page.CursorPage.Next)
	//	}
	//
	//	if page.CursorPage.Previous != nil {
	//		pg.SinceID = mastodon.ID(*page.CursorPage.Previous)
	//	}
	//}

	err := func() (err error) {
		switch *req.Mode.Relation.Relation {
		case sdk.StatusRelationName.StatusWasRepliedToByStatuses:
			//var c *mastodon.Context
			//c, err = c.GetStatusContext(ctx, mastodon.ID(*req.Mode.Relation.Id.Value))
			//if err != nil {
			//	return
			//}
			//
			//statuses = c.Descendants
		// todo scope to me
		case sdk.PersonRelationName.PersonFavorsStatuses:
			statuses, err = c.GetFavourites(ctx, pg)
			if err != nil {
				return
			}
		case sdk.PersonRelationName.PersonAuthorsStatuses:
			// todo v1: support IdUnion
			//var id mastodon.ID
			//if *req.Mode.Relation.Id.Kind == sdk.ID_ME {
			//	var acc *mastodon.Account
			//	acc, err = c.GetAccountCurrentUser(ctx)
			//	if err != nil {
			//		return
			//	}
			//
			//	id = acc.ID
			//}

			id := mastodon.ID(*req.Mode.Relation.Id.Value)

			statuses, err = c.GetAccountStatuses(ctx, id, pg)
			if err != nil {
				return
			}

			break
		case sdk.FeedRelationName.FeedContainsStatuses:
			switch *req.Mode.Relation.Id.Value {
			case TIMELINE_PUBLIC:
				statuses, err = c.GetTimelinePublic(ctx, false, pg)
				if err != nil {
					return
				}

				break
			case TIMELINE_PUBLIC_LOCAL:
				statuses, err = c.GetTimelinePublic(ctx, true, pg)
				if err != nil {
					return
				}

				break
			// todo scope to me
			case TIMELINE_HOME:
				statuses, err = c.GetTimelineHome(ctx, pg)
				if err != nil {
					return
				}

				break
			// todo scope to me?
			case TIMELINE_MEDIA:
				statuses, err = c.GetTimelineMedia(ctx, false, pg)
				if err != nil {
					return
				}

				break
			// todo scope to me?
			case TIMELINE_MEDIA_LOCAL:
				statuses, err = c.GetTimelineMedia(ctx, true, pg)
				if err != nil {
					return
				}

				break
			}
		}

		return
	}()
	if err != nil {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Formatting: &sdk.FormattingKind.Plain,
				Value:      ptr.String(err.Error()),
			},
		})
	}

	//if pg != nil {
	//	pagination := &sdk.Pagination{
	//		Previous: &sdk.Page{
	//			CursorPage: &sdk.CursorPage{
	//				Value: ptr.String(string(pg.SinceID)),
	//			},
	//		},
	//		Next: &sdk.Page{
	//			CursorPage: &sdk.CursorPage{
	//				Value: ptr.String(string(pg.MaxID)),
	//			},
	//		},
	//	}
	//
	//	if page != nil {
	//		pagination.Current = page
	//	}
	//
	//	rsp.Meta.Pagination = pagination
	//}

	rsp.Statuses = MapStatusesFromMastodonStatuses(statuses)

	return
}

func putStatusesRelation(ctx context.Context, c *mastodon.Client, req sdk.PutStatusesRequest) (rsp sdk.PutStatusesResponse) {
	rsp.Meta = &sdk.ResponseMeta{}

	add := func() (errs []error) {
		switch *req.Mode.Relation.Relation {
		case sdk.PersonRelationName.PersonFavorsStatuses:
			for _, id := range req.Mode.Relation.Ids {
				_, err := c.Favourite(ctx, mastodon.ID(*id.Value))
				if err != nil {
					errs = append(errs, err)
				}
			}
		case sdk.StatusRelationName.StatusRebloggedByStatuses:
			for _, id := range req.Mode.Relation.Ids {
				_, err := c.Reblog(ctx, mastodon.ID(*id.Value))
				if err != nil {
					errs = append(errs, err)
				}
			}
		default:
		}

		return
	}

	remove := func() (errs []error) {
		switch *req.Mode.Relation.Relation {
		case sdk.PersonRelationName.PersonFavorsStatuses:
			for _, id := range req.Mode.Relation.Ids {
				_, err := c.Unfavourite(ctx, mastodon.ID(*id.Value))
				if err != nil {
					errs = append(errs, err)
				}
			}
		case sdk.StatusRelationName.StatusRebloggedByStatuses:
			for _, id := range req.Mode.Relation.Ids {
				_, err := c.Unreblog(ctx, mastodon.ID(*id.Value))
				if err != nil {
					errs = append(errs, err)
				}
			}
		default:
		}

		return
	}

	var errs []error
	if *req.Mode.Relation.Operation == sdk.RelationOperation.Add {
		errs = add()
	} else {
		errs = remove()
	}

	for _, err := range errs {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Formatting: &sdk.FormattingKind.Plain,
				Value:      ptr.String(err.Error()),
			},
		})
	}

	return
}

func postStatuses(ctx context.Context, c *mastodon.Client, req sdk.PostStatusesRequest) (rsp sdk.PostStatusesResponse) {
	rsp.Meta = &sdk.ResponseMeta{}

	var statuses []*mastodon.Status
	var errs []error
	for _, status := range req.Statuses {
		status, err := c.PostStatus(ctx, MapStatusToMastodonToot(status))
		if err != nil {
			errs = append(errs, err)
		} else {
			statuses = append(statuses, status)
		}
	}

	for _, err := range errs {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Formatting: &sdk.FormattingKind.Plain,
				Value:      ptr.String(err.Error()),
			},
		})
	}

	rsp.Statuses = MapStatusesFromMastodonStatuses(statuses)

	return
}

func deleteStatuses(ctx context.Context, c *mastodon.Client, req sdk.DeleteStatusesRequest) (rsp sdk.DeleteStatusesResponse) {
	rsp.Meta = &sdk.ResponseMeta{}

	var errs []error
	for _, id := range req.Ids {
		err := c.DeleteStatus(ctx, mastodon.ID(*id.Value))
		if err != nil {
			errs = append(errs, err)
		}
	}

	for _, err := range errs {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Formatting: &sdk.FormattingKind.Plain,
				Value:      ptr.String(err.Error()),
			},
		})
	}

	return
}