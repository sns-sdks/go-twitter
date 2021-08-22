package twitter

/*
	Spaces API
*/

type SpaceResource Resource

func newSpaceResource(cli *Client) *SpaceResource {
	return &SpaceResource{
		Cli: cli,
	}
}

// Space represent a Twitter space
// Refer: https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/space
type Space struct {
	ID               *string   `json:"id"`
	State            *string   `json:"state"`
	CreatedAt        *string   `json:"created_at,omitempty"`
	HostIDs          []*string `json:"host_ids,omitempty"`
	Lang             *string   `json:"lang,omitempty"`
	IsTicketed       *bool     `json:"is_ticketed,omitempty"`
	InvitedUserIDs   []*string `json:"invited_user_ids,omitempty"`
	ParticipantCount *int      `json:"participant_count"`
	ScheduledStart   *string   `json:"scheduled_start,omitempty"`
	SpeakerIDs       []*string `json:"speaker_ids,omitempty"`
	StartedAt        *string   `json:"started_at,omitempty"`
	Title            *string   `json:"title,omitempty"`
	UpdatedAt        *string   `json:"updated_at,omitempty"`
}

// SpaceOpts specifies the parameters for get space
type SpaceOpts struct {
	SpaceFields string `url:"space.fields,omitempty"`
	Expansions  string `url:"expansions,omitempty"`
	UserFields  string `url:"user.fields,omitempty"`
}

// LookupByID Returns a variety of information about a single Space specified by the requested ID.
// Refer: https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-id
func (r *SpaceResource) LookupByID(id string, args SpaceOpts) (*SpaceResp, *APIError) {
	path := Baseurl + "/spaces/" + id
	resp := new(SpaceResp)

	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// spaceOptsByIDs specifies the parameters for get spaces by ids
type spaceOptsByIDs struct {
	IDs string `url:"ids"`
	SpaceOpts
}

// LookupByIDs Returns details about multiple Spaces. Up to 100 comma-separated Spaces IDs can be looked up using this endpoint.
// Refer: https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces
func (r *SpaceResource) LookupByIDs(ids string, args SpaceOpts) (*SpacesResp, *APIError) {
	path := Baseurl + "/spaces"
	newArgs := spaceOptsByIDs{ids, args}
	resp := new(SpacesResp)

	err := r.Cli.DoGet(path, newArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// spaceOptsByCreators specifies the parameters for get spaces by creators
type spaceOptsByCreators struct {
	UserIDs string `url:"user_ids"`
	SpaceOpts
}

// LookupByCreators Returns live or scheduled Spaces created by the specified user IDs.
// Refer: https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-by-creator-ids
func (r *SpaceResource) LookupByCreators(userIDs string, args SpaceOpts) (*SpacesResp, *APIError) {
	path := Baseurl + "/spaces/by/creator_ids"
	newArgs := spaceOptsByCreators{userIDs, args}
	resp := new(SpacesResp)

	err := r.Cli.DoGet(path, newArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
