package twitter

import (
	"net/http"
	"regexp"
	"strconv"
)

/*
	Rate limits for every API.
	Refer: https://developer.twitter.com/en/docs/twitter-api/rate-limits
*/

type RateLimitData struct {
	Limit     int `json:"x-rate-limit-limit,omitempty"`
	Remaining int `json:"x-rate-limit-remaining,omitempty"`
	Reset     int `json:"x-rate-limit-reset,omitempty"`
}

func (r RateLimitData) String() string {
	return Stringify(r)
}

type Endpoint struct {
	Resource string `json:"resource"`
	Regex    string `json:"regex"`
}

var RateLimitResource = [...]Endpoint{
	{Resource: "/users", Regex: `^/users$`},
	{Resource: "/users/:id", Regex: `^/users/\d+$`},
	{Resource: "/users/by", Regex: `^/users/by$`},
	{Resource: "/users/by/username/:username", Regex: `^/users/by/username/\w+$`},
	{Resource: "/users/:id/following", Regex: `^/users/\d+/following$`},
	{Resource: "/users/:id/following/:target_user_id", Regex: `^/users/\d+/following/\d+$`},
	{Resource: "/users/:id/followers", Regex: `^/users/\d+/followers$`},
	{Resource: "/users/:id/blocking", Regex: `^/users/\d+/blocking$`},
	{Resource: "/users/:id/blocking/:target_user_id", Regex: `^/users/\d+/blocking/\d+$`},
	{Resource: "/users/:id/muting", Regex: `^/users/\d+/muting$`},
	{Resource: "/users/:id/muting/:target_user_id", Regex: `^/users/\d+/muting/\d+$`},
	{Resource: "/tweets/:id", Regex: `^/tweets/\d+$`},
	{Resource: "/tweets", Regex: `^/tweets$`},
	{Resource: "/tweets/search/recent", Regex: `^/tweets/search/recent$`},
	{Resource: "/tweets/search/all", Regex: `^/tweets/search/all$`},
	{Resource: "/users/:id/tweets", Regex: `^/users/\d+/tweets$`},
	{Resource: "/users/:id/timelines/reverse_chronological", Regex: `^/users/\d+/timelines/reverse_chronological$`},
	{Resource: "/users/:id/mentions", Regex: `^/users/\d+/mentions$`},
	{Resource: "/tweets/:id/liking_users", Regex: `^/tweets/\d+/liking_users$`},
	{Resource: "/tweets/:id/retweeted_by", Regex: `^/tweets/\d+/retweeted_by$`},
	{Resource: "/tweets/:id/quote_tweets", Regex: `^/tweets/\d+/quote_tweets$`},
	{Resource: "/users/:id/liked_tweets", Regex: `^/users/\d+/liked_tweets$`},
	{Resource: "/users/:id/likes", Regex: `^/users/\d+/likes$`},
	{Resource: "/users/:id/likes/:tweet_id", Regex: `^/users/\d+/likes/\d+$`},
	{Resource: "/users/:id/retweets", Regex: `^/users/\d+/retweets$`},
	{Resource: "/users/:id/retweets/:tweet_id", Regex: `^/users/\d+/retweets/\d+$`},
	{Resource: "/users/:id/followed_lists", Regex: `^/users/\d+/followed_lists`},
	{Resource: "/users/:id/followed_lists/:list_id", Regex: `^/users/\d+/followed_lists/\d+$`},
	{Resource: "/users/:id/pinned_lists", Regex: `^/users/\d+/retweets$`},
	{Resource: "/users/:id/pinned_lists/:list_id", Regex: `^/users/\d+/pinned_lists/\d+$`},
	{Resource: "/users/:id/bookmarks", Regex: `^/users/\d+/bookmarks$`},
	{Resource: "users/:id/bookmarks/:tweet_id", Regex: `^/users/\d+/bookmarks$/\d+$`},
	{Resource: "/tweets/:id/hidden", Regex: `^/tweets/\d+/hidden$`},
	{Resource: "/tweets/counts", Regex: `^tweets/counts/\w+$`},
	{Resource: "/spaces/:id", Regex: `^/spaces/\w+$`},
	{Resource: "/spaces", Regex: `^/spaces$`},
	{Resource: "/spaces/by/creator_ids", Regex: `^/spaces/by/creator_ids$`},
	{Resource: "/spaces/search", Regex: `^/spaces/search$`},
	{Resource: "/lists", Regex: `^/lists$`},
	{Resource: "/lists/:id", Regex: `^/lists/\d+$`},
	{Resource: "/users/:id/owned_lists", Regex: `^/users/\d+/owned_lists$`},
	{Resource: "/lists/:id/tweets", Regex: `^/lists/\d+/tweets$`},
	{Resource: "/lists/:id/members", Regex: `^/lists/\d+/members$`},
	{Resource: "/users/:id/list_memberships", Regex: `^/users/\d+/list_memberships$`},
	{Resource: "/lists/:id/members/:user_id", Regex: `^/lists/\d+/members/\d+$`},
	{Resource: "/lists/:id/followers", Regex: `^/lists/\d+/followers$`},
	{Resource: "/users/:id/followed_lists", Regex: `^/users/\d+/followed_lists`},
	{Resource: "/users/:id/followed_lists/:list_id", Regex: `^/users/\d+/followed_lists/\d+$`},
	{Resource: "/users/:id/pinned_lists", Regex: `^/users/\d+/pinned_lists`},
	{Resource: "/users/:id/pinned_lists/:list_id", Regex: `^/users/\d+/pinned_lists/\d+$`},
	{Resource: "/compliance/jobs/:job_id", Regex: `^/compliance/jobs/\d+$`},
	{Resource: "/compliance/jobs", Regex: `^/compliance/jobs$`},
}

func getResource(url string) string {
	for _, e := range RateLimitResource {
		b := checkResource(url, e)
		if b {
			return e.Resource
		}
	}
	return ""
}

func checkResource(url string, e Endpoint) bool {
	reg := regexp.MustCompile(e.Regex)
	r := reg.MatchString(url)
	return r
}

type RateLimit struct {
	// Mapping container rate limit for resource.
	// Ex: {map[users:map[GET:twitter.RateLimitData{Limit:300, Remaining:300, Reset:1123124}]]}
	Mapping map[string]map[string]RateLimitData `json:"mapping"`
}

func newRateLimit() *RateLimit {
	return &RateLimit{}
}

func (r *RateLimit) String() string {
	return Stringify(r)
}

func (r *RateLimit) Set(url, method string, header http.Header) {
	resource := getResource(url)
	if resource == "" {
		return
	}

	limit, _ := strconv.Atoi(header.Get("x-rate-limit-limit"))
	remaining, _ := strconv.Atoi(header.Get("x-rate-limit-remaining"))
	reset, _ := strconv.Atoi(header.Get("x-rate-limit-reset"))

	if r.Mapping == nil {
		r.Mapping = make(map[string]map[string]RateLimitData)
	}

	methodLimit := map[string]RateLimitData{method: {
		Limit: limit, Remaining: remaining, Reset: reset,
	}}
	r.Mapping[resource] = methodLimit
}

func (r *RateLimit) Get() string {
	return r.String()
}

func (r *RateLimit) GetByURL(url, method string) string {
	source := getResource(url)
	endpoint, ok := r.Mapping[source]
	if ok {
		limitData, ok := endpoint[method]
		if ok {
			return limitData.String()
		}
	}
	return RateLimitData{}.String()
}
