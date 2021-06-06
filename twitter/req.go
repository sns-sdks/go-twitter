package twitter

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"net/http"
)

/*
	functions for http requests
*/

// Resp response structure for twitter data
type Resp struct {
	Data     json.RawMessage `json:"data,omitempty"`
	Includes json.RawMessage `json:"includes,omitempty"`
	Meta     json.RawMessage `json:"meta,omitempty"`
	Errors   json.RawMessage `json:"errors,omitempty"`
}

func DoRequest(cli *resty.Client, method, path string, queryParam map[string]string, jsonParam map[string]string) (*Resp, *APIError) {
	resp, err := cli.R().SetQueryParams(queryParam).SetBody(jsonParam).Execute(method, path)
	if err != nil{
		apiError := APIError{Title: "HTTP Error", Detail: err.Error()}
		return nil, &apiError
	}
	data, apiError := ParseDataResponse(resp)
	if data != nil {
		return data, nil
	}
	return nil, apiError
}

func ParseDataResponse(response *resty.Response) (*Resp, *APIError) {
	if response.StatusCode() == http.StatusOK {
		resp := new(Resp)
		err := json.Unmarshal(response.Body(), &resp)
		if err != nil {
			return nil, &APIError{Title: "Json Unmarshal data", Detail: err.Error()}
		}
		return resp, nil
	}
	apiErr := new(APIError)
	err := json.Unmarshal(response.Body(), &apiErr)
	if err != nil {
		return nil, &APIError{Title: "Json Unmarshal error", Detail: err.Error()}
	}
	return nil, apiErr
}
