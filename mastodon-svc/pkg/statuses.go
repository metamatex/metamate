package pkg

import (
	"context"
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamate/gen/v0/sdk"
)

func getPostId(ctx context.Context, c *mastodon.Client, req sdk.GetPostsRequest) (rsp sdk.GetPostsResponse) {
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

		rsp.Posts = []sdk.Post{MapPostFromStatus(*status)}

		return
	}()
	if err != nil {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Formatting: &sdk.FormattingKind.Plain,
				Value:      sdk.String(err.Error()),
			},
		})
	}

	return
}

func getPostsSearch(ctx context.Context, c *mastodon.Client, req sdk.GetPostsRequest) (rsp sdk.GetPostsResponse) {
	rsp.Meta = &sdk.CollectionMeta{}

	err := func() (err error) {
		results, err := c.Search(ctx, *req.Mode.Search.Term, false)
		if err != nil {
			return
		}

		rsp.Posts = MapPostsFromStatuses(results.Statuses)

		return
	}()
	if err != nil {
		rsp.Meta.Errors = append(rsp.Meta.Errors, sdk.Error{
			Message: &sdk.Text{
				Formatting: &sdk.FormattingKind.Plain,
				Value:      sdk.String(err.Error()),
			},
		})
	}

	return
}

func getPostsRelation(ctx context.Context, c *mastodon.Client, req sdk.GetPostsRequest) (rsp sdk.GetPostsResponse) {
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
		case sdk.PostRelationName.PostWasRepliedToByPosts:
			//var c *mastodon.Context
			//c, err = c.GetStatusContext(ctx, mastodon.ID(*req.Mode.Relation.Id.Value))
			//if err != nil {
			//	return
			//}
			//
			//statuses = c.Descendants
		// todo scope to me
		case sdk.SocialAccountRelationName.SocialAccountFavorsPosts:
			statuses, err = c.GetFavourites(ctx, pg)
			if err != nil {
				return
			}
		case sdk.SocialAccountRelationName.SocialAccountAuthorsPosts:
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
		case sdk.PostFeedRelationName.PostFeedContainsPosts:
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
				Value:      sdk.String(err.Error()),
			},
		})
	}

	//if pg != nil {
	//	pagination := &sdk.Pagination{
	//		Previous: &sdk.Page{
	//			CursorPage: &sdk.CursorPage{
	//				Value: sdk.String(string(pg.SinceID)),
	//			},
	//		},
	//		Next: &sdk.Page{
	//			CursorPage: &sdk.CursorPage{
	//				Value: sdk.String(string(pg.MaxID)),
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

	rsp.Posts = MapPostsFromStatuses(statuses)

	return
}