package twitter

// TweetsCounts represents the struct for counts for tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference
type TweetsCounts struct {
	End        *string `json:"end,omitempty"`
	Start      *string `json:"start,omitempty"`
	TweetCount *int    `json:"tweet_count,omitempty"`
}

func (t TweetsCounts) String() string {
	return Stringify(t)
}

// TweetCountsOpts specifies the parameters for get tweet counts.
type TweetCountsOpts struct {
	Query       string `url:"query"`
	Granularity string `url:"granularity,omitempty"`
	StartTime   string `url:"start_time,omitempty"`
	EndTime     string `url:"end_time,omitempty"`
	SinceID     string `url:"since_id,omitempty"`
	UntilID     string `url:"until_id,omitempty"`
}

// CountsRecent Returns count of Tweets from the last seven days that match a search query.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-recent
func (r *TweetResource) CountsRecent(args TweetCountsOpts) (*TweetsCountsResp, *APIError) {
	path := Baseurl + "/tweets/counts/recent"

	resp := new(TweetsCountsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CountsALL Returns the complete history of public Tweets matching a search query
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-all
func (r *TweetResource) CountsALL(args TweetCountsOpts) (*TweetsCountsResp, *APIError) {
	path := Baseurl + "/tweets/counts/all"

	resp := new(TweetsCountsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
