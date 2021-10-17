package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestFollowList() {
	lid := "1441162269824405510"
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/followed_lists",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.FollowList(uid, lid)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/followed_lists",
		httpmock.NewStringResponder(
			200,
			`{"data":{"following":true}}`,
		),
	)

	resp, _ := bc.Tw.Lists.FollowList(uid, lid)
	bc.Equal(*resp.Data.Following, true)
}

func (bc *BCSuite) TestRemoveFollowList() {
	lid := "1441162269824405510"
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/followed_lists/"+lid,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.RemoveFollowList(uid, lid)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/followed_lists/"+lid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"following":false}}`,
		),
	)

	resp, _ := bc.Tw.Lists.RemoveFollowList(uid, lid)
	bc.Equal(*resp.Data.Following, false)
}
