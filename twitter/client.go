package twitter

import "github.com/go-resty/resty/v2"

const BASEURL = "https://api.twitter.com/2"

type Client struct {
	cli *resty.Client
	// API Resource
	Users *UserResource
}

func NewBearerClient(bearerToken string) *Client {
	rCli := resty.New()
	rCli.SetAuthToken(bearerToken)

	return &Client{
		cli:   rCli,
		Users: newUserResource(rCli),
	}
}
