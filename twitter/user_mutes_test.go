package twitter

import "github.com/jarcoal/httpmock"

func (uc *UCSuite) TestGetMuting() {
	uid := "1324848235714736129"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/muting",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Users.GetMuting(uid, MutingOpts{})
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/muting",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"2244994945","name":"Twitter Dev","username":"TwitterDev"}],"meta":{"result_count":1}}`,
		),
	)

	resp, _ := uc.Tw.Users.GetMuting(uid, MutingOpts{})
	uc.Equal(*resp.Meta.ResultCount, 1)
	uc.Equal(*resp.Data[0].ID, "2244994945")

}

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
