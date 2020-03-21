package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"

	"net/http"
)

type story struct {
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

func getStatusesId(c *http.Client, req sdk.GetStatusesRequest) (ss []sdk.Status, errs []error) {
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

		s := story{}
		err = json.NewDecoder(rsp.Body).Decode(&s)
		if err != nil {
			return
		}

		ss = append(ss, mapStoryToStatus(s))

		return
	}()
	if err != nil {
		errs = append(errs, err)
	}

	return
}

func getStatusesRelation(c *http.Client, req sdk.GetStatusesRequest) (ss []sdk.Status, errs []error) {
	err := func() (err error) {
		switch sdk.IdKind.ServiceId {
		case sdk.IdKind.ServiceId:
			switch *req.Mode.Relation.Relation {
			case sdk.SocialAccountRelationName.SocialAccountAuthorsStatuses:
				var as []sdk.SocialAccount
				as, errs = getSocialAccountId(c, sdk.GetSocialAccountsRequest{
					Mode: &sdk.GetMode{
						Kind: &sdk.GetModeKind.Id,
						Id: &sdk.Id{
							Kind: &sdk.IdKind.ServiceId,
							ServiceId: req.Mode.Relation.Id,
						},
					},
				})
				if len(errs) != 0 {
					return
				}

				ss = as[0].Relations.AuthorsStatuses.Statuses

				break
			default:
				err = errors.New(fmt.Sprintf("can't handle id %v", req.Mode.Id))

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

func mapStoryToStatus(s story) (s0 sdk.Status) {
	//type story struct {
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

	s0 = sdk.Status{
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
		Meta: &sdk.TypeMeta{
			CreatedAt: &sdk.Timestamp{
				Kind: &sdk.TimestampKind.Unix,
				Value: &sdk.DurationScalar{
					Unit:  &sdk.DurationUnit.S,
					Value: sdk.Float64(float64(*s.Time)),
				},
			},
		},
		Relations: &sdk.StatusRelations{
			AuthoredBySocialAccount: &sdk.SocialAccount{
				Id: &sdk.ServiceId{
					Value: s.By,
				},
			},
			WasRepliedToByStatuses: func() (c *sdk.StatusesCollection) {
				if s.Descandants == nil && s.Kids == nil {
					return
				}

				return &sdk.StatusesCollection{
					Meta: &sdk.CollectionMeta{
						Count: s.Descandants,
					},
					Statuses: func() (ss []sdk.Status) {
						for _, k := range s.Kids {
							ss = append(ss, sdk.Status{
								Id: &sdk.ServiceId{
									Value: sdk.String(fmt.Sprintf("%v", k)),
								},
							})
						}

						return
					}(),
				}
			}(),
			RepliesToStatus: func() (s1 *sdk.Status) {
				if s.Parent == nil {
					return
				}

				return &sdk.Status{
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
					Meta: &sdk.CollectionMeta{
						Count: s.Score,
					},
				}
			}(),
		},
	}

	return s0
}
