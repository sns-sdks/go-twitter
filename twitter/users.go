package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	ent "go-twitter/twitter/entities"
)

/*
	Users API

*/

type UserResource struct {
	cli *resty.Client
}

func newUserResource(cli *resty.Client) *UserResource {
	return &UserResource{
		cli: cli,
	}
}

func (r *UserResource) lookupByID(id string) (*ent.User, *APIError) {
	resp, err := r.cli.R().SetQueryParams(map[string]string{
		"user.fields": "description",
	}).Get(BASEURL + "/users/" + id)
	fmt.Println(resp)
	if err == nil {
		data, err := ParseDataResponse(resp)
		if data != nil {
			user := new(ent.User)
			_ = json.Unmarshal(data.Data, &user)
			return user, nil
		}
		return nil, err
	}
	apiError := APIError{Title: "HTTP Error", Detail: err.Error()}
	return nil, &apiError
}
