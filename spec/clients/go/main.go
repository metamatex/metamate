package main

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/mql"
)

func main() {
	c := mql.NewClient(mql.ClientOpts{Host: "https://metamate.one"})

	rsp, err := c.GetPosts(context.Background(), mql.GetPostsRequest{
		ServiceFilter: &mql.ServiceFilter{
			Id: &mql.ServiceIdFilter{
				Value: &mql.StringFilter{
					Is: mql.String("hackernews"),
				},
			},
		},
		Mode: &mql.GetMode{
			Kind: &mql.GetModeKind.Search,
			Search: &mql.SearchGetMode{
				Term: mql.String("book recommendations"),
			},
		},
	})
	if err != nil {
		panic(err)
	}

	println(rsp.Posts)
}
