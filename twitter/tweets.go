package twitter

/*
	Tweets include api for tweets
*/

type TweetResource Resource

func newTweetResource(cli *Client) *TweetResource {
	return &TweetResource{Cli: cli}
}

// Tweet are the basic building block of all things Twitter
// Refer: https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/tweet
type Tweet struct {
	ID                  *string             `json:"id,omitempty"`
	Text                *string             `json:"text,omitempty"`
	EditHistoryTweetIDs []*string           `json:"edit_history_tweet_ids"`
	Attachments         *Attachments        `json:"attachments,omitempty"`
	AuthorID            *string             `json:"author_id,omitempty"`
	ContextAnnotations  *ContextAnnotation  `json:"context_annotations,omitempty"`
	ConversationID      *string             `json:"conversation_id,omitempty"`
	CreatedAt           *string             `json:"created_at,omitempty"`
	EditControls        *TweetEditControls  `json:"edit_controls,omitempty"`
	Entities            *TweetEntities      `json:"entities,omitempty"`
	Geo                 *TweetGeo           `json:"geo,omitempty"`
	InReplyToUserID     *string             `json:"in_reply_to_user_id,omitempty"`
	Lang                *string             `json:"lang,omitempty"`
	NonPublicMetrics    *NonPublicMetrics   `json:"non_public_metrics,omitempty"`
	OrganicMetrics      *OrganicMetrics     `json:"organic_metrics,omitempty"`
	PossiblySensitive   *bool               `json:"possibly_sensitive,omitempty"`
	PromotedMetrics     *PromotedMetrics    `json:"promoted_metrics,omitempty"`
	PublicMetrics       *TweetPublicMetrics `json:"public_metrics,omitempty"`
	ReferencedTweets    []*ReferencedTweet  `json:"referenced_tweets,omitempty"`
	ReplySettings       *string             `json:"reply_settings,omitempty"`
	Source              *string             `json:"source,omitempty"`
	Withheld            *TweetWithheld      `json:"withheld,omitempty"`
}

type Attachments struct {
	PollIDs   []*string `json:"poll_ids,omitempty"`
	MediaKeys []*string `json:"media_keys,omitempty"`
}

type ContextAnnotationDomain struct {
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ContextAnnotationEntity struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ContextAnnotation struct {
	Domain *ContextAnnotationDomain `json:"domain,omitempty"`
	Entity *ContextAnnotationEntity `json:"entity,omitempty"`
}

type TweetEditControls struct {
	EditsRemaining *int    `json:"edits_remaining,omitempty"`
	IsEditEligible *bool   `json:"is_edit_eligible,omitempty"`
	EditableUntil  *string `json:"editable_until,omitempty"`
}

type Coordinates struct {
	Type        *string    `json:"type,omitempty"`
	Coordinates []*float64 `json:"coordinates,omitempty"`
}

type TweetGeo struct {
	Coordinates *Coordinates `json:"coordinates,omitempty"`
	PlaceID     *string      `json:"place_id,omitempty"`
}

type NonPublicMetrics struct {
	ImpressionCount   *int `json:"impression_count,omitempty"`
	URLLinkClicks     *int `json:"url_link_clicks,omitempty"`
	UserProfileClicks *int `json:"user_profile_clicks,omitempty"`
}

type OrganicMetrics struct {
	LikeCount    *int `json:"like_count,omitempty"`
	ReplyCount   *int `json:"reply_count,omitempty"`
	RetweetCount *int `json:"retweet_count,omitempty"`
}

type PromotedMetrics struct {
	ImpressionCount   *int `json:"impression_count,omitempty"`
	LikeCount         *int `json:"like_count,omitempty"`
	ReplyCount        *int `json:"reply_count,omitempty"`
	RetweetCount      *int `json:"retweet_count,omitempty"`
	URLLinkClicks     *int `json:"url_link_clicks,omitempty"`
	UserProfileClicks *int `json:"user_profile_clicks,omitempty"`
}

type TweetPublicMetrics struct {
	RetweetCount    *int `json:"retweet_count,omitempty"`
	ReplyCount      *int `json:"reply_count,omitempty"`
	LikeCount       *int `json:"like_count,omitempty"`
	QuoteCount      *int `json:"quote_count,omitempty"`
	ImpressionCount *int `json:"impression_count,omitempty"`
}

type TweetWithheld struct {
	Copyright    *bool     `json:"copyright,omitempty"`
	CountryCodes []*string `json:"country_codes,omitempty"`
}

type ReferencedTweet struct {
	Type *string `json:"type,omitempty"`
	ID   *string `json:"id,omitempty"`
}

func (t Tweet) String() string {
	return Stringify(t)
}

// TweetOpts specifies the parameters for get tweet
type TweetOpts struct {
	TweetFields string `url:"tweet.fields,omitempty"`
	Expansions  string `url:"expansions,omitempty"`
	MediaFields string `url:"media.fields,omitempty"`
	PlaceFields string `url:"place.fields,omitempty"`
	PollFields  string `url:"poll.fields,omitempty"`
	UserFields  string `url:"user.fields,omitempty"`
}

// LookupByID Returns a variety of information about a single Tweet specified by the requested ID.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets-id
func (r *TweetResource) LookupByID(id string, args TweetOpts) (*TweetResp, *APIError) {
	path := "/tweets/" + id

	resp := new(TweetResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// tweetOptsByIDs specifies the parameters for tweets by ids
type tweetOptsByIDs struct {
	IDs string `url:"ids,omitempty"`
	TweetOpts
}

// LookupByIDs Returns a variety of information about the Tweet specified by the requested ID or list of IDs.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func (r *TweetResource) LookupByIDs(ids string, args TweetOpts) (*TweetsResp, *APIError) {
	path := "/tweets"

	newArgs := tweetOptsByIDs{ids, args}
	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, newArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
