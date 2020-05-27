package communication

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/metamatex/metamate/reddit-svc/pkg/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	Opts      ClientOpts
	Token     string
	ExpiresIn time.Duration
}

type ClientOpts struct {
	Client      *http.Client
	Credentials types.Credentials
	UserAgent   string
}

func NewClient(opts ClientOpts) (c Client, err error) {
	if opts.Client == nil {
		err = errors.New("opts.Client is nil")

		return
	}

	c = Client{
		Opts: opts,
	}

	return
}

const RedditBase = "https://www.reddit.com/"

func (c *Client) Authenticate() (err error) {
	authUrl := RedditBase + "api/v1/access_token"

	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("username", c.Opts.Credentials.Username)
	form.Add("password", c.Opts.Credentials.Password)

	raw := c.Opts.Credentials.ClientId + ":" + c.Opts.Credentials.ClientSecret
	encoded := base64.StdEncoding.EncodeToString([]byte(raw))

	r, err := http.NewRequest("POST", authUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return
	}

	r.Header.Set("User-Agent", c.Opts.UserAgent)
	r.Header.Set("Authorization", "Basic "+encoded)

	response, err := c.Opts.Client.Do(r)
	if err != nil {
		return
	}
	defer response.Body.Close()

	var rsp types.AccessTokenResponse
	err = json.NewDecoder(response.Body).Decode(&rsp)
	if err != nil {
		return
	}

	err = types.GetError(rsp.Error)
	if err != nil {
		return
	}

	c.Token = rsp.AccessToken
	c.ExpiresIn = time.Duration(rsp.ExpiresIn) * time.Second

	return
}

func (c *Client) GetSubredditSubmissions(name string, order string, limit *int, time *string, after *string) (rsp types.GetSubredditSubmissionsResponse, err error) {
	u, err := url.Parse(fmt.Sprintf("https://oauth.reddit.com/r/%v/%v.json", name, order))
	if err != nil {
		return
	}

	q := url.Values{}
	if limit != nil {
		q.Set("limit", fmt.Sprintf("%v", *limit))
	}

	if time != nil {
		q.Set("t", fmt.Sprintf("%v", *time))
	}

	if after != nil {
		q.Set("after", fmt.Sprintf("%v", *after))
	}

	u.RawQuery = q.Encode()

	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return
	}

	r.Header.Set("User-Agent", c.Opts.UserAgent)
	r.Header.Set("Authorization", "Bearer "+c.Token)

	rsp0, err := c.Opts.Client.Do(r)
	if err != nil {
		return
	}
	defer rsp0.Body.Close()

	b, err := ioutil.ReadAll(rsp0.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &rsp)
	//if err != nil {
	//    return
	//}

	err = types.GetError(rsp.Error)
	if err != nil {
		return
	}

	return
}

func (c *Client) GetSubredditAbout(name string) (rsp types.GetSubredditAboutResponse, err error) {
	u, err := url.Parse(fmt.Sprintf("https://oauth.reddit.com/r/%v/about", name))
	if err != nil {
		return
	}

	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return
	}

	r.Header.Set("User-Agent", c.Opts.UserAgent)
	r.Header.Set("Authorization", "Bearer "+c.Token)

	rsp0, err := c.Opts.Client.Do(r)
	if err != nil {
		return
	}
	defer rsp0.Body.Close()

	b, err := ioutil.ReadAll(rsp0.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &rsp)
	//if err != nil {
	//	return
	//}

	return
}

func (c *Client) GetUserAbout(name string) (rsp types.GetUserAboutResponse, err error) {
	u, err := url.Parse(fmt.Sprintf("https://oauth.reddit.com/user/%v/about", name))
	if err != nil {
		return
	}

	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return
	}

	r.Header.Set("User-Agent", c.Opts.UserAgent)
	r.Header.Set("Authorization", "Bearer "+c.Token)

	rsp0, err := c.Opts.Client.Do(r)
	if err != nil {
		return
	}
	defer rsp0.Body.Close()

	b, err := ioutil.ReadAll(rsp0.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &rsp)
	//if err != nil {
	//	return
	//}

	err = types.GetError(rsp.Error)
	if err != nil {
		return
	}

	return
}

func (c *Client) GetUserSubmissions(name string, order string, limit *int, time *string, after *string) (rsp types.GetSubredditSubmissionsResponse, err error) {
	u, err := url.Parse(fmt.Sprintf("https://oauth.reddit.com/u/%v/submitted/%v.json", name, order))
	if err != nil {
		return
	}

	q := url.Values{}
	if limit != nil {
		q.Set("limit", fmt.Sprintf("%v", *limit))
	}

	if time != nil {
		q.Set("t", fmt.Sprintf("%v", *time))
	}

	if after != nil {
		q.Set("after", fmt.Sprintf("%v", *after))
	}

	q.Set("show", "all")

	u.RawQuery = q.Encode()

	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return
	}

	r.Header.Set("User-Agent", c.Opts.UserAgent)
	r.Header.Set("Authorization", "Bearer "+c.Token)

	rsp0, err := c.Opts.Client.Do(r)
	if err != nil {
		return
	}
	defer rsp0.Body.Close()

	b, err := ioutil.ReadAll(rsp0.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &rsp)
	//if err != nil {
	//	return
	//}

	err = types.GetError(rsp.Error)
	if err != nil {
		return
	}

	return
}
