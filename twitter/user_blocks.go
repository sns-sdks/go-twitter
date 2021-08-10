package twitter

// UserBlockingOpts specifies the parameters for get blocking users
type UserBlockingOpts struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	UserOpts
}

// GetBlocking Returns a list of users who are blocked by the specified user ID.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/get-users-blocking
func (r *UserResource) GetBlocking(id string, args UserBlockingOpts) (*UsersResp, *APIError) {
	path := Baseurl + "/users/" + id + "/blocking"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// blockingOpts specifies the parameters for blocks create
type blockingOpts struct {
	TargetUserID string `json:"target_user_id"`
}

// BlockingStatus represents the status for blocking
type BlockingStatus struct {
	Blocking *bool `json:"blocking"`
}

func (b BlockingStatus) String() string {
	return Stringify(b)
}

// BlockingResp data struct represents response for blocking
type BlockingResp struct {
	Data *BlockingStatus `json:"data,omitempty"`
}

func (b BlockingResp) String() string {
	return Stringify(b)
}

// BlockingCreate Causes the user (in the path) to block the target user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/post-users-user_id-blocking
func (r *UserResource) BlockingCreate(id, targetUserID string) (*BlockingResp, *APIError) {
	path := Baseurl + "/users/" + id + "/blocking"
	postArgs := blockingOpts{TargetUserID: targetUserID}

	resp := new(BlockingResp)
	err := r.Cli.DoPost(path, postArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BlockingDestroy Allows authenticated user ID to unblock another user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/delete-users-user_id-blocking
func (r *UserResource) BlockingDestroy(id, targetUserID string) (*BlockingResp, *APIError) {
	path := Baseurl + "/users/" + id + "/blocking/" + targetUserID

	resp := new(BlockingResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
