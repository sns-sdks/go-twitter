package twitter

import (
	"github.com/jarcoal/httpmock"
)

func (bc *BCSuite) TestGetListTweets() {
	lid := "84839422"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/lists/"+lid+"/tweets",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.GetListTweets(lid, ListTweetsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/lists/"+lid+"/tweets",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"author_id":"2244994945","created_at":"2018-11-26T16:37:10.000Z","text":"Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN","id":"1067094924124872705"}],"includes":{"users":[{"verified":true,"username":"TwitterDev","id":"2244994945","name":"Twitter Dev"}]},"meta":{"result_count":1}}`,
		),
	)

	resp, _ := bc.Tw.Lists.GetListTweets(lid, ListTweetsOpts{TweetOpts: TweetOpts{Expansions: "author_id", UserFields: "verified"}})
	bc.Equal(*resp.Data[0].ID, "1067094924124872705")
	bc.Equal(*resp.Includes.Users[0].ID, "2244994945")
}
