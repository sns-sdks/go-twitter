package twitter

// GetQuoteTweetsOpts specifies the parameters for get quote tweets
type GetQuoteTweetsOpts struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	TweetOpts
}

// GetQuoteTweets Returns Quote Tweets for a Tweet specified by the requested Tweet ID.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/quote-tweets/api-reference/get-tweets-id-quote_tweets
func (r *TweetResource) GetQuoteTweets(id string, args GetQuoteTweetsOpts) (*TweetsResp, *APIError) {
	path := "/tweets/" + id + "/quote_tweets"

	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
