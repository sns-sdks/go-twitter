package twitter

import "github.com/jarcoal/httpmock"

func (uc *UCSuite) TestCreateMuting() {
	uid := "123456789"
	targetID := "2244994945"

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/muting",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Users.MutingCreate(uid, targetID)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/muting",
		httpmock.NewStringResponder(
			200,
			`{"data":{"muting":true}}`,
		),
	)

	resp, _ := uc.Tw.Users.MutingCreate(uid, targetID)
	uc.Equal(*resp.Data.Muting, true)
}

func (uc *UCSuite) TestDestroyMuting() {
	uid := "123456789"
	targetID := "2244994945"

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/muting/"+targetID,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Users.MutingDestroy(uid, targetID)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/muting/"+targetID,
		httpmock.NewStringResponder(
			200,
			`{"data":{"muting":false}}`,
		),
	)

	resp, _ := uc.Tw.Users.MutingDestroy(uid, targetID)
	uc.Equal(*resp.Data.Muting, false)
}
