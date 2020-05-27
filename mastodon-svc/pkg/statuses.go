package pkg

import (
	"context"
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamate/gen/v0/mql"
)

func getPostId(ctx context.Context, c *mastodon.Client, req mql.GetPostsRequest) (rsp mql.GetPostsResponse) {
	err := func() (err error) {
		var status *mastodon.Status
		switch *req.Mode.Id.Kind {
		case mql.IdKind.ServiceId:
			status, err = c.GetStatus(ctx, mastodon.ID(*req.Mode.Id.ServiceId.Value))
			if err != nil {
				return
			}
		default:
		}

		rsp.Posts = []mql.Post{MapPostFromStatus(*status)}

		return
	}()
	if err != nil {
		rsp.Errors = append(rsp.Errors, mql.Error{
			Message: mql.String(err.Error()),
		})
	}

	return
}

func getPostsSearch(ctx context.Context, c *mastodon.Client, req mql.GetPostsRequest) (rsp mql.GetPostsResponse) {
	err := func() (err error) {
		results, err := c.Search(ctx, *req.Mode.Search.Term, false)
		if err != nil {
			return
		}

		rsp.Posts = MapPostsFromStatuses(results.Statuses)

		return
	}()
	if err != nil {
		rsp.Errors = append(rsp.Errors, mql.Error{
			Message: mql.String(err.Error()),
		})
	}

	return
}

func getPostsRelation(ctx context.Context, c *mastodon.Client, req mql.GetPostsRequest) (rsp mql.GetPostsResponse) {
	var statuses []*mastodon.Status

	//var page *mql.Page
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
		case mql.PostRelationName.PostWasRepliedToByPosts:
			//var c *mastodon.Context
			//c, err = c.GetStatusContext(ctx, mastodon.ID(*req.Mode.Relation.Id.Value))
			//if err != nil {
			//	return
			//}
			//
			//statuses = c.Descendants
		// todo scope to me
		case mql.SocialAccountRelationName.SocialAccountFavorsPosts:
			statuses, err = c.GetFavourites(ctx, pg)
			if err != nil {
				return
			}
		case mql.SocialAccountRelationName.SocialAccountAuthorsPosts:
			// todo v1: support IdUnion
			//var id mastodon.ID
			//if *req.Mode.Relation.Id.Kind == mql.ID_ME {
			//	var acc *mastodon.Account
			//	acc, err = c.GetAccountCurrentUser(ctx)
			//	if err != nil {
			//		return
			//	}
			//
			//	id = acc.ID
			//}

			id := mastodon.ID(*req.Mode.Relation.Id.ServiceId.Value)

			statuses, err = c.GetAccountStatuses(ctx, id, pg)
			if err != nil {
				return
			}

			break
		case mql.PostFeedRelationName.PostFeedContainsPosts:
			switch *req.Mode.Relation.Id.ServiceId.Value {
			case TimelinePublic:
				statuses, err = c.GetTimelinePublic(ctx, false, pg)
				if err != nil {
					return
				}

				break
			case TimelinePublicLocal:
				statuses, err = c.GetTimelinePublic(ctx, true, pg)
				if err != nil {
					return
				}

				break
			// todo scope to me
			case TimelineHome:
				statuses, err = c.GetTimelineHome(ctx, pg)
				if err != nil {
					return
				}

				break
			// todo scope to me?
			case TimelineMedia:
				statuses, err = c.GetTimelineMedia(ctx, false, pg)
				if err != nil {
					return
				}

				break
			// todo scope to me?
			case TimelineMediaLocal:
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
		rsp.Errors = append(rsp.Errors, mql.Error{
			Message: mql.String(err.Error()),
		})
	}

	//if pg != nil {
	//	pagination := &mql.Pagination{
	//		Previous: &mql.Page{
	//			CursorPage: &mql.CursorPage{
	//				Value: mql.String(string(pg.SinceID)),
	//			},
	//		},
	//		Next: &mql.Page{
	//			CursorPage: &mql.CursorPage{
	//				Value: mql.String(string(pg.MaxID)),
	//			},
	//		},
	//	}
	//
	//	if page != nil {
	//		pagination.Current = page
	//	}
	//
	//	rsp.Pagination = pagination
	//}

	rsp.Posts = MapPostsFromStatuses(statuses)

	return
}
