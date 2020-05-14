package pkg

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/mql"
)

const (
	TimelinePublic      = "public"
	TimelinePublicLocal = "public_local"
	TimelineHome        = "home"
	TimelineMedia       = "media"
	TimelineMediaLocal  = "media_local"
)

func getPostFeedsCollection(ctx context.Context, req mql.GetPostFeedsRequest) (rsp mql.GetPostFeedsResponse) {
	feeds := []mql.PostFeed{
		{
			Id: &mql.ServiceId{
				Value: mql.String(TimelinePublic),
			},
		},
		{
			Id: &mql.ServiceId{
				Value: mql.String(TimelinePublicLocal),
			},
		},
		{
			Id: &mql.ServiceId{
				Value: mql.String(TimelineHome),
			},
		},
		{
			Id: &mql.ServiceId{
				Value: mql.String(TimelineMedia),
			},
		},
		{
			Id: &mql.ServiceId{
				Value: mql.String(TimelineMediaLocal),
			},
		},
	}

	rsp.PostFeeds = feeds

	return
}
