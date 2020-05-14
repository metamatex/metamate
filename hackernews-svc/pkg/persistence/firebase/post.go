package firebase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/mql"
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

func GetPostsId(c *http.Client, req mql.GetPostsRequest) (ss []mql.Post, errs []mql.Error) {
	err := func() (err error) {
		var url string

		switch *req.Mode.Id.Kind {
		case mql.IdKind.ServiceId:
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
		errs = append(errs, mql.Error{
			Message: mql.String(err.Error()),
		})
	}

	return
}

func GetPostFeedContainsPosts(c *http.Client, feed string) (ss []mql.Post, errs []mql.Error) {
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
			ss = append(ss, mql.Post{
				Id: &mql.ServiceId{
					Value: mql.String(strconv.Itoa(id)),
				},
			})
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

func mapFirebaseStoryToPost(s firebaseStory) mql.Post {
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

	return mql.Post{
		Id: &mql.ServiceId{
			Value: mql.String(fmt.Sprintf("%v", *s.Id)),
		},
		Kind: func() *string {
			if s.Parent == nil {
				return &mql.PostKind.Post
			} else {
				return &mql.PostKind.Reply
			}
		}(),
		TotalWasRepliedToByPostsCount: s.Descendants,
		AlternativeIds: []mql.Id{
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(fmt.Sprintf("https://news.ycombinator.com/item?id=%v", *s.Id)),
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
			if s.Text == nil {
				return nil
			}

			return &mql.Text{
				Formatting: &mql.FormattingKind.Html,
				Value:      s.Text,
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
			if s.Time == nil {
				return
			}

			return &mql.Timestamp{
				Kind: &mql.TimestampKind.Unix,
				Unix: &mql.DurationScalar{
					Unit:  &mql.DurationUnit.S,
					Value: mql.Float64(float64(*s.Time)),
				},
			}
		}(),
		Relations: &mql.PostRelations{
			AuthoredBySocialAccount: &mql.SocialAccount{
				Id: &mql.ServiceId{
					Value: s.By,
				},
			},
			WasRepliedToByPosts: func() (c *mql.PostsCollection) {
				if s.Kids == nil {
					return
				}

				return &mql.PostsCollection{
					Count: mql.Int32(int32(len(s.Kids))),
					Posts: func() (ss []mql.Post) {
						for _, k := range s.Kids {
							ss = append(ss, mql.Post{
								Id: &mql.ServiceId{
									Value: mql.String(fmt.Sprintf("%v", k)),
								},
							})
						}

						return
					}(),
				}
			}(),
			RepliesToPost: func() (s1 *mql.Post) {
				if s.Parent == nil {
					return
				}

				return &mql.Post{
					Id: &mql.ServiceId{
						Value: mql.String(fmt.Sprintf("%v", *s.Parent)),
					},
				}
			}(),
			FavoredBySocialAccounts: func() (c *mql.SocialAccountsCollection) {
				if s.Score == nil {
					return
				}

				return &mql.SocialAccountsCollection{
					Count: s.Score,
				}
			}(),
		},
	}
}

func GetSocialAccountAuthorsPosts(c *http.Client, req mql.GetPostsRequest) (ss []mql.Post, errs []mql.Error) {
	var as []mql.SocialAccount
	as, errs = GetSocialAccountId(c, mql.GetSocialAccountsRequest{
		Mode: &mql.GetMode{
			Kind: &mql.GetModeKind.Id,
			Id: &mql.Id{
				Kind:      &mql.IdKind.ServiceId,
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
