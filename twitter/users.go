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
	cli *resty.Client
}

type U struct {
	E  string
	T  string `json:"t"`
	Us string `json:"us"`
}

type UserParams struct {
	Expansions  string `json:"expansions,omitempty"`
	TweetFields string `json:"tweet.fields,omitempty"`
	UserFields  string `json:"user.fields,omitempty"`
}

func newUserResource(cli *resty.Client) *UserResource {
	return &UserResource{
		cli: cli,
	}
}

func (r *UserResource) lookupByID(id string, params UserParams) (*ent.User, *APIError) {
	path := BASEURL + "/users/" + id
	data, err := DoRequest(r.cli, resty.MethodGet, path, params, nil)
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
