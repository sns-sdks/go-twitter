package twitter

import "github.com/jarcoal/httpmock"

func (uc *UCSuite) TestTweetCreate() {
	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/tweets",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.TweetCreate(CreateTweetOpts{Text: "test tweet"})
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/tweets",
		httpmock.NewStringResponder(
			201,
			`{"data":{"id":"1457675300473552900","text":"test tweet"}}`,
		),
	)

	resp, _ := uc.Tw.Tweets.TweetCreate(CreateTweetOpts{Text: "test tweet"})
	uc.Equal(*resp.Data.ID, "1457675300473552900")
}

func (uc *UCSuite) TestTweetRemove() {
	tweetID := "1457675300473552900"

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/tweets/"+tweetID,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.TweetRemove(tweetID)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/tweets/"+tweetID,
		httpmock.NewStringResponder(
			200,
			`{"data":{"deleted":true}}`,
		),
	)
	resp, _ := uc.Tw.Tweets.TweetRemove(tweetID)
	uc.Equal(*resp.Data.Deleted, true)
}
