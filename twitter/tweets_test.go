package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestTweetByID() {
	tid := "1067094924124872705"

	httpmock.RegisterResponder(
		HTTP_GET, BASEURL+"/tweets/"+tid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"1067094924124872705","text":"Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN"}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.LookupByID(tid, TweetParams{})
	bc.Equal(*resp.Data.ID, tid)
}
