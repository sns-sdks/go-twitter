package twitter

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

const (
	Baseurl    = "https://api.twitter.com/2"
	HttpGet    = resty.MethodGet
	HttpPost   = resty.MethodPost
	HttpPut    = resty.MethodPut
	HttpDelete = resty.MethodDelete
)

type Client struct {
	Cli *resty.Client
	// API Resource
	Users         *UserResource
	Tweets        *TweetResource
	Spaces        *SpaceResource
	Lists         *ListsResource
	Compliance    *ComplianceResource
	DirectMessage *DirectMessageResource

	RateLimit *RateLimit
}

type Resource struct {
	Cli *Client
}

func NewClient(client *resty.Client) *Client {
	c := &Client{Cli: client}
	c.RateLimit = newRateLimit()
	c.Users = newUserResource(c)
	c.Tweets = newTweetResource(c)
	c.Spaces = newSpaceResource(c)
	c.Lists = newListsResource(c)
	c.Compliance = newComplianceResource(c)
	c.DirectMessage = newDirectMessageResource(c)
	return c
}

func NewBearerClient(bearerToken string) *Client {
	rCli := resty.New()
	rCli.SetAuthToken(bearerToken)

	return NewClient(rCli)
}

func NewUserClint(hc *http.Client) *Client {
	rCli := resty.NewWithClient(hc)
	return NewClient(rCli)
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
