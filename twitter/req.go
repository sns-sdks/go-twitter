package twitter

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	goquery "github.com/google/go-querystring/query"
	"net/http"
)

// Meta is Response Data for request.
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
	Users []*User `json:"users,omitempty"`
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

type BaseData struct {
	Includes *Includes `json:"includes,omitempty"`
	Meta     *Meta     `json:"meta,omitempty"`
	Error    []*Error  `json:"error,omitempty"`
}

type UserResp struct {
	Data *User `json:"data,omitempty"`
	*BaseData
}

type UsersResp struct {
	Data []*User `json:"data,omitempty"`
	*BaseData
}

type TweetResp struct {
	Data *Tweet `json:"data,omitempty"`
	*BaseData
}

type TweetsResp struct {
	Data []*Tweet `json:"data,omitempty"`
	*BaseData
}

type TweetsCounts struct {
	End        *string `json:"end,omitempty"`
	Start      *string `json:"start,omitempty"`
	TweetCount *int    `json:"tweet_count,omitempty"`
}

type TweetsCountsResp struct {
	Data []*TweetsCounts `json:"data,omitempty"`
	*BaseData
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
		v, err := goquery.Values(jsonParams)
		if err != nil {
			apiError := APIError{Title: "Form param Error", Detail: err.Error()}
			return &apiError
		}
		req.SetFormDataFromValues(v)
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
