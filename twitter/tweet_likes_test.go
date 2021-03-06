package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestLikingUsers() {
	tid := "1354143047324299264"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/"+tid+"/liking_users",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Tweets.GetLikingUsers(tid, LikingUsersOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/"+tid+"/liking_users",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1065249714214457345","name":"Spaces","username":"TwitterSpaces"},{"id":"783214","name":"Twitter","username":"Twitter"},{"id":"1526228120","name":"Twitter Data","username":"TwitterData"},{"id":"2244994945","name":"Twitter Dev","username":"TwitterDev"},{"id":"6253282","name":"Twitter API","username":"TwitterAPI"}]}`,
		),
	)

	resp, _ := bc.Tw.Tweets.GetLikingUsers(tid, LikingUsersOpts{})
	bc.Equal(len(resp.Data), 5)
}

func (bc *BCSuite) TestLikedTweets() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/liked_tweets",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Tweets.GetLikedTweets(uid, LikedTweetsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/liked_tweets",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1362449997430542337","text":"Honored to be the first developer to be featured in @TwitterDev's love fest 🥰♥️😍 https://t.co/g8TsPoZsij"},{"id":"1365416026435854338","text":"We're so happy for our Official Partner @Brandwatch and their big news. https://t.co/3DwWBNSq0o https://t.co/bDUGbgPkKO"},{"id":"1296487407475462144","text":"Check out this feature on @TwitterDev to learn more about how we're mining social media data to make sense of this evolving #publichealth crisis https://t.co/sIFLXRSvEX."},{"id":"1294346980072624128","text":"I awake from five years of slumber https://t.co/OEPVyAFcfB"},{"id":"1283153843367206912","text":"@wongmjane Wish we could tell you more, but I’m only a teapot 👀"}]}`,
		),
	)
	resp, _ := bc.Tw.Tweets.GetLikedTweets(uid, LikedTweetsOpts{})
	bc.Equal(len(resp.Data), 5)
}

func (uc *UCSuite) TestCreateLike() {
	uid := "123456789"
	tweetID := "1228393702244134912"

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/likes",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.LikeCreate(uid, tweetID)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/likes",
		httpmock.NewStringResponder(
			200,
			`{"data":{"liked":true}}`,
		),
	)

	resp, _ := uc.Tw.Tweets.LikeCreate(uid, tweetID)
	uc.Equal(*resp.Data.Liked, true)
}

func (uc *UCSuite) TestDestroyLike() {
	uid := "123456789"
	tweetID := "1228393702244134912"

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/likes/"+tweetID,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.LikeDestroy(uid, tweetID)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/likes/"+tweetID,
		httpmock.NewStringResponder(
			200,
			`{"data":{"liked":false}}`,
		),
	)

	resp, _ := uc.Tw.Tweets.LikeDestroy(uid, tweetID)
	uc.Equal(*resp.Data.Liked, false)
}
