package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestSpaceByID() {
	sid := "1DXxyRYNejbKM"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/"+sid,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Spaces.LookupByID(sid, SpaceOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/"+sid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"1DXxyRYNejbKM","state":"live"}}`,
		),
	)

	resp, _ := bc.Tw.Spaces.LookupByID(sid, SpaceOpts{})
	bc.Equal(*resp.Data.ID, sid)
	bc.Equal(*resp.Data.State, "live")
}

func (bc *BCSuite) TestSpacesByIDs() {
	ids := "1DXxyRYNejbKM,1nAJELYEEPvGL"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Spaces.LookupByIDs(ids, SpaceOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"host_ids":["2244994945"],"id":"1DXxyRYNejbKM","state":"live"},{"host_ids":["6253282"],"id":"1nAJELYEEPvGL","state":"scheduled"}]}`,
		),
	)

	resp, _ := bc.Tw.Spaces.LookupByIDs(ids, SpaceOpts{SpaceFields: "host_ids"})
	bc.Equal(*resp.Data[0].ID, "1DXxyRYNejbKM")
	bc.Equal(len(resp.Data), 2)
}

func (bc *BCSuite) TestSpacesByCreators() {
	userIDs := "2244994945,6253282"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/by/creator_ids",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Spaces.LookupByCreators(userIDs, SpaceOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/by/creator_ids",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"host_ids":["2244994945"],"id":"1DXxyRYNejbKM","state":"live"},{"host_ids":["6253282"],"id":"1nAJELYEEPvGL","state":"scheduled"}],"meta":{"result_count":2}}`,
		),
	)

	resp, _ := bc.Tw.Spaces.LookupByCreators(userIDs, SpaceOpts{SpaceFields: "host_ids"})
	bc.Equal(*resp.Data[0].ID, "1DXxyRYNejbKM")
	bc.Equal(*resp.Data[0].HostIDs[0], "2244994945")
	bc.Equal(len(resp.Data), 2)
}
