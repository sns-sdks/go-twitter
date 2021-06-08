package integration

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"go-twitter/twitter"
	"testing"
)

func (bc *BCSuite) TestUserByID() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		"GET",
		twitter.BASEURL+"/users/"+uid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"2244994945","name":"Twitter Dev","username":"TwitterDev","public_metrics":{"follower_count": 10}}}`,
		),
	)

	user, _ := bc.Tw.Users.LookupByID("2244994945", twitter.UserParams{})
	bc.Equal(*user.ID, uid)
	bc.Equal(*user.PublicMetrics.FollowerCount, 10)
}

func TestBCSuite(t *testing.T) {
	suite.Run(t, new(BCSuite))
}
