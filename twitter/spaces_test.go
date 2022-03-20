package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestSpaceByID() {
	sid := "1DXxyRYNejbKM"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/"+sid,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Spaces.LookupByID(sid, SpaceOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/"+sid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"1DXxyRYNejbKM","state":"live"}}`,
		),
	)

	resp, _ := bc.Tw.Spaces.LookupByID(sid, SpaceOpts{})
	bc.Equal(*resp.Data.ID, sid)
	bc.Equal(*resp.Data.State, "live")
}

func (bc *BCSuite) TestSpacesByIDs() {
	ids := "1DXxyRYNejbKM,1nAJELYEEPvGL"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Spaces.LookupByIDs(ids, SpaceOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"host_ids":["2244994945"],"id":"1DXxyRYNejbKM","state":"live"},{"host_ids":["6253282"],"id":"1nAJELYEEPvGL","state":"scheduled"}]}`,
		),
	)

	resp, _ := bc.Tw.Spaces.LookupByIDs(ids, SpaceOpts{SpaceFields: "host_ids"})
	bc.Equal(*resp.Data[0].ID, "1DXxyRYNejbKM")
	bc.Equal(len(resp.Data), 2)
}

func (bc *BCSuite) TestSpacesByCreators() {
	userIDs := "2244994945,6253282"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/by/creator_ids",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Spaces.LookupByCreators(userIDs, SpaceOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/by/creator_ids",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"host_ids":["2244994945"],"id":"1DXxyRYNejbKM","state":"live"},{"host_ids":["6253282"],"id":"1nAJELYEEPvGL","state":"scheduled"}],"meta":{"result_count":2}}`,
		),
	)

	resp, _ := bc.Tw.Spaces.LookupByCreators(userIDs, SpaceOpts{SpaceFields: "host_ids"})
	bc.Equal(*resp.Data[0].ID, "1DXxyRYNejbKM")
	bc.Equal(*resp.Data[0].HostIDs[0], "2244994945")
	bc.Equal(len(resp.Data), 2)
}

func (bc *BCSuite) TestSpacesGetBuyers() {
	spaceID := "1DXxyRYNejbKM"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/"+spaceID+"/buyers",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Spaces.GetBuyers(spaceID, UserOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/"+spaceID+"/buyers",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"created_at":"2013-12-14T04:35:55.000Z","username":"TwitterDev","pinned_tweet_id":"1255542774432063488","id":"2244994945","name":"Twitter Dev"},{"created_at":"2007-02-20T14:35:54.000Z","username":"Twitter","pinned_tweet_id":"1274087687469715457","id":"783214","name":"Twitter"}],"includes":{"tweets":[{"created_at":"2020-04-29T17:01:38.000Z","text":"During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId","id":"1255542774432063488"},{"created_at":"2020-06-19T21:12:30.000Z","text":"📍 Minneapolisn🗣️ @FredTJoseph https://t.co/lNTOkyguG1","id":"1274087687469715457"}]}}`,
		),
	)

	resp, _ := bc.Tw.Spaces.GetBuyers(spaceID, UserOpts{Expansions: "pinned_tweet_id", UserFields: "created_at", TweetFields: "created_at"})
	bc.Equal(*resp.Data[0].ID, "2244994945")
	bc.Equal(*resp.Data[0].PinnedTweetID, "1255542774432063488")
	bc.Equal(*resp.Includes.Tweets[0].ID, "1255542774432063488")
}

func (bc *BCSuite) TestSpacesGetTweets() {
	spaceID := "1DXxyRYNejbKM"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/"+spaceID+"/tweets",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Spaces.GetTweets(spaceID, TweetOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/spaces/"+spaceID+"/tweets",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1389270063807598594","author_id":"1065249714214457345","text":"now, everyone with 600 or more followers can host a Space.nnbased on what we've learned, these accounts are likely to have a good experience hosting because of their existing audience. before bringing the ability to create a Space to everyone, we're focused on a few things. :thread:"},{"id":"1354143047324299264","author_id":"783214","text":"Academics are one of the biggest groups using the #TwitterAPI to research what's happening. Their work helps make the world (&amp; Twitter) a better place, and now more than ever, we must enable more of it. nIntroducing :drum_with_drumsticks: the Academic Research product track!nhttps://t.co/nOFiGewAV2"},{"id":"1293595870563381249","author_id":"783214","text":"Twitter API v2: Early Access releasednnToday we announced Early Access to the first endpoints of the new Twitter API!nn#TwitterAPI #EarlyAccess #VersionBump https://t.co/g7v3aeIbtQ"}],"includes":{"users":[{"id":"1065249714214457345","created_at":"2018-11-21T14:24:58.000Z","name":"Spaces","pinned_tweet_id":"1389270063807598594","description":"Twitter Spaces is where live audio conversations happen.","username":"TwitterSpaces"},{"id":"783214","created_at":"2007-02-20T14:35:54.000Z","name":"Twitter","description":"What's happening?!","username":"Twitter"},{"id":"1526228120","created_at":"2013-06-17T23:57:45.000Z","name":"Twitter Data","description":"Data-driven insights about notable moments and conversations from Twitter, Inc., plus tips and tricks to help you get the most out of Twitter data.","username":"TwitterData"}],"meta":{"result_count":3}}}`,
		),
	)

	resp, _ := bc.Tw.Spaces.GetTweets(spaceID, TweetOpts{Expansions: "author_id", UserFields: "created_at,description"})
	bc.Equal(*resp.Data[0].ID, "1389270063807598594")
	bc.Equal(*resp.Data[0].AuthorID, "1065249714214457345")
	bc.Equal(*resp.Includes.Users[0].ID, "1065249714214457345")
}
