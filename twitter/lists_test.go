package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestLookupList() {
	lid := "84839422"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/lists/"+lid,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.LookupByID(lid, ListOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/lists/"+lid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"follower_count":906,"id":"84839422","name":"Official Twitter Accounts","owner_id":"783214"},"includes":{"users":[{"id":"783214","name":"Twitter","username":"Twitter"}]}}`,
		),
	)

	resp, _ := bc.Tw.Lists.LookupByID(lid, ListOpts{})
	bc.Equal(*resp.Data.FollowerCount, 906)
	bc.Equal(*resp.Includes.Users[0].ID, "783214")
}

func (bc *BCSuite) TestGetOwnedLists() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/owned_lists",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.GetOwnedLists(uid, OwnedListsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/owned_lists",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"follower_count":0,"id":"1451305624956858369","name":"Test List","owner_id":"2244994945"}],"includes":{"users":[{"username":"TwitterDev","id":"2244994945","created_at":"2013-12-14T04:35:55.000Z","name":"Twitter Dev"}]},"meta":{"result_count":1}}`,
		),
	)

	resp, _ := bc.Tw.Lists.GetOwnedLists(uid, OwnedListsOpts{})
	bc.Equal(*resp.Data[0].ID, "1451305624956858369")
	bc.Equal(*resp.Includes.Users[0].ID, "2244994945")
}
