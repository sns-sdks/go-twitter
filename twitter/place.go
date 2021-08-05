package twitter

// Place tagged in a Tweet is not a primary object on any endpoint, but can be found and expanded in the Tweet resource.
// Refer: https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/place
type Place struct {
	FullName *string `json:"full_name,omitempty"`
	ID       *string `json:"id,omitempty"`
	// ContainedWithin []*Place `json:"contained_within,omitempty"`
	Country     *string   `json:"country,omitempty"`
	CountryCode *string   `json:"country_code,omitempty"`
	Geo         *PlaceGeo `json:"geo,omitempty"`
	Name        *string   `json:"name,omitempty"`
	PlaceType   *string   `json:"place_type,omitempty"`
}

type PlaceGeo struct {
	Type       *string           `json:"type,omitempty"`
	BBox       []*float64        `json:"bbox,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

func (p Place) String() string {
	return Stringify(p)
}
