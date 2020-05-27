package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/metamatex/metamate/reddit-svc/pkg/communication"
	"github.com/metamatex/metamate/reddit-svc/pkg/types"
	"net/http"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() (err error) {
	c, err := communication.NewClient(communication.ClientOpts{
		Credentials: types.Credentials{
			ClientId:     "5Spu4UHEEVbvsQ",
			ClientSecret: "OzXCIrbPZVbTlgy37YeZDfCiYWQ",
			Username:     "metamatex",
			Password:     "vJ6g3ouQbZ4ztiA",
		},
		UserAgent: "abc",
		Client:    &http.Client{},
	})
	if err != nil {
		return
	}

	err = c.Authenticate()
	if err != nil {
		return
	}

	l := 1
	time := "all"
	after := "t3_gkv5p9"
	rsp, err := c.GetSubredditSubmissions("graphql", "new", &l, &time, &after)
	if err != nil {
		return
	}
	spew.Dump(rsp)

	//rsp0, err := c.GetSubredditAbout("graphql")
	//if err != nil {
	//	return
	//}
	//spew.Dump(rsp0)

	rsp0, err := c.GetUserAbout("maniishjaiin")
	if err != nil {
		return
	}
	spew.Dump(rsp0)

	//l := 100
	//time := "all"
	//rsp, err := c.GetUserSubmissions("maniishjaiin", "new", &l, &time, nil)
	//if err != nil {
	//	return
	//}
	//spew.Dump(rsp.Data.Children)

	return
}

func runMira() (err error) {
	//c := mira.Credentials{
	//	ClientId:     "5Spu4UHEEVbvsQ",
	//	ClientSecret: "OzXCIrbPZVbTlgy37YeZDfCiYWQ",
	//	Username:     "metamatex",
	//	Password:     "vJ6g3ouQbZ4ztiA",
	//	UserAgent:    "abc",
	//}
	//
	//r, err := mira.Init(c)
	//if err != nil {
	//	return
	//}

	//sort := "top"
	//var limit int = 100
	//duration := "all"
	//_, err = r.Subreddit("graphql").Submissions(sort, duration, limit)
	//if err != nil {
	//	return
	//}

	//sort := "top"
	//limit := 100
	//duration := "all"
	//i, err := r.Redditor("maniishjaiin").Submissions(sort, duration, limit)
	//if err != nil {
	//    return
	//}

	return
}
