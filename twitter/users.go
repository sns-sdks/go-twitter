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

func (r *UserResource) lookupByID(id string) *ent.User {
	resp, err := r.cli.R().SetQueryParams(map[string]string{
		"user.fields": "description",
	}).Get(BASEURL + "/users/" + id)
	fmt.Println(resp)
	user := new(ent.User)
	if err!=nil{
		data, _ := json.Marshal(resp)
		_ = json.Unmarshal(data, &user)
		return user
	}

	return nil
}

