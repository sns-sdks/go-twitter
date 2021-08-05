package twitter

// Media refers to any image, GIF, or video attached to a Tweet. The media object is not a primary object on any endpoint, but can be found and expanded in the Tweet object.
// Refer: https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/media
type Media struct {
	MediaKey         *string                `json:"media_key,omitempty"`
	Type             *string                `json:"type,omitempty"`
	URL              *string                `json:"url,omitempty"`
	DurationMS       *int                   `json:"duration_ms,omitempty"`
	Height           *int                   `json:"height,omitempty"`
	Width            *int                   `json:"width,omitempty"`
	NonPublicMetrics *MediaNonPublicMetrics `json:"non_public_metrics,omitempty"`
	OrganicMetrics   *MediaOrganicMetrics   `json:"organic_metrics,omitempty"`
	PreviewImageURL  *string                `json:"preview_image_url,omitempty"`
	PromotedMetrics  *MediaPromotedMetrics  `json:"promoted_metrics,omitempty"`
	PublicMetrics    *MediaNonPublicMetrics `json:"public_metrics,omitempty"`
}

type MediaNonPublicMetrics struct {
	Playback0Count   *int `json:"playback_0_count,omitempty"`
	Playback100Count *int `json:"playback_100_count,omitempty"`
	Playback25Count  *int `json:"playback_25_count,omitempty"`
	Playback50Count  *int `json:"playback_50_count,omitempty"`
	Playback75Count  *int `json:"playback_75_count,omitempty"`
}
type MediaOrganicMetrics struct {
	Playback0Count   *int `json:"playback_0_count,omitempty"`
	Playback100Count *int `json:"playback_100_count,omitempty"`
	Playback25Count  *int `json:"playback_25_count,omitempty"`
	Playback50Count  *int `json:"playback_50_count,omitempty"`
	Playback75Count  *int `json:"playback_75_count,omitempty"`
	ViewCount        *int `json:"view_count,omitempty"`
}

type MediaPromotedMetrics struct {
	Playback0Count   *int `json:"playback_0_count,omitempty"`
	Playback100Count *int `json:"playback_100_count,omitempty"`
	Playback25Count  *int `json:"playback_25_count,omitempty"`
	Playback50Count  *int `json:"playback_50_count,omitempty"`
	Playback75Count  *int `json:"playback_75_count,omitempty"`
	ViewCount        *int `json:"view_count,omitempty"`
}

type MediaPublicMetrics struct {
	ViewCount *int `json:"view_count,omitempty"`
}
