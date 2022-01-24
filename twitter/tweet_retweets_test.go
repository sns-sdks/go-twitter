package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestRetweetedBy() {
	tid := "1354143047324299264"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/"+tid+"/retweeted_by",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Tweets.GetRetweetedBy(tid, RetweetedByOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/"+tid+"/retweeted_by",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1065249714214457345","name":"Spaces","username":"TwitterSpaces"},{"id":"783214","name":"Twitter","username":"Twitter"},{"id":"1526228120","name":"Twitter Data","username":"TwitterData"},{"id":"2244994945","name":"Twitter Dev","username":"TwitterDev"},{"id":"6253282","name":"Twitter API","username":"TwitterAPI"}],"meta":{"result_count":5}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.GetRetweetedBy(tid, RetweetedByOpts{})
	bc.Equal(len(resp.Data), 5)
}

func (uc *UCSuite) TestCreateRetweet() {
	uid := "123456789"
	tweetID := "1228393702244134912"

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/retweets",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.RetweetCreate(uid, tweetID)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/retweets",
		httpmock.NewStringResponder(
			200,
			`{"data":{"retweeted":true}}`,
		),
	)

	resp, _ := uc.Tw.Tweets.RetweetCreate(uid, tweetID)
	uc.Equal(*resp.Data.Retweeted, true)
}

func (uc *UCSuite) TestDestroyRetweet() {
	uid := "123456789"
	tweetID := "1228393702244134912"

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/retweets/"+tweetID,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.RetweetDestroy(uid, tweetID)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/retweets/"+tweetID,
		httpmock.NewStringResponder(
			200,
			`{"data":{"retweeted":false}}`,
		),
	)

	resp, _ := uc.Tw.Tweets.RetweetDestroy(uid, tweetID)
	uc.Equal(*resp.Data.Retweeted, false)
}
