package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestGetBookmarks() {
	uid := ""
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/bookmarks",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Tweets.GetBookmarks(uid, GetBookmarksOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/bookmarks",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"created_at":"2021-02-18T17:12:47.000Z","source":"Twitter Web App","id":"1362449997430542337","text":"Honored to be the first developer to be featured in @TwitterDev's love fest 🥰♥️😍 https://t.co/g8TsPoZsij"},{"created_at":"2021-02-26T21:38:43.000Z","source":"Twitter Web App","id":"1365416026435854338","text":"We're so happy for our Official Partner @Brandwatch and their big news. https://t.co/3DwWBNSq0o https://t.co/bDUGbgPkKO"},{"created_at":"2020-08-20T16:41:00.000Z","source":"Twitter Web App","id":"1296487407475462144","text":"Check out this feature on @TwitterDev to learn more about how we're mining social media data to make sense of this evolving #publichealth crisis https://t.co/sIFLXRSvEX."},{"created_at":"2020-08-14T18:55:42.000Z","source":"Twitter for Android","id":"1294346980072624128","text":"I awake from five years of slumber https://t.co/OEPVyAFcfB"},{"created_at":"2020-07-14T21:38:10.000Z","source":"Twitter for  iPhone","id":"1283153843367206912","text":"@wongmjane Wish we could tell you more, but I’m only a teapot 👀"}],"meta":{"result_count":5,"next_token":"zldjwdz3w6sba13nbs0mbravfipbtqvbiqplg9h0p4k"}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.GetBookmarks(uid, GetBookmarksOpts{MaxResults: 5, TweetOpts: TweetOpts{Expansions: "attachments.media_keys", MediaFields: "type,duration_ms"}})
	bc.Equal(len(resp.Data), 5)
	bc.Equal(*resp.Data[0].ID, "1362449997430542337")
	bc.Equal(*resp.Meta.ResultCount, 5)
}

func (uc *UCSuite) TestBookmarkTweet() {
	uid := "2244994945"
	tweetID := "1228393702244134912"

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/bookmarks",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.BookmarkTweet(uid, tweetID)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/bookmarks",
		httpmock.NewStringResponder(
			200,
			`{"data":{"bookmarked":true}}`,
		),
	)

	resp, _ := uc.Tw.Tweets.BookmarkTweet(uid, tweetID)
	uc.Equal(*resp.Data.Bookmarked, true)
}

func (uc *UCSuite) TestBookmarkTweetRemove() {
	uid := "2244994945"
	tweetID := "1228393702244134912"

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/bookmarks/"+tweetID,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.BookmarkTweetRemove(uid, tweetID)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/bookmarks/"+tweetID,
		httpmock.NewStringResponder(
			200,
			`{"data":{"bookmarked":false}}`,
		),
	)

	resp, _ := uc.Tw.Tweets.BookmarkTweetRemove(uid, tweetID)
	uc.Equal(*resp.Data.Bookmarked, false)
}
