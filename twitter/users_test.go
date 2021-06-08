package twitter

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsers(t *testing.T) {
	uid := "2244994945"
	tw := NewBearerClient("")

	httpmock.ActivateNonDefault(tw.cli.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		BASEURL+"/users/"+uid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"2244994945","name":"Twitter Dev","username":"TwitterDev","public_metrics":{"follower_count": 10}}}`,
		),
	)

	user, _ := tw.Users.lookupByID("2244994945", UserParams{})
	assert.Equal(t, *user.ID, uid)
}
