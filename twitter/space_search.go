package twitter

// SpaceSearchOpts specifies the parameters for search spaces
type SpaceSearchOpts struct {
	Query      string `url:"query"`
	State      string `url:"state"`
	MaxResults int    `url:"max_results,omitempty"`
	SpaceOpts
}

// Search
// Refer: https://developer.twitter.com/en/docs/twitter-api/spaces/search/api-reference/get-spaces-search
func (r *SpaceResource) Search(args SpaceSearchOpts) (*SpacesResp, *APIError) {
	path := "/spaces/search"
	resp := new(SpacesResp)

	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
