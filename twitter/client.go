package twitter

import "github.com/go-resty/resty/v2"

const BASEURL = "https://api.twitter.com/2"

type Client struct {
	Cli *resty.Client
	// API Resource
	Users *UserResource
}

func NewBearerClient(bearerToken string) *Client {
	rCli := resty.New()
	rCli.SetAuthToken(bearerToken)

	return &Client{
		Cli:   rCli,
		Users: newUserResource(rCli),
	}
}
