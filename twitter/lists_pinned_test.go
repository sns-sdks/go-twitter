package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestPinList() {
	lid := "1441162269824405510"
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/pinned_lists",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.PinList(uid, lid)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/pinned_lists",
		httpmock.NewStringResponder(
			200,
			`{"data":{"pinned":true}}`,
		),
	)

	resp, _ := bc.Tw.Lists.PinList(uid, lid)
	bc.Equal(*resp.Data.Pinned, true)
}

func (bc *BCSuite) TestRemovePinnedList() {
	lid := "1441162269824405510"
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/pinned_lists/"+lid,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.RemovePinnedList(uid, lid)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/pinned_lists/"+lid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"pinned":false}}`,
		),
	)

	resp, _ := bc.Tw.Lists.RemovePinnedList(uid, lid)
	bc.Equal(*resp.Data.Pinned, false)
}
