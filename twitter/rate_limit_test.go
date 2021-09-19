package twitter

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRateLimit(t *testing.T) {
	r := RateLimit{}

	header := http.Header{}
	header.Add("x-rate-limit-limit", "300")
	header.Add("x-rate-limit-remaining", "199")
	header.Add("x-rate-limit-reset", "1234567")
	r.Set("users/aaaaaa", "GET", header)
	r.Set("users/123124125", "GET", header)

	assert.Equal(t, r.Get(), `twitter.RateLimit{Mapping:map[users/:id:map[GET:twitter.RateLimitData{Limit:300, Remaining:199, Reset:1234567}]]}`)
	assert.Equal(t, r.GetByURL("users/123124125", "GET"), `twitter.RateLimitData{Limit:300, Remaining:199, Reset:1234567}`)
	assert.Equal(t, r.GetByURL("users", "GET"), `twitter.RateLimitData{Limit:0, Remaining:0, Reset:0}`)
}
