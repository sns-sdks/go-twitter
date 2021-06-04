package twitter

import "testing"

func TestUsers(t *testing.T) {
	cli := NewBearerClient("")
	r := cli.Users.lookupByID("")
	t.Log(r)
}
