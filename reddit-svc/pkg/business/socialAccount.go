package business

import (
	"fmt"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/reddit-svc/pkg/communication"
	"github.com/metamatex/metamate/reddit-svc/pkg/types"
)

func GetSocialAccountId(c communication.Client, req mql.GetSocialAccountsRequest) (as []mql.SocialAccount, errs []mql.Error) {
	switch *req.Mode.Id.Kind {
	case mql.IdKind.ServiceId:
		rsp, err := c.GetUserAbout(*req.Mode.Id.ServiceId.Value)
		if err != nil {
			errs = append(errs, mql.Error{
				Message: mql.String(err.Error()),
			})

			return
		}

		err = types.GetError(rsp.Error)
		if err != nil {
			errs = append(errs, mql.Error{
				Message: mql.String(err.Error()),
			})

			return
		}

		as = append(as, mapUserDataToSocialAccount(*rsp.Data))
	case mql.IdKind.Username:
		rsp, err := c.GetUserAbout(*req.Mode.Id.Username)
		if err != nil {
			errs = append(errs, mql.Error{
				Message: mql.String(err.Error()),
			})

			return
		}

		err = types.GetError(rsp.Error)
		if err != nil {
			errs = append(errs, mql.Error{
				Message: mql.String(err.Error()),
			})

			return
		}

		as = append(as, mapUserDataToSocialAccount(*rsp.Data))
	}

	return
}

func mapUserDataToSocialAccount(d types.UserData) mql.SocialAccount {
	return mql.SocialAccount{
		Id: &mql.ServiceId{
			Value: d.Name,
		},
		AlternativeIds: []mql.Id{
			{
				Kind:     &mql.IdKind.Username,
				Username: d.Name,
			},
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String(fmt.Sprintf("https://reddit.com%v", *d.Subreddit.Url)),
				},
			},
		},
		Points: func() *int32 {
			var n int32 = 0
			if d.CommentKarma != nil {
				n += *d.CommentKarma
			}
			if d.LinkKarma != nil {
				n += *d.LinkKarma
			}

			return &n
		}(),
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
		Username: d.Name,
	}
}
