package twitter

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/go-resty/resty/v2"
	ent "go-twitter/twitter/entities"
)

/*
	Users API
*/

type UserResource struct {
	cli *resty.Client
}

type UserParams struct {
	Expansions string `json:"expansions,omitempty"`
	TweetFields string `json:"tweet.fields,omitempty"`
	UserFields string `json:"user.fields,omitempty"`
}

func newUserResource(cli *resty.Client) *UserResource {
	return &UserResource{
		cli: cli,
	}
}

func (r *UserResource) lookupByID(id string, params UserParams) (*ent.User, *APIError) {
	path := BASEURL + "/users/" + id
	data, err := DoRequest(r.cli, resty.MethodGet, path, structs.Map(&params), nil)
	if err != nil{
		return nil, err
	}
	user := new(ent.User)
	_ = json.Unmarshal(data.Data, &user)
	return user, nil
}
