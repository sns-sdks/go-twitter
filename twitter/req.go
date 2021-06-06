package twitter

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"net/http"
)

/*
	functions for http requests
*/

type Resp struct {
	Data     json.RawMessage `json:"data,omitempty"`
	Includes json.RawMessage `json:"includes,omitempty"`
	Meta     json.RawMessage `json:"meta,omitempty"`
	Errors   json.RawMessage `json:"errors,omitempty"`
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
