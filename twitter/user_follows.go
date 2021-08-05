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
func (r *UserResource) GetFollowing(id string, params FollowsOpts) (*UsersResp, *APIError) {
	path := Baseurl + "/users/" + id + "/following"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetFollowers Returns a list of users who are followers of the specified user ID.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-followers
func (r *UserResource) GetFollowers(id string, params FollowsOpts) (*UsersResp, *APIError) {
	path := Baseurl + "/users/" + id + "/followers"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
