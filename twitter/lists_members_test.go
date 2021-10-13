package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestAddListMember() {
	lid := "1448302476780871685"
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/lists/"+lid+"/members",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.AddListMember(lid, uid)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/lists/"+lid+"/members",
		httpmock.NewStringResponder(
			200,
			`{"data":{"is_member":true}}`,
		),
	)

	resp, _ := bc.Tw.Lists.AddListMember(lid, uid)
	bc.Equal(*resp.Data.IsMember, true)
}

func (bc *BCSuite) TestRemoveListMember() {
	lid := "1448302476780871685"
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/lists/"+lid+"/members/"+uid,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.RemoveListMember(lid, uid)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/lists/"+lid+"/members/"+uid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"is_member":false}}`,
		),
	)

	resp, _ := bc.Tw.Lists.RemoveListMember(lid, uid)
	bc.Equal(*resp.Data.IsMember, false)
}
