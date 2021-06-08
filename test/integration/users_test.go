package integration

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"go-twitter/twitter"
	"testing"
)

func TestUsers(t *testing.T) {
	uid := "2244994945"
	tw := twitter.NewBearerClient("")

	httpmock.ActivateNonDefault(tw.Cli.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		twitter.BASEURL+"/users/"+uid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"2244994945","name":"Twitter Dev","username":"TwitterDev","public_metrics":{"follower_count": 10}}}`,
		),
	)

	user, _ := tw.Users.LookupByID("2244994945", twitter.UserParams{})
	assert.Equal(t, *user.ID, uid)
}
