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
