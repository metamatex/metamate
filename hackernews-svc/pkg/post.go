package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"
	"net/url"
	"strconv"

	"net/http"
)

type firebaseStory struct {
	Id          *int32
	Deleted     *bool
	Type        *string
	By          *string
	Time        *int32
	Text        *string
	Dead        *bool
	Parent      *int32
	Poll        *int32
	Kids        []int32
	Url         *string
	Score       *int32
	Title       *string
	Parts       []int32
	Descandants *int32
}

type searchHnStory struct {
	CreatedAt      *string `json:"created_at"`
	Title          *string
	Url            *string
	Author         *string
	Points         *int
	StoryText      *string `json:"story_text"`
	CommentText    *string `json:"comment_text"`
	NumComments    *int    `json:"num_comments"`
	StoryId        *int    `json:"story_id"`
	StoryTitle     *string `json:"story_title"`
	StoryUrl       *string `json:"story_url"`
	ParentId       *int    `json:"parent_id"`
	CreateAtI      *int    `json:"created_at_i"`
	RelevanceScore *int    `json:"relevance_score"`
	ObjectId       *string `json:"objectID"`
}

func getPostsId(c *http.Client, req sdk.GetPostsRequest) (ss []sdk.Post, errs []error) {
	err := func() (err error) {
		var url string

		switch *req.Mode.Id.Kind {
		case sdk.IdKind.ServiceId:
			url = fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", *req.Mode.Id.ServiceId.Value)
		default:
			err = errors.New(fmt.Sprintf("can't handle id %v", req.Mode.Id))

			return
		}

		rsp, err := c.Get(url)
		if err != nil {
			return
		}

		s := firebaseStory{}
		err = json.NewDecoder(rsp.Body).Decode(&s)
		if err != nil {
			return
		}

		ss = append(ss, mapFirebaseStoryToPost(s))

		return
	}()
	if err != nil {
		errs = append(errs, err)
	}

	return
}

func getPostsRelation(c *http.Client, req sdk.GetPostsRequest) (ss []sdk.Post, errs []error) {
	err := func() (err error) {
		switch sdk.IdKind.ServiceId {
		case sdk.IdKind.ServiceId:
			switch *req.Mode.Relation.Relation {
			case sdk.SocialAccountRelationName.SocialAccountAuthorsPosts:
				var as []sdk.SocialAccount
				as, errs = getSocialAccountId(c, sdk.GetSocialAccountsRequest{
					Mode: &sdk.GetMode{
						Kind: &sdk.GetModeKind.Id,
						Id: &sdk.Id{
							Kind:      &sdk.IdKind.ServiceId,
							ServiceId: req.Mode.Relation.Id,
						},
					},
				})
				if len(errs) != 0 {
					return
				}

				ss = as[0].Relations.AuthorsPosts.Posts

				break
			case sdk.PostFeedRelationName.PostFeedContainsPosts:
				ss, errs = getFeedContainsPosts(c, *req.Mode.Relation.Id.Value)

				break
			default:
				err = errors.New(fmt.Sprintf("can't handle relation %v", *req.Mode.Relation.Relation))

				return
			}
		default:
			err = errors.New(fmt.Sprintf("can't handle id %v", req.Mode.Id))

			return
		}

		return
	}()
	if err != nil {
		errs = append(errs, err)
	}

	return
}

func getFeedContainsPosts(c *http.Client, feed string) (ss []sdk.Post, errs []error) {
	err := func() (err error) {
		m := map[string]string{
			TopStories:  "https://hacker-news.firebaseio.com/v0/topstories.json",
			NewStories:  "https://hacker-news.firebaseio.com/v0/newstories.json",
			BestStories: "https://hacker-news.firebaseio.com/v0/beststories.json",
			AskStories:  "https://hacker-news.firebaseio.com/v0/askstories.json",
			ShowStories: "https://hacker-news.firebaseio.com/v0/showstories.json",
			JobStories:  "https://hacker-news.firebaseio.com/v0/jobstories.json",
		}

		rsp, err := c.Get(m[feed])
		if err != nil {
			return
		}

		var ids []int
		err = json.NewDecoder(rsp.Body).Decode(&ids)
		if err != nil {
			return
		}

		for _, id := range ids {
			ss = append(ss, sdk.Post{
				Id: &sdk.ServiceId{
					Value: sdk.String(strconv.Itoa(id)),
				},
			})
		}

		return
	}()
	if err != nil {
		errs = append(errs, err)
	}

	return
}

func mapFirebaseStoryToPost(s firebaseStory) (sdk.Post) {
	//type firebaseStory struct {
	//	Id          int32          x
	//	Deleted     bool
	//	Type        string
	//	By          string       x
	//	Time        int32          x
	//	Text        string       x-
	//	Dead        bool
	//	Parent      int32        x
	//	Poll        int32
	//	Kids        []int32      x
	//	Url         string       x-
	//	Score       int32        x
	//	Title       string       x-
	//	Parts       []int32
	//	Descandants int32        x
	//}

	return sdk.Post{
		Id: &sdk.ServiceId{
			Value: sdk.String(fmt.Sprintf("%v", *s.Id)),
		},
		AlternativeIds: []sdk.Id{
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(fmt.Sprintf("https://news.ycombinator.com/item?id=%v", *s.Id)),
				},
			},
		},
		Content: &sdk.Text{
			Formatting: &sdk.FormattingKind.Html,
			Value: func() *string {
				var v string

				if s.Title != nil {
					v += fmt.Sprintf("%v ", *s.Title)
				}

				if s.Text != nil {
					v += fmt.Sprintf("%v ", *s.Text)
				}

				if s.Url != nil {
					v += fmt.Sprintf("%v ", *s.Url)
				}

				return &v
			}(),
		},
		CreatedAt: &sdk.Timestamp{
			Kind: &sdk.TimestampKind.Unix,
			Unix: &sdk.DurationScalar{
				Unit:  &sdk.DurationUnit.S,
				Value: sdk.Float64(float64(*s.Time)),
			},
		},
		Relations: &sdk.PostRelations{
			AuthoredBySocialAccount: &sdk.SocialAccount{
				Id: &sdk.ServiceId{
					Value: s.By,
				},
			},
			WasRepliedToByPosts: func() (c *sdk.PostsCollection) {
				if s.Descandants == nil && s.Kids == nil {
					return
				}

				return &sdk.PostsCollection{
					Count: s.Descandants,
					Posts: func() (ss []sdk.Post) {
						for _, k := range s.Kids {
							ss = append(ss, sdk.Post{
								Id: &sdk.ServiceId{
									Value: sdk.String(fmt.Sprintf("%v", k)),
								},
							})
						}

						return
					}(),
				}
			}(),
			RepliesToPost: func() (s1 *sdk.Post) {
				if s.Parent == nil {
					return
				}

				return &sdk.Post{
					Id: &sdk.ServiceId{
						Value: sdk.String(fmt.Sprintf("%v", *s.Parent)),
					},
				}
			}(),
			FavoredBySocialAccounts: func() (c *sdk.SocialAccountsCollection) {
				if s.Score == nil {
					return
				}

				return &sdk.SocialAccountsCollection{
					Count: s.Score,
				}
			}(),
		},
	}
}

func getPostsSearch(c *http.Client, req sdk.GetPostsRequest) (ss []sdk.Post, errs []error) {
	err := func() (err error) {
		var u string

		u = fmt.Sprintf("http://hn.algolia.com/api/v1/search?query=%v", url.QueryEscape(*req.Mode.Search.Term))

		rsp, err := c.Get(u)
		if err != nil {
			return
		}



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
		errs = append(errs, err)
	}

	return
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
	}
}