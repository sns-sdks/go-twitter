package twitter

// CreateTweetGeoOpts A JSON object that contains location information for a Tweet. You can only add a location to Tweets if you have geo enabled in your profile settings. If you don't have geo enabled, you can still add a location parameter in your request body, but it won't get attached to your Tweet.
type CreateTweetGeoOpts struct {
	PlaceID string `json:"place_id,omitempty"`
}

// CreateTweetMediaOpts A JSON object that contains media information being attached to created Tweet. This is mutually exclusive from Quote Tweet ID and Poll.
type CreateTweetMediaOpts struct {
	MediaIDs      []string `json:"media_ids,omitempty"`
	TaggedUserIDs []string `json:"tagged_user_ids,omitempty"`
}

// CreateTweetPollOpts A JSON object that contains options for a Tweet with a poll. This is mutually exclusive from Media and Quote Tweet ID.
type CreateTweetPollOpts struct {
	DurationMinutes int      `json:"duration_minutes,omitempty"`
	Options         []string `json:"options,omitempty"`
}

// CreateTweetReplyOpts A JSON object that contains information of the Tweet being replied to.
type CreateTweetReplyOpts struct {
	ExcludeReplyUserIDs []string `json:"exclude_reply_user_ids,omitempty"`
	InReplyToTweetID    string   `json:"in_reply_to_tweet_id,omitempty"`
}

// CreateTweetOpts parameters for create a Tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/post-tweets
type CreateTweetOpts struct {
	DirectMessageDeepLink string                `json:"direct_message_deep_link,omitempty"`
	ForSuperFollowersOnly bool                  `json:"for_super_followers_only,omitempty"`
	Geo                   *CreateTweetGeoOpts   `json:"geo,omitempty"`
	Media                 *CreateTweetMediaOpts `json:"media,omitempty"`
	Poll                  *CreateTweetPollOpts  `json:"poll,omitempty"`
	QuoteTweetID          string                `json:"quote_tweet_id,omitempty"`
	Reply                 *CreateTweetReplyOpts `json:"reply,omitempty"`
	ReplySettings         string                `json:"reply_settings,omitempty"`
	Text                  string                `json:"text,omitempty"`
}

// TweetCreate Creates a Tweet on behalf of an authenticated user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/post-tweets
func (r *TweetResource) TweetCreate(args CreateTweetOpts) (*TweetResp, *APIError) {
	path := "/tweets"

	resp := new(TweetResp)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TweetDeletedStatus represents the status for tweet delete.
type TweetDeletedStatus struct {
	Deleted *bool `json:"deleted,omitempty"`
}

func (t TweetDeletedStatus) String() string {
	return Stringify(t)
}

// TweetDeletedResp represents the response for delete tweet
type TweetDeletedResp struct {
	Data *TweetDeletedStatus `json:"data,omitempty"`
}

func (t TweetDeletedResp) String() string {
	return Stringify(t)
}

// TweetRemove Allows a user or authenticated user ID to delete a Tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id
func (r *TweetResource) TweetRemove(id string) (*TweetDeletedResp, *APIError) {
	path := "/tweets/" + id

	resp := new(TweetDeletedResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
