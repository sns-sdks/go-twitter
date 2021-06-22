package twitter

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (bc *BCSuite) TestUserByID() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		"GET",
		BASEURL+"/users/"+uid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"2244994945","name":"Twitter Dev","username":"TwitterDev","public_metrics":{"follower_count": 10}}}`,
		),
	)

	resp, _ := bc.Tw.Users.LookupByID(uid, UserParams{})
	bc.Equal(*resp.Data.ID, uid)
	bc.Equal(*resp.Data.PublicMetrics.FollowerCount, 10)
}

func (bc *BCSuite) TestUsersByIDs() {
	ids := "2244994945,783214"

	httpmock.RegisterResponder(
		"GET",
		BASEURL+"/users",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"2244994945","username":"TwitterDev","name":"Twitter Dev"},{"id":"783214","username":"Twitter","name":"Twitter"}]}`,
		),
	)

	resp, _ := bc.Tw.Users.LookupByIDs(UserParams{IDs: ids})
	bc.Equal(*resp.Data[0].ID, "2244994945")
	bc.Equal(len(resp.Data), 2)
}

func (bc *BCSuite) TestUserByUsername() {
	username := "TwitterDev"

	httpmock.RegisterResponder(
		"GET",
		BASEURL+"/users/by/username/"+username,
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"2244994945","name":"Twitter Dev","username":"TwitterDev"}}`,
		),
	)
	resp, _ := bc.Tw.Users.LookupByUsername(username, UserParams{})
	bc.Equal(*resp.Data.Username, username)
}

func (bc *BCSuite) TestUsersByUsernames() {
	usernames := "TwitterDev,Twitter"

	httpmock.RegisterResponder(
		"GET",
		BASEURL+"/users/by",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"2244994945","username":"TwitterDev","name":"Twitter Dev"},{"id":"783214","username":"Twitter","name":"Twitter"}]}`,
		),
	)

	resp, _ := bc.Tw.Users.LookupByUsernames(UserParams{Usernames: usernames})
	bc.Equal(len(resp.Data), 2)
	bc.Equal(*resp.Data[0].Username, "TwitterDev")
}

func TestBCSuite(t *testing.T) {
	suite.Run(t, new(BCSuite))
}
