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

func (bc *BCSuite) TestGetUserPinnedLists() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/pinned_lists",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.GetUserPinnedLists(uid, ListOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/pinned_lists",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"follower_count":0,"id":"1451305624956858369","name":"Test List","owner_id":"2244994945"}],"includes":{"users":[{"username":"TwitterDev","id":"2244994945","created_at":"2013-12-14T04:35:55.000Z","name":"Twitter Dev"}]},"meta":{"result_count":1}}`,
		),
	)

	resp, _ := bc.Tw.Lists.GetUserPinnedLists(uid, ListOpts{Expansions: "owner_id", ListFields: "follower_count", UserFields: "created_at"})
	bc.Equal(*resp.Data[0].ID, "1451305624956858369")
	bc.Equal(*resp.Includes.Users[0].ID, "2244994945")
}
