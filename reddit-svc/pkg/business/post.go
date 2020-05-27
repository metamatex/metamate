package business

import (
	"errors"
	"fmt"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/reddit-svc/pkg/communication"
	"github.com/metamatex/metamate/reddit-svc/pkg/types"
)

func GetPostRelation(c communication.Client, req mql.GetPostsRequest) (ps []mql.Post, pagination mql.Pagination, errs []mql.Error) {
	err := func() (err error) {
		var after *string
		if len(req.Pages) > 0 {
			after = req.Pages[0].Page.CursorPage.Value
		}

		limit := 100
		time := "all"
		var rsp types.GetSubredditSubmissionsResponse

		switch *req.Mode.Relation.Relation {
		case mql.PathNames.PostFeedContainsPosts:
			switch *req.Mode.Relation.Id.Kind {
			case mql.IdKind.Name:
				rsp, err = c.GetSubredditSubmissions(*req.Mode.Relation.Id.Name, "new", &limit, &time, after)
				if err != nil {
					return
				}
			case mql.IdKind.ServiceId:
				rsp, err = c.GetSubredditSubmissions(*req.Mode.Relation.Id.ServiceId.Value, "new", &limit, &time, after)
				if err != nil {
					return
				}
			}

			err = types.GetError(rsp.Error)
			if err != nil {
				return
			}

			for _, c := range rsp.Data.Children {
				ps = append(ps, mapPostListingChildToPost(c.Data))
			}
		case mql.PathNames.SocialAccountAuthorsPosts:
			var rsp types.GetSubredditSubmissionsResponse
			switch *req.Mode.Relation.Id.Kind {
			case mql.IdKind.Username:
				rsp, err = c.GetUserSubmissions(*req.Mode.Relation.Id.Username, "new", &limit, &time, after)
				if err != nil {
					return
				}
			case mql.IdKind.ServiceId:
				rsp, err = c.GetUserSubmissions(*req.Mode.Relation.Id.ServiceId.Value, "new", &limit, &time, after)
				if err != nil {
					return
				}
			}

			err = types.GetError(rsp.Error)
			if err != nil {
				return
			}

			for _, c := range rsp.Data.Children {
				ps = append(ps, mapPostListingChildToPost(c.Data))
			}
		default:
			err = errors.New(fmt.Sprintf("can't handle relation %v", *req.Mode.Relation.Relation))
		}

		if after != nil {
			pagination.Current = []mql.ServicePage{
				{
					Page: &mql.Page{
						Kind: &mql.PageKind.CursorPage,
						CursorPage: &mql.CursorPage{
							Value: after,
						},
					},
				},
			}
		}

		if rsp.Data.After != nil {
			pagination.Next = []mql.ServicePage{
				{
					Page: &mql.Page{
						Kind: &mql.PageKind.CursorPage,
						CursorPage: &mql.CursorPage{
							Value: rsp.Data.After,
						},
					},
				},
			}
		}

		return
	}()
	if err != nil {
		errs = append(errs, mql.Error{
			Message: mql.String(err.Error()),
		})
	}

	return
}

func mapPostListingChildToPost(d types.PostListingChildData) (f mql.Post) {
	return mql.Post{
		Id: &mql.ServiceId{
			Value: d.Id,
		},
		//Kind: func() *string {
		//	if s.Parent == nil {
		//		return &mql.PostKind.Post
		//	} else {
		//		return &mql.PostKind.Reply
		//	}
		//}(),
		TotalWasRepliedToByPostsCount: d.NumComments,
		AlternativeIds: []mql.Id{
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: d.Url,
				},
			},
		},
		Title: func() *mql.Text {
			if d.Title == nil {
				return nil
			}

			return &mql.Text{
				Formatting: &mql.FormattingKind.Plain,
				Value:      d.Title,
			}
		}(),
		Content: func() *mql.Text {
			if d.SelftextHtml == nil {
				return nil
			}

			return &mql.Text{
				Formatting: &mql.FormattingKind.Html,
				Value:      d.SelftextHtml,
			}
		}(),
		//Links: func() []mql.HyperLink {
		//	if s.Url == nil {
		//		return nil
		//	}
		//
		//	return []mql.HyperLink{
		//		{
		//			Label: s.Title,
		//			Url: &mql.Url{
		//				Value: s.Url,
		//			},
		//		},
		//	}
		//}(),
		IsSensitive: d.Over18,
		CreatedAt: func() (ts *mql.Timestamp) {
			if d.CreatedUtc == nil {
				return
			}

			return &mql.Timestamp{
				Kind: &mql.TimestampKind.Unix,
				Unix: &mql.DurationScalar{
					Unit:  &mql.DurationUnit.S,
					Value: d.CreatedUtc,
				},
			}
		}(),
		Relations: &mql.PostRelations{
			AuthoredBySocialAccount: &mql.SocialAccount{
				Id: &mql.ServiceId{
					Value: d.Author,
				},
				AlternativeIds: []mql.Id{
					{
						Kind:     &mql.IdKind.Username,
						Username: d.Author,
					},
				},
			},
			FavoredBySocialAccounts: func() (c *mql.SocialAccountsCollection) {
				if d.Ups == nil {
					return
				}

				return &mql.SocialAccountsCollection{
					Count: d.Ups,
				}
			}(),
		},
	}
}
