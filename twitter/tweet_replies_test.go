package twitter

import "github.com/jarcoal/httpmock"

func (uc *UCSuite) TestHideReply() {
	tweetID := "1228393702244134912"

	httpmock.RegisterResponder(
		HttpPut, Baseurl+"/tweets/"+tweetID+"/hidden",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.HideReply(tweetID)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPut, Baseurl+"/tweets/"+tweetID+"/hidden",
		httpmock.NewStringResponder(
			200,
			`{"data":{"hidden":true}}`,
		),
	)

	resp, _ := uc.Tw.Tweets.HideReply(tweetID)
	uc.Equal(*resp.Data.Hidden, true)
}

func (uc *UCSuite) TestHideReplyRemove() {
	tweetID := "1228393702244134912"

	httpmock.RegisterResponder(
		HttpPut, Baseurl+"/tweets/"+tweetID+"/hidden",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.HideReplyDestroy(tweetID)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPut, Baseurl+"/tweets/"+tweetID+"/hidden",
		httpmock.NewStringResponder(
			200,
			`{"data":{"hidden":false}}`,
		),
	)

	resp, _ := uc.Tw.Tweets.HideReplyDestroy(tweetID)
	uc.Equal(*resp.Data.Hidden, false)
}
