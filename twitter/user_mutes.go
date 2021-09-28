package twitter

// mutingOpts specifies the parameters for mutes create
type mutingOpts struct {
	TargetUserID string `json:"target_user_id"`
}

// MutingStatus represents status for muting
type MutingStatus struct {
	Muting *bool `json:"muting,omitempty"`
}

func (m MutingStatus) String() string {
	return Stringify(m)
}

// MutingResp data struct represents response for muting
type MutingResp struct {
	Data *MutingStatus `json:"data,omitempty"`
}

func (m MutingResp) String() string {
	return Stringify(m)
}

// MutingOpts specifies the parameters for get muting
type MutingOpts struct {
	MaxResults      string `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	UserFields      string `url:"user.fields,omitempty"`
	Expansions      string `url:"expansions,omitempty"`
	TweetFields     string `url:"tweet.fields,omitempty"`
}

// GetMuting Returns a list of users who are muted by the specified user ID.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/get-users-muting
func (r *UserResource) GetMuting(id string, args MutingOpts) (*UsersResp, *APIError) {
	path := "/users/" + id + "/muting"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MutingCreate Allows an authenticated user ID to mute the target user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/post-users-user_id-muting
func (r *UserResource) MutingCreate(id, targetUserID string) (*MutingResp, *APIError) {
	path := "/users/" + id + "/muting"
	postArgs := mutingOpts{TargetUserID: targetUserID}

	resp := new(MutingResp)
	err := r.Cli.DoPost(path, postArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MutingDestroy Allows an authenticated user ID to unmute the target user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/delete-users-user_id-muting
func (r *UserResource) MutingDestroy(id, targetUserID string) (*MutingResp, *APIError) {
	path := "/users/" + id + "/muting/" + targetUserID

	resp := new(MutingResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
