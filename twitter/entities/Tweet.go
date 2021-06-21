package entities

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
	Entities           *TwitterEntities    `json:"entities,omitempty"`
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
