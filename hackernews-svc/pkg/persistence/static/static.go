package static

import (
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/mql"
	"github.com/metamatex/metamate/hackernews-svc/pkg/types"
)

func GetPostFeedsCollection() (fs []mql.PostFeed, errs []mql.Error) {
	fs = []mql.PostFeed{
		{
			Id: &mql.ServiceId{
				Value: mql.String(types.TopStories),
			},
			Info: &mql.Info{
				Name: &mql.Text{
					Value:      mql.String("Top stories"),
					Formatting: &mql.FormattingKind.Plain,
				},
			},
		},
		{
			Id: &mql.ServiceId{
				Value: mql.String(types.NewStories),
			},
			Info: &mql.Info{
				Name: &mql.Text{
					Value:      mql.String("New stories"),
					Formatting: &mql.FormattingKind.Plain,
				},
			},
		},
		{
			Id: &mql.ServiceId{
				Value: mql.String(types.BestStories),
			},
			Info: &mql.Info{
				Name: &mql.Text{
					Value:      mql.String("Best stories"),
					Formatting: &mql.FormattingKind.Plain,
				},
			},
		},
		{
			Id: &mql.ServiceId{
				Value: mql.String(types.AskStories),
			},
			Info: &mql.Info{
				Name: &mql.Text{
					Value:      mql.String("Ask stories"),
					Formatting: &mql.FormattingKind.Plain,
				},
			},
		},
		{
			Id: &mql.ServiceId{
				Value: mql.String(types.ShowStories),
			},
			Info: &mql.Info{
				Name: &mql.Text{
					Value:      mql.String("Show stories"),
					Formatting: &mql.FormattingKind.Plain,
				},
			},
		},
		{
			Id: &mql.ServiceId{
				Value: mql.String(types.JobStories),
			},
			Info: &mql.Info{
				Name: &mql.Text{
					Value:      mql.String("Job stories"),
					Formatting: &mql.FormattingKind.Plain,
				},
			},
		},
	}

	return
}
