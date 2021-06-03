package entities

// Entities represent metadata and context info parsed from Twitter components.
// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/user

type PublicMetrics struct {
	FollowerCount  int64 `json:"follower_count"`
	FollowingCount int64 `json:"following_count"`
	TweetCount     int   `json:"tweet_count"`
	ListCount      int   `json:"list_count"`
}

// User represent a Twitter user.
type User struct {
	ID              int64         `json:"id"`
	Name            string        `json:"name"`
	Username        string        `json:"username"`
	CreatedAt       string        `json:"created_at"`
	Description     string        `json:"description"`
	Location        string        `json:"location"`
	PinnedTweetID   string        `json:"pinned_tweet_id"`
	ProfileImageUrl string        `json:"profile_image_url"`
	Protected       string        `json:"protected"`
	URL             string        `json:"url"`
	Verified        bool          `json:"verified"`
	Entities        UserEntities  `json:"entities"`
	PublicMetrics   PublicMetrics `json:"public_metrics"`
	WithHeld        WithHeld      `json:"withheld"`
}
