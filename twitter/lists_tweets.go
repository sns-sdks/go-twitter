package twitter

// ListTweetsOpts specifies the parameters for get list tweets
type ListTweetsOpts struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	TweetOpts
}

// GetListTweets Returns a list of Tweets from the specified List.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/list-tweets/api-reference/get-lists-id-tweets
func (r *ListsResource) GetListTweets(id string, args ListTweetsOpts) (*TweetsResp, *APIError) {
	path := "/lists/" + id + "/tweets"

	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
