package twitter

// FollowsOpts specifies the parameters for get follows
type FollowsOpts struct {
	MaxResults      string `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	UserFields      string `url:"user.fields,omitempty"`
	Expansions      string `url:"expansions,omitempty"`
	TweetFields     string `url:"tweet.fields,omitempty"`
}

// GetFollowing Returns a list of users the specified user ID is following.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following
func (r *UserResource) GetFollowing(id string, args FollowsOpts) (*UsersResp, *APIError) {
	path := "/users/" + id + "/following"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetFollowers Returns a list of users who are followers of the specified user ID.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-followers
func (r *UserResource) GetFollowers(id string, args FollowsOpts) (*UsersResp, *APIError) {
	path := "/users/" + id + "/followers"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// followingOpts specifies the parameters for follows create
type followingOpts struct {
	TargetUserID string `json:"target_user_id"`
}

// FollowingStatus represents status for following
type FollowingStatus struct {
	Following     *bool `json:"following,omitempty"`
	PendingFollow *bool `json:"pending_follow,omitempty"`
}

func (f FollowingStatus) String() string {
	return Stringify(f)
}

// FollowingResp data struct represents response for following
type FollowingResp struct {
	Data *FollowingStatus `json:"data,omitempty"`
}

func (f FollowingResp) String() string {
	return Stringify(f)
}

// FollowingCreate Allows a user ID to follow another user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/post-users-source_user_id-following
func (r *UserResource) FollowingCreate(id, targetUserID string) (*FollowingResp, *APIError) {
	path := "/users/" + id + "/following"
	postArgs := followingOpts{TargetUserID: targetUserID}

	resp := new(FollowingResp)
	err := r.Cli.DoPost(path, postArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FollowingDestroy Allows a user ID to unfollow another user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/delete-users-source_id-following
func (r *UserResource) FollowingDestroy(id, targetUserID string) (*FollowingResp, *APIError) {
	path := "/users/" + id + "/following/" + targetUserID

	resp := new(FollowingResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
