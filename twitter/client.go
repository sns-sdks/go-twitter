package twitter

import "github.com/go-resty/resty/v2"

const (
	Baseurl = "https://api.twitter.com/2"
	HttpGet = resty.MethodGet
)

type Client struct {
	Cli *resty.Client
	// API Resource
	Users  *UserResource
	Tweets *TweetResource
}

type Resource struct {
	Cli *Client
}

func NewBearerClient(bearerToken string) *Client {
	rCli := resty.New()
	rCli.SetAuthToken(bearerToken)

	c := &Client{Cli: rCli}
	c.Users = newUserResource(c)
	c.Tweets = newTweetResource(c)

	return c
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
