package pkg

import (
	"context"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
)

const (
	TIMELINE_PUBLIC       = "public"
	TIMELINE_PUBLIC_LOCAL = "public_local"
	TIMELINE_HOME         = "home"
	TIMELINE_MEDIA        = "media"
	TIMELINE_MEDIA_LOCAL  = "media_local"
)

func getFeedsCollection(ctx context.Context, req sdk.GetFeedsRequest) (rsp sdk.GetFeedsResponse) {
	feeds := []sdk.Feed{
		{
			Id: &sdk.ServiceId{
				Value: ptr.String(TIMELINE_PUBLIC),
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: ptr.String(TIMELINE_PUBLIC_LOCAL),
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: ptr.String(TIMELINE_HOME),
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: ptr.String(TIMELINE_MEDIA),
			},
		},
		{
			Id: &sdk.ServiceId{
				Value: ptr.String(TIMELINE_MEDIA_LOCAL),
			},
		},
	}

	rsp.Feeds = feeds

	return
}
