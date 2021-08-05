package twitter

type FollowsOpts struct {
	MaxResults      string `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	UserFields      string `url:"user.fields,omitempty"`
	Expansions      string `url:"expansions,omitempty"`
	TweetFields     string `url:"tweet.fields,omitempty"`
}

func (r *UserResource) GetFollowing(id string, params FollowsOpts) (*UsersResp, *APIError) {
	path := Baseurl + "/users/" + id + "/following"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *UserResource) GetFollowers(id string, params FollowsOpts) (*UsersResp, *APIError) {
	path := Baseurl + "/users/" + id + "/followers"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
