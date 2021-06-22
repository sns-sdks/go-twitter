package twitter

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	goquery "github.com/google/go-querystring/query"
	"net/http"
)

/*
	functions for http requests
*/

//// Resp response structure for twitter data
//type Resp struct {
//	Data     json.RawMessage `json:"data,omitempty"`
//	Includes json.RawMessage `json:"includes,omitempty"`
//	Meta     json.RawMessage `json:"meta,omitempty"`
//	Errors   json.RawMessage `json:"errors,omitempty"`
//}

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
