package twitter

// Entities are JSON objects that provide additional information about hashtags, urls, user mentions, and cashtags associated with the description.
type Entities struct {
	Urls     []*EntityUrl     `json:"urls"`
	HashTags []*EntityHashTag `json:"Hashtags"`
	Mentions []*EntityMention `json:"mentions"`
	CashTags []*EntityCashTag `json:"cashtags"`
}

type UserEntities struct {
	URL         *Entities `json:"url"`
	Description *Entities `json:"description"`
}

type TweetEntities struct {
	*Entities
	Annotations []*EntityAnnotation `json:"annotations"`
}

type Image struct {
	URL    *string `json:"url"`
	Width  *int    `json:"width"`
	Height *int    `json:"height"`
}

type EntityUrl struct {
	Start       *int     `json:"start"`
	End         *int     `json:"end"`
	URL         *string  `json:"url"`
	ExpandedURL *string  `json:"expanded_url"`
	DisplayURL  *string  `json:"display_url"`
	Images      []*Image `json:"images"`
	Status      *int     `json:"status"`
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	UnwoundURL  *string  `json:"unwound_url"`
}

type EntityHashTag struct {
	Start *int    `json:"start"`
	End   *int    `json:"end"`
	Tag   *string `json:"tag"`
}

type EntityMention struct {
	Start    *int    `json:"start"`
	End      *int    `json:"end"`
	Username *string `json:"username"`
}

type EntityCashTag struct {
	Start *int    `json:"start"`
	End   *int    `json:"end"`
	Tag   *string `json:"tag"`
}

type EntityAnnotation struct {
	Start          *int     `json:"start"`
	End            *int     `json:"end"`
	Probability    *float64 `json:"probability"`
	Type           *string  `json:"type"`
	NormalizedText *string  `json:"normalized_text"`
}
