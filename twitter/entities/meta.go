package entities

// Meta is Response Data for request.
type Meta struct {
	ResultCount     *int         `json:"result_count,omitempty"`
	PreviousToken   *string      `json:"previous_token,omitempty"`
	NextToken       *string      `json:"next_token,omitempty"`
	OldestID        *string      `json:"oldest_id,omitempty"`
	NewestID        *string      `json:"newest_id,omitempty"`
	Sent            *string      `json:"sent,omitempty"`
	Summary         *MetaSummary `json:"summary,omitempty"`
	TotalTweetCount *int         `json:"total_tweet_count,omitempty"`
}

type MetaSummary struct {
	Created    *int `json:"created,omitempty"`
	NotCreated *int `json:"not_created,omitempty"`
	Valid      *int `json:"valid,omitempty"`
	Invalid    *int `json:"invalid,omitempty"`
	Deleted    *int `json:"deleted,omitempty"`
	NotDeleted *int `json:"not_deleted,omitempty"`
}
