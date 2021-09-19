package twitter

// GetLikingUsers  Return information about a Tweet’s liking users.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/get-tweets-id-liking_users
func (r *TweetResource) GetLikingUsers(id string, args UserOpts) (*UsersResp, *APIError) {
	path := "/tweets/" + id + "/liking_users"

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
	path := "/users/" + id + "/liked_tweets"

	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// likeTweetOpts specifies the parameters for like tweet
type likeTweetOpts struct {
	TweetID string `json:"tweet_id"`
}

// LikedStatus represents the status for like tweet
type LikedStatus struct {
	Liked *bool `json:"liked,omitempty"`
}

func (s LikedStatus) String() string {
	return Stringify(s)
}

// LikedResp represents the response for like tweet
type LikedResp struct {
	Data *LikedStatus `json:"data,omitempty"`
}

func (r LikedResp) String() string {
	return Stringify(r)
}

// LikeCreate Allows an authenticated user ID to Like the target Tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/post-users-id-likes
func (r *TweetResource) LikeCreate(id, tweetID string) (*LikedResp, *APIError) {
	path := "/users/" + id + "/likes"
	postArgs := likeTweetOpts{TweetID: tweetID}

	resp := new(LikedResp)
	err := r.Cli.DoPost(path, postArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LikeDestroy Allows a user or authenticated user ID to unlike a Tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/delete-users-id-likes-tweet_id
func (r *TweetResource) LikeDestroy(id, tweetID string) (*LikedResp, *APIError) {
	path := "/users/" + id + "/likes/" + tweetID

	resp := new(LikedResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
