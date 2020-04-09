package pkg

import (
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"

	"net/http"
)

const (
	TopStories = "topstories"
	NewStories = "newstories"
	BestStories = "beststories"
	AskStories = "askstories"
	ShowStories = "showstories"
	JobStories = "jobstories"
)

func getPostFeedsCollection(c *http.Client, req sdk.GetPostFeedsRequest) (fs []sdk.PostFeed, errs []error) {
	fs = []sdk.PostFeed{
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(TopStories),
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
				Value: sdk.String(NewStories),
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
				Value: sdk.String(BestStories),
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
				Value: sdk.String(AskStories),
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
				Value: sdk.String(ShowStories),
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
				Value: sdk.String(JobStories),
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
