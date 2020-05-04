package static

import (
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"
	"github.com/metamatex/metamate/hackernews-svc/pkg/types"
)

func GetPostFeedsCollection() (fs []sdk.PostFeed, errs []sdk.Error) {
	fs = []sdk.PostFeed{
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(types.TopStories),
			},
			Info: &sdk.Info{
				Name: &sdk.Text{
					Value:      sdk.String("Top stories"),
					Formatting: &sdk.FormattingKind.Plain,
				},
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(types.NewStories),
			},
			Info: &sdk.Info{
				Name: &sdk.Text{
					Value:      sdk.String("New stories"),
					Formatting: &sdk.FormattingKind.Plain,
				},
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(types.BestStories),
			},
			Info: &sdk.Info{
				Name: &sdk.Text{
					Value:      sdk.String("Best stories"),
					Formatting: &sdk.FormattingKind.Plain,
				},
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(types.AskStories),
			},
			Info: &sdk.Info{
				Name: &sdk.Text{
					Value:      sdk.String("Ask stories"),
					Formatting: &sdk.FormattingKind.Plain,
				},
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(types.ShowStories),
			},
			Info: &sdk.Info{
				Name: &sdk.Text{
					Value:      sdk.String("Show stories"),
					Formatting: &sdk.FormattingKind.Plain,
				},
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(types.JobStories),
			},
			Info: &sdk.Info{
				Name: &sdk.Text{
					Value:      sdk.String("Job stories"),
					Formatting: &sdk.FormattingKind.Plain,
				},
			},
		},
	}

	return
}
