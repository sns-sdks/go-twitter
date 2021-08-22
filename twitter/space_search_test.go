package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestSpacesSearch() {
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/search",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Spaces.Search(SpaceSearchOpts{Query: "hello", State: "live"})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/search",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"host_ids":["2244994945"],"id":"1DXxyRYNejbKM","state":"live","title":"hello world 👋"},{"host_ids":["6253282"],"id":"1nAJELYEEPvGL","state":"live","title":"Say hello to the Spaces endpoints"}],"meta":{"result_count":2}}`,
		),
	)

	resp, _ := bc.Tw.Spaces.Search(SpaceSearchOpts{Query: "hello", State: "live", SpaceOpts: SpaceOpts{SpaceFields: "title"}})
	bc.Equal(*resp.Data[0].ID, "1DXxyRYNejbKM")
	bc.Equal(len(resp.Data), 2)
	bc.Equal(*resp.Meta.ResultCount, 2)
}
