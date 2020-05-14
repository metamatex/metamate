package angolia

import (
	"encoding/json"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/mql"
	"net/http"
	"net/url"
)

func GetPostsSearch(c *http.Client, req mql.GetPostsRequest) (ss []mql.Post, errs []mql.Error, pagination *mql.Pagination) {
	err := func() (err error) {
		var u string

		var page *mql.ServicePage
		if len(req.Pages) > 0 {
			page = &req.Pages[0]
		}

		if page == nil {
			page = &mql.ServicePage{
				Page: &mql.Page{
					IndexPage: &mql.IndexPage{
						Value: mql.Int32(0),
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

		hitsPerPage := 1000

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

		return
	}()
	if err != nil {
		errs = append(errs, mql.Error{
			Message: mql.String(err.Error()),
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

func mapSearchHNStoryToPostX(s searchHnStory) (p mql.Post) {
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

	return mql.Post{
		Id: &mql.ServiceId{
			Value: s.ObjectId,
		},
	}
}

func mapSearchHNStoryToPost(s searchHnStory) (p mql.Post) {
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

	return mql.Post{
		Id: &mql.ServiceId{
			Value: s.ObjectId,
		},
		Kind: func() *string {
			if s.ParentId == nil {
				return &mql.PostKind.Post
			} else {
				return &mql.PostKind.Reply
			}
		}(),
		TotalWasRepliedToByPostsCount: s.NumComments,
		AlternativeIds: []mql.Id{
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(fmt.Sprintf("https://news.ycombinator.com/item?id=%v", *s.ObjectId)),
				},
			},
		},
		Title: func() *mql.Text {
			if s.Title == nil {
				return nil
			}

			return &mql.Text{
				Formatting: &mql.FormattingKind.Plain,
				Value:      s.Title,
			}
		}(),
		Content: func() *mql.Text {
			if s.StoryText == nil {
				return nil
			}

			return &mql.Text{
				Formatting: &mql.FormattingKind.Html,
				Value:      s.StoryText,
			}
		}(),
		Links: func() []mql.HyperLink {
			if s.Url == nil {
				return nil
			}

			return []mql.HyperLink{
				{
					Label: s.Title,
					Url: &mql.Url{
						Value: s.Url,
					},
				},
			}
		}(),
		CreatedAt: func() (ts *mql.Timestamp) {
			if s.CreateAtI == nil {
				return
			}

			return &mql.Timestamp{
				Kind: &mql.TimestampKind.Unix,
				Unix: &mql.DurationScalar{
					Unit:  &mql.DurationUnit.S,
					Value: mql.Float64(float64(*s.CreateAtI)),
				},
			}
		}(),
		Relations: &mql.PostRelations{
			AuthoredBySocialAccount: &mql.SocialAccount{
				Id: &mql.ServiceId{
					Value: s.Author,
				},
			},
			RepliesToPost: func() (s1 *mql.Post) {
				if s.ParentId == nil {
					return
				}

				return &mql.Post{
					Id: &mql.ServiceId{
						Value: mql.String(fmt.Sprintf("%v", *s.ParentId)),
					},
				}
			}(),
			FavoredBySocialAccounts: func() (c *mql.SocialAccountsCollection) {
				if s.Points == nil {
					return
				}

				return &mql.SocialAccountsCollection{
					Count: s.Points,
				}
			}(),
		},
	}
}
