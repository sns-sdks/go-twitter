package twitter

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	goquery "github.com/google/go-querystring/query"
	"net/http"
)

// Meta represents the Response meta data for request.
type Meta struct {
	ResultCount     *int         `json:"result_count,omitempty"`
	PreviousToken   *string      `json:"previous_token,omitempty"`
	NextToken       *string      `json:"next_token,omitempty"`
	OldestID        *string      `json:"oldest_id,omitempty"`
	NewestID        *string      `json:"newest_id,omitempty"`
	Sent            *string      `json:"sent,omitempty"`
	Summary         *MetaSummary `json:"summary,omitempty"`
	TotalTweetCount *int         `json:"total_tweet_count,omitempty"`
}

// MetaSummary represents the Response meta summary data for request.
type MetaSummary struct {
	Created    *int `json:"created,omitempty"`
	NotCreated *int `json:"not_created,omitempty"`
	Valid      *int `json:"valid,omitempty"`
	Invalid    *int `json:"invalid,omitempty"`
	Deleted    *int `json:"deleted,omitempty"`
	NotDeleted *int `json:"not_deleted,omitempty"`
}

// Includes represents the expansions objects for main request thread.
type Includes struct {
	Users  []*User  `json:"users,omitempty"`
	Tweets []*Tweet `json:"tweets,omitempty"`
	Media  []*Media `json:"media,omitempty"`
	Places []*Place `json:"places,omitempty"`
	Polls  []*Poll  `json:"polls,omitempty"`
}

// Error represents the common error response
type Error struct {
	Detail       *string `json:"detail"`
	Title        *string `json:"title"`
	ResourceType *string `json:"resource_type,omitempty"`
	Parameter    *string `json:"parameter,omitempty"`
	Value        *string `json:"value,omitempty"`
	Type         *string `json:"type,omitempty"`
}

// BaseData Additional response data for a request.
type BaseData struct {
	Includes *Includes `json:"includes,omitempty"`
	Meta     *Meta     `json:"meta,omitempty"`
	Error    []*Error  `json:"error,omitempty"`
}

// UserResp data struct represents the single user response
type UserResp struct {
	Data *User `json:"data,omitempty"`
	*BaseData
}

func (u UserResp) String() string {
	return Stringify(u)
}

// UsersResp data struct represents list users response
type UsersResp struct {
	Data []*User `json:"data,omitempty"`
	*BaseData
}

func (u UsersResp) String() string {
	return Stringify(u)
}

// TweetResp data struct represents the single tweet response
type TweetResp struct {
	Data *Tweet `json:"data,omitempty"`
	*BaseData
}

func (t TweetResp) String() string {
	return Stringify(t)
}

// TweetsResp data struct represents list tweets response
type TweetsResp struct {
	Data []*Tweet `json:"data,omitempty"`
	*BaseData
}

func (t TweetsResp) String() string {
	return Stringify(t)
}

// TweetsCountsResp data struct represents tweet counts response
type TweetsCountsResp struct {
	Data []*TweetsCounts `json:"data,omitempty"`
	*BaseData
}

func (t TweetsCountsResp) String() string {
	return Stringify(t)
}

// SpaceResp data struct represents the single space response
type SpaceResp struct {
	Data *Space `json:"data,omitempty"`
	*BaseData
}

func (s SpaceResp) String() string {
	return Stringify(s)
}

// SpacesResp data struct represents the list spaces response
type SpacesResp struct {
	Data []*Space `json:"data,omitempty"`
	*BaseData
}

func (s SpacesResp) String() string {
	return Stringify(s)
}

/*
	functions for http requests
*/
func ParseDataResponse(response *resty.Response, d interface{}) *APIError {
	var err error
	if response.StatusCode() == http.StatusOK {
		switch d := d.(type) {
		case nil:
		default:
			err = json.Unmarshal(response.Body(), d)
		}
		if err != nil {
			return &APIError{Title: "Json Unmarshal data", Detail: err.Error()}
		}
		return nil
	}
	apiErr := new(APIError)
	err = json.Unmarshal(response.Body(), &apiErr)
	if err != nil {
		return &APIError{Title: "Json Unmarshal error", Detail: err.Error()}
	}
	return apiErr
}

func (r *Client) Do(method, path string, queryParams interface{}, jsonParams interface{}, d interface{}) *APIError {
	req := r.Cli.R()

	// parse struct params
	if queryParams != nil {
		v, err := goquery.Values(queryParams)
		if err != nil {
			apiError := APIError{Title: "Query param Error", Detail: err.Error()}
			return &apiError
		}
		req.SetQueryParamsFromValues(v)
	}
	if jsonParams != nil {
		req.SetBody(jsonParams)
		req.SetHeader("Content-Type", "application/json")
	}

	resp, err := req.Execute(method, path)
	if err != nil {
		apiError := APIError{Title: "HTTP Error", Detail: err.Error()}
		return &apiError
	}
	apiError := ParseDataResponse(resp, d)
	return apiError
}

func (r *Client) DoGet(path string, queryParams interface{}, d interface{}) *APIError {
	return r.Do(HttpGet, path, queryParams, nil, d)
}

func (r *Client) DoPost(path string, jsonParams interface{}, d interface{}) *APIError {
	return r.Do(HttpPost, path, nil, jsonParams, d)
}

func (r *Client) DoPut(path string, jsonParams interface{}, d interface{}) *APIError {
	return r.Do(HttpPut, path, nil, jsonParams, d)
}

func (r *Client) DoDelete(path string, d interface{}) *APIError {
	return r.Do(HttpDelete, path, nil, nil, d)
}
