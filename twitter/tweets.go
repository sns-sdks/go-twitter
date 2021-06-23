package twitter

import ent "go-twitter/twitter/entities"

/*
	Tweets include api for tweets
*/

type TweetResource Resource

func newTweetResource(cli *Client) *TweetResource {
	return &TweetResource{Cli: cli}
}

type TweetParams struct {
	IDs         string `url:"ids,omitempty"`
	TweetFields string `url:"tweet.fields,omitempty"`
	Expansions  string `url:"expansions,omitempty"`
	MediaFields string `url:"media.fields,omitempty"`
	PlaceFields string `url:"place.fields,omitempty"`
	PollFields  string `url:"poll.fields,omitempty"`
	UserFields  string `url:"user.fields,omitempty"`
}

func (r *TweetResource) LookupByID(id string, params TweetParams) (*ent.TweetResp, *APIError) {
	path := BASEURL + "/tweets/" + id

	resp := new(ent.TweetResp)
	err := r.Cli.Do(HTTP_GET, path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
