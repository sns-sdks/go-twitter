package twitter

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestParseDataResponse(t *testing.T) {
	resp := resty.Response{RawResponse: &http.Response{StatusCode: 200}}

	err := ParseDataResponse(&resp, nil)
	assert.Nil(t, err)

	err = ParseDataResponse(&resp, "123")
	assert.IsType(t, &APIError{}, err)

	errResp := resty.Response{RawResponse: &http.Response{StatusCode: 401}}

	err = ParseDataResponse(&errResp, "123")
	assert.IsType(t, "", err.Error())
	assert.IsType(t, &APIError{}, err)
}

func TestDo(t *testing.T) {
	cli := NewBearerClient("")

	cli.Cli.R().SetFile("", "")
	err := cli.Do("DELETE", "", "", "", "")
	assert.IsType(t, &APIError{}, err)
}
