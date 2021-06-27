package twitter

import ent "github.com/sns-sdks/go-twitter/twitter/entities"


/*
	Users API
*/

type UserResource Resource

type UserParams struct {
	IDs         string `url:"ids,omitempty"`
	Usernames   string `url:"usernames,omitempty"`
	UserFields  string `url:"user.fields,omitempty"`
	Expansions  string `url:"expansions,omitempty"`
	TweetFields string `url:"tweet.fields,omitempty"`
}

type FollowParams struct {
	MaxResults      string `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	UserFields      string `url:"user.fields,omitempty"`
	Expansions      string `url:"expansions,omitempty"`
	TweetFields     string `url:"tweet.fields,omitempty"`
}

func newUserResource(cli *Client) *UserResource {
	return &UserResource{
		Cli: cli,
	}
}

func (r *UserResource) LookupByID(id string, params UserParams) (*ent.UserResp, *APIError) {
	path := Baseurl + "/users/" + id
	resp := new(ent.UserResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *UserResource) LookupByIDs(params UserParams) (*ent.UsersResp, *APIError) {
	path := Baseurl + "/users"
	resp := new(ent.UsersResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *UserResource) LookupByUsername(username string, params UserParams) (*ent.UserResp, *APIError) {
	path := Baseurl + "/users/by/username/" + username

	resp := new(ent.UserResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *UserResource) LookupByUsernames(params UserParams) (*ent.UsersResp, *APIError) {
	path := Baseurl + "/users/by"

	resp := new(ent.UsersResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *UserResource) GetFollowing(id string, params FollowParams) (*ent.UsersResp, *APIError) {
	path := Baseurl + "/users/" + id + "/following"

	resp := new(ent.UsersResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *UserResource) GetFollowers(id string, params FollowParams) (*ent.UsersResp, *APIError) {
	path := Baseurl + "/users/" + id + "/followers"

	resp := new(ent.UsersResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
