package twitter

// GetRetweetedBy Return information about who has Retweeted a Tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/get-tweets-id-retweeted_by
func (r *TweetResource) GetRetweetedBy(id string, args UserOpts) (*UsersResp, *APIError) {
	path := "/tweets/" + id + "/retweeted_by"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// retweetTweetOpts specifies the parameters for retweet tweet
type retweetTweetOpts struct {
	TweetID string `json:"tweet_id"`
}

// RetweetedStatus represents the status for retweet tweet
type RetweetedStatus struct {
	Retweeted *bool `json:"retweeted,omitempty"`
}

func (r RetweetedStatus) String() string {
	return Stringify(r)
}

// RetweetedResp represents the response for retweet tweet
type RetweetedResp struct {
	Data *RetweetedStatus `json:"data,omitempty"`
}

func (r RetweetedResp) String() string {
	return Stringify(r)
}

// RetweetCreate Allows an authenticated user ID to Retweet the target Tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/post-users-id-retweets
func (r *TweetResource) RetweetCreate(id, tweetID string) (*RetweetedResp, *APIError) {
	path := "/users/" + id + "/retweets"
	postArgs := retweetTweetOpts{TweetID: tweetID}

	resp := new(RetweetedResp)
	err := r.Cli.DoPost(path, postArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RetweetDestroy Allows an authenticated user ID to remove the Retweet of a Tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/delete-users-id-retweets-tweet_id
func (r *TweetResource) RetweetDestroy(id, tweetID string) (*RetweetedResp, *APIError) {
	path := "/users/" + id + "/retweets/" + tweetID

	resp := new(RetweetedResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
