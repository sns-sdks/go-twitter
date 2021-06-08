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

	user, _ := bc.Tw.Users.LookupByID("2244994945", UserParams{})
	bc.Equal(*user.ID, uid)
	bc.Equal(*user.PublicMetrics.FollowerCount, 10)
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

	users, _ := bc.Tw.Users.LookupByIDs(UserParams{IDs: ids})
	bc.Equal(*users[0].ID, "2244994945")
	bc.Equal(len(users), 2)
}

func TestBCSuite(t *testing.T) {
	suite.Run(t, new(BCSuite))
}
