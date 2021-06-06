package twitter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsers(t *testing.T) {
	a := assert.New(t)
	cli := NewBearerClient("")
	r, err := cli.Users.lookupByID("124124125412", UserParams{})
	a.Nil(r)
	if err != nil {
		a.Equal(err.Status, 401)
	}
}
