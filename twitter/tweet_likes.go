package twitter

// GetLikingUsers  Return information about a Tweet’s liking users.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/get-tweets-id-liking_users
func (r *TweetResource) GetLikingUsers(id string, args UserOpts) (*UsersResp, *APIError) {
	path := Baseurl + "/tweets/" + id + "/liking_users"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LikedTweetsOpts specifies the parameters for get liked tweets
type LikedTweetsOpts struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	TweetOpts
}

// GetLikedTweets Return information about a user’s liked Tweets.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/get-users-id-liked_tweets
func (r *TweetResource) GetLikedTweets(id string, args LikedTweetsOpts) (*TweetsResp, *APIError) {
	path := Baseurl + "/users/" + id + "/liked_tweets"

	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
