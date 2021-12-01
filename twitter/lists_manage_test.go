package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestListCreate() {
	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/lists",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.ListCreate(ManageListOpts{Name: "test v2 create list"})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/lists",
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"1441162269824405510","name":"test v2 create list"}}`,
		),
	)

	resp, _ := bc.Tw.Lists.ListCreate(ManageListOpts{Name: "test v2 create list"})
	bc.Equal(*resp.Data.ID, "1441162269824405510")
}

func (bc *BCSuite) TestListDelete() {
	listID := "1441162269824405510"

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/lists/"+listID,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.ListDelete(listID)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/lists/"+listID,
		httpmock.NewStringResponder(
			200,
			`{"data":{"deleted":true}}`,
		),
	)

	resp, _ := bc.Tw.Lists.ListDelete(listID)
	bc.Equal(*resp.Data.Deleted, true)
}

func (bc *BCSuite) TestListUpdate() {
	listID := "1441162269824405510"

	httpmock.RegisterResponder(
		HttpPut, Baseurl+"/lists/"+listID,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.ListUpdate(listID, ManageListOpts{Name: "test v2 update list"})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPut, Baseurl+"/lists/"+listID,
		httpmock.NewStringResponder(
			200,
			`{"data":{"updated":true}}`,
		),
	)

	resp, _ := bc.Tw.Lists.ListUpdate(listID, ManageListOpts{Name: "test v2 update list"})
	bc.Equal(*resp.Data.Updated, true)
}
