package twitter

// GetRetweetedBy Return information about who has Retweeted a Tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/get-tweets-id-retweeted_by
func (r *TweetResource) GetRetweetedBy(id string, args UserOpts) (*UsersResp, *APIError) {
	path := Baseurl + "/tweets/" + id + "/retweeted_by"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
