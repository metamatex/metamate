package firebase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"
	"github.com/metamatex/metamate/hackernews-svc/pkg/types"
	"net/http"
	"strconv"
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
	Descendants *int32
}

func GetPostsId(c *http.Client, req sdk.GetPostsRequest) (ss []sdk.Post, errs []sdk.Error) {
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
		defer rsp.Body.Close()

		s := firebaseStory{}
		err = json.NewDecoder(rsp.Body).Decode(&s)
		if err != nil {
			return
		}

		ss = append(ss, mapFirebaseStoryToPost(s))

		return
	}()
	if err != nil {
		errs = append(errs, sdk.Error{
			Message: sdk.String(err.Error()),
		})
	}

	return
}

func GetPostFeedContainsPosts(c *http.Client, feed string) (ss []sdk.Post, errs []sdk.Error) {
	err := func() (err error) {
		m := map[string]string{
			types.TopStories:  "https://hacker-news.firebaseio.com/v0/topstories.json",
			types.NewStories:  "https://hacker-news.firebaseio.com/v0/newstories.json",
			types.BestStories: "https://hacker-news.firebaseio.com/v0/beststories.json",
			types.AskStories:  "https://hacker-news.firebaseio.com/v0/askstories.json",
			types.ShowStories: "https://hacker-news.firebaseio.com/v0/showstories.json",
			types.JobStories:  "https://hacker-news.firebaseio.com/v0/jobstories.json",
		}

		rsp, err := c.Get(m[feed])
		if err != nil {
			return
		}
		defer rsp.Body.Close()

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
		errs = append(errs, sdk.Error{
			Message: sdk.String(err.Error()),
		})
	}

	return
}

func mapFirebaseStoryToPost(s firebaseStory) sdk.Post {
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
	//	Descendants int32        x
	//}

	return sdk.Post{
		Id: &sdk.ServiceId{
			Value: sdk.String(fmt.Sprintf("%v", *s.Id)),
		},
		Kind: func() *string {
			if s.Parent == nil {
				return &sdk.PostKind.Post
			} else {
				return &sdk.PostKind.Reply
			}
		}(),
		TotalWasRepliedToByPostsCount: s.Descendants,
		AlternativeIds: []sdk.Id{
			{
				Kind: &sdk.IdKind.Url,
				Url: &sdk.Url{
					Value: sdk.String(fmt.Sprintf("https://news.ycombinator.com/item?id=%v", *s.Id)),
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
			if s.Text == nil {
				return nil
			}

			return &sdk.Text{
				Formatting: &sdk.FormattingKind.Html,
				Value:      s.Text,
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
			if s.Time == nil {
				return
			}

			return &sdk.Timestamp{
				Kind: &sdk.TimestampKind.Unix,
				Unix: &sdk.DurationScalar{
					Unit:  &sdk.DurationUnit.S,
					Value: sdk.Float64(float64(*s.Time)),
				},
			}
		}(),
		Relations: &sdk.PostRelations{
			AuthoredBySocialAccount: &sdk.SocialAccount{
				Id: &sdk.ServiceId{
					Value: s.By,
				},
			},
			WasRepliedToByPosts: func() (c *sdk.PostsCollection) {
				if s.Kids == nil {
					return
				}

				return &sdk.PostsCollection{
					Count: sdk.Int32(int32(len(s.Kids))),
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

func GetSocialAccountAuthorsPosts(c *http.Client, req sdk.GetPostsRequest) (ss []sdk.Post, errs []sdk.Error) {
	var as []sdk.SocialAccount
	as, errs = GetSocialAccountId(c, sdk.GetSocialAccountsRequest{
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

	return
}
