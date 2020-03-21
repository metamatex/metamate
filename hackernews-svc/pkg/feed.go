package pkg

import (
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"

	"net/http"
)

func getFeedsCollection(c *http.Client, req sdk.GetFeedsRequest) (fs []sdk.Feed, errs []error) {
	fs = []sdk.Feed{
		{
			Id: &sdk.ServiceId{
				Value: sdk.String("topstories"),
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
				Value: sdk.String("newstories"),
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
				Value: sdk.String("beststories"),
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
				Value: sdk.String("askstories"),
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
				Value: sdk.String("showstories"),
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
				Value: sdk.String("jobstories"),
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
