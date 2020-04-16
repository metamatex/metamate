package angolia

import (
	"encoding/json"
	"fmt"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"
	"net/http"
	"net/url"
)

func GetPostsSearch(c *http.Client, req sdk.GetPostsRequest) (ss []sdk.Post, errs []error) {
	err := func() (err error) {
		var u string

		u = fmt.Sprintf("http://hn.algolia.com/api/v1/search?query=%v", url.QueryEscape(*req.Mode.Search.Term))

		rsp, err := c.Get(u)
		if err != nil {
			return
		}
		defer rsp.Body.Close()

		var r struct {
			Hits []searchHnStory
		}
		err = json.NewDecoder(rsp.Body).Decode(&r)
		if err != nil {
			return
		}

		for _, s := range r.Hits {
			ss = append(ss, mapSearchHNStoryToPost(s))
		}

		return
	}()
	if err != nil {
		errs = append(errs, err)
	}

	return
}

type searchHnStory struct {
	CreatedAt      *string `json:"created_at"`
	Title          *string
	Url            *string
	Author         *string
	Points         *int
	StoryText      *string `json:"story_text"`
	CommentText    *string `json:"comment_text"`
	NumComments    *int    `json:"num_comments"`
	StoryId        *int    `json:"story_id"`
	StoryTitle     *string `json:"story_title"`
	StoryUrl       *string `json:"story_url"`
	ParentId       *int    `json:"parent_id"`
	CreateAtI      *int    `json:"created_at_i"`
	RelevanceScore *int    `json:"relevance_score"`
	ObjectId       *string `json:"objectID"`
}

func mapSearchHNStoryToPost(s searchHnStory) (p sdk.Post) {
	//type searchHnStory struct {
	//	CreatedAt      *string `json:"created_at"`
	//	Title          *string
	//	Url            *string
	//	Author         *string
	//	Points         *int
	//	StoryText      *string `json:"story_text"`
	//	CommentText    *string `json:"comment_text"`
	//	NumComments    *int    `json:"num_comments"`
	//	StoryId        *int    `json:"story_id"`
	//	StoryTitle     *string `json:"story_title"`
	//	StoryUrl       *string `json:"story_url"`
	//	ParentId       *int    `json:"parent_id"`
	//	CreateAtI      *int    `json:"created_at_i"`
	//	RelevanceScore *int    `json:"relevance_score"`
	//	x ObjectId       *string `json:"objectID"`
	//}

	return sdk.Post{
		Id: &sdk.ServiceId{
			Value: s.ObjectId,
		},
	}
}
