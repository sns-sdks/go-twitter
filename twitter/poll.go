package twitter

// Poll included in a Tweet is not a primary object on any endpoint, but can be found and expanded in the Tweet object.
// Refer: https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/poll
type Poll struct {
	ID              *string     `json:"id,omitempty"`
	Options         *PollOption `json:"options,omitempty"`
	DurationMinutes *int        `json:"duration_minutes,omitempty"`
	EndDatetime     *string     `json:"end_datetime,omitempty"`
	VotingStatus    *string     `json:"voting_status,omitempty"`
}

type PollOption struct {
	Position *int    `json:"position,omitempty"`
	Label    *string `json:"label,omitempty"`
	Votes    *int    `json:"votes,omitempty"`
}

func (p Poll) String() string {
	return Stringify(p)
}
