package pkg

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/sdk"
	
)

const (
	TIMELINE_PUBLIC       = "public"
	TIMELINE_PUBLIC_LOCAL = "public_local"
	TIMELINE_HOME         = "home"
	TIMELINE_MEDIA        = "media"
	TIMELINE_MEDIA_LOCAL  = "media_local"
)

func getPostFeedsCollection(ctx context.Context, req sdk.GetPostFeedsRequest) (rsp sdk.GetPostFeedsResponse) {
	feeds := []sdk.PostFeed{
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(TIMELINE_PUBLIC),
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(TIMELINE_PUBLIC_LOCAL),
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(TIMELINE_HOME),
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(TIMELINE_MEDIA),
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: sdk.String(TIMELINE_MEDIA_LOCAL),
			},
		},
	}

	rsp.PostFeeds = feeds

	return
}
