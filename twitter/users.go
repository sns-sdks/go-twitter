package twitter

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	ent "go-twitter/twitter/entities"
)

/*
	Users API
*/

type UserResource struct {
	Cli *resty.Client
}

type UserParams struct {
	IDs         string `url:"ids,omitempty"`
	Usernames   string `url:"usernames,omitempty"`
	Expansions  string `url:"expansions,omitempty"`
	TweetFields string `url:"tweet.fields,omitempty"`
	UserFields  string `url:"user.fields,omitempty"`
}

type FollowParams struct {
	MaxResults      string `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	Expansions      string `url:"expansions,omitempty"`
	TweetFields     string `url:"tweet.fields,omitempty"`
	UserFields      string `url:"user.fields,omitempty"`
}

func newUserResource(cli *resty.Client) *UserResource {
	return &UserResource{
		Cli: cli,
	}
}

func (r *UserResource) LookupByID(id string, params UserParams) (*ent.User, *APIError) {
	path := BASEURL + "/users/" + id
	data, err := DoRequest(r.Cli, resty.MethodGet, path, params, nil)
	if err != nil {
		return nil, err
	}
	user := new(ent.User)
	jErr := json.Unmarshal(data.Data, &user)
	if jErr != nil {
		apiError := APIError{Title: "Json Error", Detail: jErr.Error()}
		return nil, &apiError
	}
	return user, nil
}

func (r *UserResource) LookupByIDs(params UserParams) ([]*ent.User, *APIError) {
	path := BASEURL + "/users"
	data, err := DoRequest(r.Cli, resty.MethodGet, path, params, nil)
	if err != nil {
		return nil, err
	}

	users := new([]*ent.User)
	jErr := json.Unmarshal(data.Data, &users)
	if jErr != nil {
		apiError := APIError{Title: "Json Error", Detail: jErr.Error()}
		return nil, &apiError
	}
	return *users, nil
}

func (r *UserResource) LookupByUsername(username string, params UserParams) (*ent.User, *APIError) {
	path := BASEURL + "/users/by/username/" + username
	data, err := DoRequest(r.Cli, resty.MethodGet, path, params, nil)
	if err != nil {
		return nil, err
	}
	user := new(ent.User)
	jErr := json.Unmarshal(data.Data, &user)
	if jErr != nil {
		apiError := APIError{Title: "Json Error", Detail: jErr.Error()}
		return nil, &apiError
	}
	return user, nil
}

func (r *UserResource) LookupByUsernames(params UserParams) ([]*ent.User, *APIError) {
	path := BASEURL + "/users/by"
	data, err := DoRequest(r.Cli, resty.MethodGet, path, params, nil)
	if err != nil {
		return nil, err
	}
	users := new([]*ent.User)
	jErr := json.Unmarshal(data.Data, &users)
	if jErr != nil {
		apiError := APIError{Title: "Json Error", Detail: jErr.Error()}
		return nil, &apiError
	}
	return *users, nil
}

func (r *UserResource) GetFollowing(id string, params FollowParams) ([]*ent.User, *APIError) {
	path := BASEURL + "/users/" + id + "/following"
	data, err := DoRequest(r.Cli, resty.MethodGet, path, params, nil)
	if err != nil {
		return nil, err
	}
	users := new([]*ent.User)
	jErr := json.Unmarshal(data.Data, &users)
	if jErr != nil {
		apiError := APIError{Title: "Json Error", Detail: jErr.Error()}
		return nil, &apiError
	}
	return *users, nil
}

func (r *UserResource) GetFollowers(id string, params FollowParams) ([]*ent.User, *APIError) {
	path := BASEURL + "/users/" + id + "/followers"
	data, err := DoRequest(r.Cli, resty.MethodGet, path, params, nil)
	if err != nil {
		return nil, err
	}
	users := new([]*ent.User)
	jErr := json.Unmarshal(data.Data, &users)
	if jErr != nil {
		apiError := APIError{Title: "Json Error", Detail: jErr.Error()}
		return nil, &apiError
	}
	return *users, nil
}
