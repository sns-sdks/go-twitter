package twitter

// TweetSearchOpts specifies the parameters for search tweets
type TweetSearchOpts struct {
	Query      string `url:"query"`
	StartTime  string `url:"start_time,omitempty"`
	EndTime    string `url:"end_time,omitempty"`
	SinceID    string `url:"since_id,omitempty"`
	UntilID    string `url:"until_id,omitempty"`
	MaxResults int    `url:"max_results,omitempty"`
	NextToken  string `url:"next_token,omitempty"`
	TweetOpts
}

// SearchRecent Returns Tweets from the last seven days that match a search query.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent
func (r *TweetResource) SearchRecent(args TweetSearchOpts) (*TweetsResp, *APIError) {
	path := Baseurl + "/tweets/search/recent"

	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchAll Returns the complete history of public Tweets matching a search query.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-all
func (r *TweetResource) SearchAll(args TweetSearchOpts) (*TweetsResp, *APIError) {
	path := Baseurl + "/tweets/search/all"

	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
