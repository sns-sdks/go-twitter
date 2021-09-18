package twitter

import (
	"net/http"
	"regexp"
	"strconv"
)

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
	{Resource: "users", Regex: `users$`},
	{Resource: "users/:id", Regex: `users/\d+$`},
	{Resource: "users/by", Regex: `users/by$`},
	{Resource: "users/by/username/:username", Regex: `/users/by/username/\w+$`},
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
