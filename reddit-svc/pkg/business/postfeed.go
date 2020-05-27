package business

import (
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/reddit-svc/pkg/communication"
	"github.com/metamatex/metamate/reddit-svc/pkg/types"
)

func GetPostsFeedId(c communication.Client, req mql.GetPostFeedsRequest) (fs []mql.PostFeed, errs []mql.Error) {
	switch *req.Mode.Id.Kind {
	case mql.IdKind.ServiceId:
		rsp, err := c.GetSubredditAbout(*req.Mode.Id.ServiceId.Value)
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

		fs = append(fs, mapSubredditDataToPostFeed(*rsp.Data))
	case mql.IdKind.Name:
		rsp, err := c.GetSubredditAbout(*req.Mode.Id.Name)
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

		fs = append(fs, mapSubredditDataToPostFeed(*rsp.Data))
	}

	return
}

func mapSubredditDataToPostFeed(d types.SubredditData) (f mql.PostFeed) {
	return mql.PostFeed{
		Id: &mql.ServiceId{
			Value: d.DisplayName,
		},
		AlternativeIds: []mql.Id{
			{
				Kind: &mql.IdKind.Url,
				Url: &mql.Url{
					Value: mql.String("https://reddit.com" + *d.Url),
				},
			},
		},
		Info: &mql.Info{
			Name: &mql.Text{
				Formatting: &mql.FormattingKind.Plain,
				Value:      d.Title,
			},
			Description: &mql.Text{
				Formatting: &mql.FormattingKind.Markdown,
				Value:      d.Description,
			},
		},
	}
}
