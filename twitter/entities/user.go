package entities

// Entities represent metadata and context info parsed from Twitter components.
// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/user

type PublicMetrics struct {
	FollowerCount  *int `json:"follower_count,omitempty"`
	FollowingCount *int `json:"following_count,omitempty"`
	TweetCount     *int `json:"tweet_count,omitempty"`
	ListCount      *int `json:"list_count,omitempty"`
}

// User represent a Twitter user.
type User struct {
	ID              *string        `json:"id,omitempty"`
	Name            *string        `json:"name,omitempty"`
	Username        *string        `json:"username,omitempty"`
	CreatedAt       *string        `json:"created_at,omitempty"`
	Description     *string        `json:"description,omitempty"`
	Location        *string        `json:"location,omitempty"`
	PinnedTweetID   *string        `json:"pinned_tweet_id,omitempty"`
	ProfileImageUrl *string        `json:"profile_image_url,omitempty"`
	Protected       *string        `json:"protected,omitempty"`
	URL             *string        `json:"url,omitempty,omitempty"`
	Verified        *bool          `json:"verified,omitempty"`
	Entities        *UserEntities  `json:"entities,omitempty"`
	PublicMetrics   *PublicMetrics `json:"public_metrics,omitempty"`
	WithHeld        *WithHeld      `json:"withheld,omitempty"`
}
