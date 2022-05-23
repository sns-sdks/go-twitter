package twitter

// TimelinesOpts specifies the parameters for get timelines
type TimelinesOpts struct {
	Exclude         string `url:"exclude,omitempty"`
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	SinceID         string `url:"since_id,omitempty"`
	UntilID         string `url:"until_id,omitempty"`
	StartTime       string `url:"start_time,omitempty"`
	EndTime         string `url:"end_time,omitempty"`
	TweetOpts
}

// GetTimelines Returns Tweets composed by a single user, specified by the requested user ID.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-tweets
func (r *TweetResource) GetTimelines(id string, args TimelinesOpts) (*TweetsResp, *APIError) {
	path := "/users/" + id + "/tweets"

	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetTimelinesReverseChronological Allows you to retrieve a collection of the most recent Tweets and Retweets posted by you and users you follow
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-reverse-chronological
func (r *TweetResource) GetTimelinesReverseChronological(id string, args TimelinesOpts) (*TweetsResp, *APIError) {
	path := "/users/" + id + "/timelines/reverse_chronological"

	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MentionsOpts specifies the parameters for get mentions
type MentionsOpts struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	SinceID         string `url:"since_id,omitempty"`
	UntilID         string `url:"until_id,omitempty"`
	StartTime       string `url:"start_time,omitempty"`
	EndTime         string `url:"end_time,omitempty"`
	TweetOpts
}

// GetMentions Returns Tweets mentioning a single user specified by the requested user ID
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-mentions
func (r *TweetResource) GetMentions(id string, args MentionsOpts) (*TweetsResp, *APIError) {
	path := "/users/" + id + "/mentions"

	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
