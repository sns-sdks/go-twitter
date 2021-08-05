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
	ID                 *string             `json:"id,omitempty"`
	Text               *string             `json:"text,omitempty"`
	Attachments        *Attachments        `json:"attachments,omitempty"`
	AuthorID           *string             `json:"author_id,omitempty"`
	ContextAnnotations *ContextAnnotation  `json:"context_annotations,omitempty"`
	ConversationID     *string             `json:"conversation_id,omitempty"`
	CreatedAt          *string             `json:"created_at,omitempty"`
	Entities           *TweetEntities      `json:"entities,omitempty"`
	Geo                *TweetGeo           `json:"geo,omitempty"`
	InReplyToUserID    *string             `json:"in_reply_to_user_id,omitempty"`
	Lang               *string             `json:"lang,omitempty"`
	NonPublicMetrics   *NonPublicMetrics   `json:"non_public_metrics,omitempty"`
	OrganicMetrics     *OrganicMetrics     `json:"organic_metrics,omitempty"`
	PossiblySensitive  *bool               `json:"possibly_sensitive,omitempty"`
	PromotedMetrics    *PromotedMetrics    `json:"promoted_metrics,omitempty"`
	PublicMetrics      *TweetPublicMetrics `json:"public_metrics,omitempty"`
	ReferencedTweets   []*ReferencedTweet  `json:"referenced_tweets,omitempty"`
	ReplySettings      *string             `json:"reply_settings,omitempty"`
	Source             *string             `json:"source,omitempty"`
	Withheld           *TweetWithheld      `json:"withheld,omitempty"`
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
	RetweetCount *int `json:"retweet_count,omitempty"`
	ReplyCount   *int `json:"reply_count,omitempty"`
	LikeCount    *int `json:"like_count,omitempty"`
	QuoteCount   *int `json:"quote_count,omitempty"`
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

type TweetOpts struct {
	Tweet      string `url:"tweet.fields,omitempty"`
	Expansions string `url:"expansions,omitempty"`
	Media      string `url:"media.fields,omitempty"`
	Place      string `url:"place.fields,omitempty"`
	Poll       string `url:"poll.fields,omitempty"`
	User       string `url:"user.fields,omitempty"`
}

func (r *TweetResource) LookupByID(id string, args TweetOpts) (*TweetResp, *APIError) {
	path := Baseurl + "/tweets/" + id

	resp := new(TweetResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type tweetOptsByIDs struct {
	IDs string `url:"ids,omitempty"`
	TweetOpts
}

func (r *TweetResource) LookupByIDs(ids string, args TweetOpts) (*TweetsResp, *APIError) {
	path := Baseurl + "/tweets"

	newArgs := tweetOptsByIDs{ids, args}
	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, newArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
