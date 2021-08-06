package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestGetUserGetBlocking() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/blocking",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1065249714214457345","created_at":"2018-11-21T14:24:58.000Z","name":"Spaces","pinned_tweet_id":"1389270063807598594","description":"Twitter Spaces is where live audio conversations happen.","username":"TwitterSpaces"},{"id":"783214","created_at":"2007-02-20T14:35:54.000Z","name":"Twitter","description":"What's happening?!","username":"Twitter"},{"id":"1526228120","created_at":"2013-06-17T23:57:45.000Z","name":"Twitter Data","description":"Data-driven insights about notable moments and conversations from Twitter, Inc., plus tips and tricks to help you get the most out of Twitter data.","username":"TwitterData"},{"id":"2244994945","created_at":"2013-12-14T04:35:55.000Z","name":"Twitter Dev","pinned_tweet_id":"1354143047324299264","description":"The voice of the #TwitterDev team and your official source for updates, news, and events, related to the #TwitterAPI.","username":"TwitterDev"},{"id":"6253282","created_at":"2007-05-23T06:01:13.000Z","name":"Twitter API","pinned_tweet_id":"1293595870563381249","description":"Tweets about changes and service issues. Follow @TwitterDev for more.","username":"TwitterAPI"}],"includes":{"tweets":[{"id":"1389270063807598594","text":"now, everyone with 600 or more followers can host a Space.nnbased on what we've learned, these accounts are likely to have a good experience hosting because of their existing audience. before bringing the ability to create a Space to everyone, we’re focused on a few things. 🧵"},{"id":"1354143047324299264","text":"Academics are one of the biggest groups using the #TwitterAPI to research what’s happening. Their work helps make the world (&amp; Twitter) a better place, and now more than ever, we must enable more of it. nIntroducing 🥁 the Academic Research product track!nhttps://t.co/nOFiGewAV2"},{"id":"1293595870563381249","text":"Twitter API v2: Early Access releasednnToday we announced Early Access to the first endpoints of the new Twitter API!nn#TwitterAPI #EarlyAccess #VersionBump https://t.co/g7v3aeIbtQ"}]}}`,
		),
	)

	resp, _ := bc.Tw.Users.GetBlocking(uid, UserBlockingOpts{})
	bc.Equal(len(resp.Data), 5)
	bc.Equal(*resp.Includes.Tweets[0].ID, "1389270063807598594")
}
