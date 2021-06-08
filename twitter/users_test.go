package twitter

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsers(t *testing.T) {
	uid := "2244994945"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		BASEURL + "/users/"+uid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"2244994945","name":"Twitter Dev","username":"TwitterDev"}}`,
		),
	)

	cli := NewBearerClient("")
	user, _ := cli.Users.lookupByID("2244994945", UserParams{})
	assert.Equal(t, user.ID, uid)
}
