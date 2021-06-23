package twitter

import ent "go-twitter/twitter/entities"

/*
	Tweets include api for tweets
*/

type TweetResource Resource

func newTweetResource(cli *Client) *TweetResource {
	return &TweetResource{Cli: cli}
}

type TweetCommonParams struct {
	TweetFields string `url:"tweet.fields,omitempty"`
	Expansions  string `url:"expansions,omitempty"`
	MediaFields string `url:"media.fields,omitempty"`
	PlaceFields string `url:"place.fields,omitempty"`
	PollFields  string `url:"poll.fields,omitempty"`
	UserFields  string `url:"user.fields,omitempty"`
}

type TweetParams struct {
	IDs string `url:"ids,omitempty"`
	*TweetCommonParams
}

type TimelineParams struct {
	Exclude         string `url:"exclude,omitempty"`
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	SinceID         string `url:"since_id,omitempty"`
	UntilID         string `url:"until_id,omitempty"`
	StartTime       string `url:"start_time,omitempty"`
	EndTime         string `url:"end_time,omitempty"`
	*TweetCommonParams
}

type MentionParams struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	SinceID         string `url:"since_id,omitempty"`
	UntilID         string `url:"until_id,omitempty"`
	StartTime       string `url:"start_time,omitempty"`
	EndTime         string `url:"end_time,omitempty"`
	*TweetCommonParams
}

func (r *TweetResource) LookupByID(id string, params TweetParams) (*ent.TweetResp, *APIError) {
	path := Baseurl + "/tweets/" + id

	resp := new(ent.TweetResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) LookupByIDs(params TweetParams) (*ent.TweetsResp, *APIError) {
	path := Baseurl + "/tweets"

	resp := new(ent.TweetsResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) GetTimelines(id string, params TimelineParams) (*ent.TweetsResp, *APIError) {
	path := Baseurl + "/users/" + id + "/tweets"

	resp := new(ent.TweetsResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) GetMentions(id string, params MentionParams) (*ent.TweetsResp, *APIError) {
	path := Baseurl + "/users/" + id + "/mentions"

	resp := new(ent.TweetsResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
