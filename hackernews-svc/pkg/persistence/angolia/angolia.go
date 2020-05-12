package angolia

import (
	"encoding/json"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"
	"net/http"
	"net/url"
)

func GetPostsSearch(c *http.Client, req sdk.GetPostsRequest) (ss []sdk.Post, errs []sdk.Error, pagination *sdk.Pagination) {
	err := func() (err error) {
		var u string

		var page *sdk.ServicePage
		if len(req.Pages) > 0 {
			page = &req.Pages[0]
		}

		if page == nil {
			page = &sdk.ServicePage{
				Page: &sdk.Page{
					IndexPage: &sdk.IndexPage{
						Value: sdk.Int32(0),
					},
				},
			}
		}

		var pageIndex int32 = 0
		if page.Page != nil &&
			page.Page.IndexPage != nil &&
			page.Page.IndexPage.Value != nil {
			pageIndex = *page.Page.IndexPage.Value
		}

		pagination = &sdk.Pagination{
			Current: []sdk.ServicePage{
				{
					Page: &sdk.Page{
						Kind: &sdk.PageKind.IndexPage,
						IndexPage: &sdk.IndexPage{
							Value: sdk.Int32(pageIndex),
						},
					},
				},
			},
		}

		if pageIndex > 0 {
			pagination.Previous = []sdk.ServicePage{
				{
					Page: &sdk.Page{
						Kind: &sdk.PageKind.IndexPage,
						IndexPage: &sdk.IndexPage{
							Value: sdk.Int32(pageIndex - 1),
						},
					},
				},
			}
		}
		hitsPerPage := 101

		u = fmt.Sprintf("http://hn.algolia.com/api/v1/search?query=%v&hitsPerPage=%v&page=%v", url.QueryEscape(*req.Mode.Search.Term), hitsPerPage, pageIndex)

		rsp, err := c.Get(u)
		if err != nil {
			return
		}
		defer rsp.Body.Close()

		var r struct {
			Hits []searchHnStory
		}
		err = json.NewDecoder(rsp.Body).Decode(&r)
		if err != nil {
			return
		}

		for _, s := range r.Hits {
			ss = append(ss, mapSearchHNStoryToPost(s))
		}

		if len(ss) == hitsPerPage {
			pagination.Next = []sdk.ServicePage{
				{
					Page: &sdk.Page{
						Kind: &sdk.PageKind.IndexPage,
						IndexPage: &sdk.IndexPage{
							Value: sdk.Int32(pageIndex + 1),
						},
					},
				},
			}
		}

		return
	}()
	if err != nil {
		errs = append(errs, sdk.Error{
			Message: sdk.String(err.Error()),
		})
	}

	return
}

type searchHnStory struct {
	CreatedAt      *string `json:"created_at"`
	Title          *string
	Url            *string
	Author         *string
	Points         *int32
	StoryText      *string `json:"story_text"`
	CommentText    *string `json:"comment_text"`
	NumComments    *int32  `json:"num_comments"`
	StoryId        *int32  `json:"story_id"`
	StoryTitle     *string `json:"story_title"`
	StoryUrl       *string `json:"story_url"`
	ParentId       *int32  `json:"parent_id"`
	CreateAtI      *int32  `json:"created_at_i"`
	RelevanceScore *int32  `json:"relevance_score"`
	ObjectId       *string `json:"objectID"`
}

func mapSearchHNStoryToPostX(s searchHnStory) (p sdk.Post) {
	//type searchHnStory struct {
	//	CreatedAt      *string `json:"created_at"`
	//	Title          *string
	//	Url            *string
	//	Author         *string
	//	Points         *int
	//	StoryText      *string `json:"story_text"`
	//	CommentText    *string `json:"comment_text"`
	//	NumComments    *int    `json:"num_comments"`
	//	StoryId        *int    `json:"story_id"`
	//	StoryTitle     *string `json:"story_title"`
	//	StoryUrl       *string `json:"story_url"`
	//	ParentId       *int    `json:"parent_id"`
	//	CreateAtI      *int    `json:"created_at_i"`
	//	RelevanceScore *int    `json:"relevance_score"`
	//	x ObjectId       *string `json:"objectID"`
	//}

	return sdk.Post{
		Id: &sdk.ServiceId{
			Value: s.ObjectId,
		},
	}
}

func mapSearchHNStoryToPost(s searchHnStory) (p sdk.Post) {
	//type searchHnStory struct {
	//	CreatedAt      *string `json:"created_at"`
	//	Title          *string
	//	Url            *string
	//	Author         *string
	//	Points         *int
	//	StoryText      *string `json:"story_text"`
	//	CommentText    *string `json:"comment_text"`
	//	NumComments    *int    `json:"num_comments"`
	//	StoryId        *int    `json:"story_id"`
	//	StoryTitle     *string `json:"story_title"`
	//	StoryUrl       *string `json:"story_url"`
	//	ParentId       *int    `json:"parent_id"`
	//	CreateAtI      *int    `json:"created_at_i"`
	//	RelevanceScore *int    `json:"relevance_score"`
	//	x ObjectId       *string `json:"objectID"`
	//}

	return sdk.Post{
		Id: &sdk.ServiceId{
			Value: s.ObjectId,
		},
		Kind: func() *string {
			if s.ParentId == nil {
				return &sdk.PostKind.Post
			} else {
				return &sdk.PostKind.Reply
			}
		}(),
		TotalWasRepliedToByPostsCount: s.NumComments,
		AlternativeIds: []sdk.Id{
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(fmt.Sprintf("https://news.ycombinator.com/item?id=%v", *s.ObjectId)),
				},
			},
		},
		Title: func() *sdk.Text {
			if s.Title == nil {
				return nil
			}

			return &sdk.Text{
				Formatting: &sdk.FormattingKind.Plain,
				Value:      s.Title,
			}
		}(),
		Content: func() *sdk.Text {
			if s.StoryText == nil {
				return nil
			}

			return &sdk.Text{
				Formatting: &sdk.FormattingKind.Html,
				Value:      s.StoryText,
			}
		}(),
		Links: func() []sdk.HyperLink {
			if s.Url == nil {
				return nil
			}

			return []sdk.HyperLink{
				{
					Label: s.Title,
					Url: &sdk.Url{
						Value: s.Url,
					},
				},
			}
		}(),
		CreatedAt: func() (ts *sdk.Timestamp) {
			if s.CreateAtI == nil {
				return
			}

			return &sdk.Timestamp{
				Kind: &sdk.TimestampKind.Unix,
				Unix: &sdk.DurationScalar{
					Unit:  &sdk.DurationUnit.S,
					Value: sdk.Float64(float64(*s.CreateAtI)),
				},
			}
		}(),
		Relations: &sdk.PostRelations{
			AuthoredBySocialAccount: &sdk.SocialAccount{
				Id: &sdk.ServiceId{
					Value: s.Author,
				},
			},
			RepliesToPost: func() (s1 *sdk.Post) {
				if s.ParentId == nil {
					return
				}

				return &sdk.Post{
					Id: &sdk.ServiceId{
						Value: sdk.String(fmt.Sprintf("%v", *s.ParentId)),
					},
				}
			}(),
			FavoredBySocialAccounts: func() (c *sdk.SocialAccountsCollection) {
				if s.Points == nil {
					return
				}

				return &sdk.SocialAccountsCollection{
					Count: s.Points,
				}
			}(),
		},
	}
}
