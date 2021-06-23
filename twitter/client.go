package twitter

import "github.com/go-resty/resty/v2"

const (
	BASEURL  = "https://api.twitter.com/2"
	HTTP_GET = resty.MethodGet
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
