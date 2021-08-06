package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestTweetsByID() {
	tid := "1067094924124872705"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/"+tid,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Tweets.LookupByID(tid, TweetOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/"+tid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"1067094924124872705","text":"Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN"}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.LookupByID(tid, TweetOpts{})
	bc.Equal(*resp.Data.ID, tid)
}

func (bc *BCSuite) TestTweetsByIDs() {
	ids := "1261326399320715264,1278347468690915330"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Tweets.LookupByIDs(ids, TweetOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1261326399320715264","text":"Tune in to the @MongoDB @Twitch stream featuring our very own @suhemparack to learn about Twitter Developer Labs - starting now! https://t.co/fAWpYi3o5O"},{"id":"1278347468690915330","text":"Good news and bad news: nn2020 is half over"}]}`),
	)

	resp, _ := bc.Tw.Tweets.LookupByIDs(ids, TweetOpts{})
	bc.Equal(len(resp.Data), 2)
}
