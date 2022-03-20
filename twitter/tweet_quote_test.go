package twitter

import "github.com/jarcoal/httpmock"

func (uc *UCSuite) TestGetQuoteTweets() {
	tweetID := "1409931481552543749"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/"+tweetID+"/quote_tweets",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.GetQuoteTweets(tweetID, GetQuoteTweetsOpts{})
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/"+tweetID+"/quote_tweets",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1495979553889697792","text":"RT @chris_bail: Twitter has created an entire course (with videos, code, and other materials) to help researchers learn how to collect data…"},{"id":"1486385372401737728","text":"RT @suhemparack: Super excited to share our course on Getting started with the #TwitterAPI v2 for academic researchnnIf you know students w…"},{"id":"1480954678447857669","text":"RT @suhemparack: Super excited to share our course on Getting started with the #TwitterAPI v2 for academic researchnnIf you know students w…"},{"id":"1480639272721940486","text":"RT @suhemparack: Super excited to share our course on Getting started with the #TwitterAPI v2 for academic researchnnIf you know students w…"},{"id":"1471614967207976961","text":"RT @chris_bail: Twitter has created an entire course (with videos, code, and other materials) to help researchers learn how to collect data…"},{"id":"1470423243513372679","text":"RT @suhemparack: Super excited to share our course on Getting started with the #TwitterAPI v2 for academic researchnnIf you know students w…"},{"id":"1469125403373568001","text":"RT @suhemparack: Super excited to share our course on Getting started with the #TwitterAPI v2 for academic researchnnIf you know students w…"},{"id":"1468633446935318529","text":"RT @suhemparack: Super excited to share our course on Getting started with the #TwitterAPI v2 for academic researchnnIf you know students w…"},{"id":"1438256410417143809","text":"RT @suhemparack: Super excited to share our course on Getting started with the #TwitterAPI v2 for academic researchnnIf you know students w…"},{"id":"1430934381829492746","text":"RT @suhemparack: Super excited to share our course on Getting started with the #TwitterAPI v2 for academic researchnnIf you know students w…"}],"meta":{"result_count":10,"next_token":"avdjwk0udyx6"}}`,
		),
	)

	resp, _ := uc.Tw.Tweets.GetQuoteTweets(tweetID, GetQuoteTweetsOpts{MaxResults: 10})
	uc.Equal(*resp.Data[0].ID, "1495979553889697792")
	uc.Equal(*resp.Meta.ResultCount, 10)
}
