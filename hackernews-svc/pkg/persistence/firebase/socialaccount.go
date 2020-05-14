package firebase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/mql"
	"net/http"
)

type user struct {
	Id        *string
	Created   *int
	Delay     *int
	Karma     *int
	About     *string
	Submitted []int
	Error     *string
}

func GetSocialAccountId(c *http.Client, req mql.GetSocialAccountsRequest) (as []mql.SocialAccount, errs []mql.Error) {
	err := func() (err error) {
		var url string

		switch *req.Mode.Id.Kind {
		case mql.IdKind.ServiceId:
			url = fmt.Sprintf("https://hacker-news.firebaseio.com/v0/user/%v.json", *req.Mode.Id.ServiceId.Value)
		case mql.IdKind.Username:
			url = fmt.Sprintf("https://hacker-news.firebaseio.com/v0/user/%v.json", *req.Mode.Id.Username)
		default:
			err = errors.New(fmt.Sprintf("can't handle id %v", req.Mode.Id))

			return
		}

		rsp, err := c.Get(url)
		if err != nil {
			return
		}
		defer rsp.Body.Close()

		u := &user{}
		err = json.NewDecoder(rsp.Body).Decode(&u)
		if err != nil {
			return
		}

		if u == nil || u.Error != nil {
			errs = append(errs, mql.Error{
				Kind: &mql.ErrorKind.IdNotPresent,
				Id:   req.Mode.Id,
			})
		} else {
			as = append(as, mapUserToSocialAccount(*u))
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

func mapUserToSocialAccount(u user) (a mql.SocialAccount) {
	//type user struct {
	//	Id        string    x
	//	Created   int       x
	//	Delay     int
	//	Karma     int
	//	About     string    x
	//	Submitted []int     x
	//}

	a = mql.SocialAccount{
		Id: &mql.ServiceId{
			Value: u.Id,
		},
		AlternativeIds: []mql.Id{
			{
				Kind:     &mql.IdKind.Username,
				Username: u.Id,
			},
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(fmt.Sprintf("https://news.ycombinator.com/user?id=%v", *u.Id)),
				},
			},
		},
		Points: mql.Int32(int32(*u.Karma)),
		Note: &mql.Text{
			Formatting: &mql.FormattingKind.Html,
			Value:      u.About,
		},
		CreatedAt: &mql.Timestamp{
			Kind: &mql.TimestampKind.Unix,
			Unix: &mql.DurationScalar{
				Unit:  &mql.DurationUnit.S,
				Value: mql.Float64(float64(*u.Created)),
			},
		},
		Username: u.Id,
		Relations: &mql.SocialAccountRelations{
			AuthorsPosts: &mql.PostsCollection{
				Count: mql.Int32(int32(len(u.Submitted))),
				Posts: func() (ss []mql.Post) {
					for _, s := range u.Submitted {
						ss = append(ss, mql.Post{
							Id: &mql.ServiceId{
								Value: mql.String(fmt.Sprintf("%v", s)),
							},
						})
					}

					return
				}(),
			},
		},
	}

	return a
}
