package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestSearchRecentTweets() {
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/search/recent",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1373001119480344583","text":"Looking to get started with the Twitter API but new to APIs in general? @jessicagarson will walk you through everything you need to know in APIs 101 session. She’ll use examples using our v2 endpoints, Tuesday, March 23rd at 1 pm EST.nnJoin us on Twitchnhttps://t.co/GrtBOXyHmB"},{"id":"1372627771717869568","text":"Thanks to everyone who joined and made today a great session! 🙌 nnIf weren't able to attend, we've got you covered. Academic researchers can now sign up for office hours for help using the new product track. See how you can sign up, here 👇nhttps://t.co/duIkd27lPx https://t.co/AP9YY4F8FG"},{"id":"1367519323925843968","text":"Meet Aviary, a modern client for iOS 14 built using the new Twitter API. It has a beautiful UI and great widgets to keep you up to date with the latest Tweets. https://t.co/95cbd253jK"},{"id":"1366832168333234177","text":"The new #TwitterAPI provides the ability to build the Tweet payload with the fields that you want. nnIn this tutorial @suhemparack explains how to build the new Tweet payload and how it compares with the old Tweet payload in v1.1 👇 https://t.co/eQZulq4Ik3"},{"id":"1364984313154916352","text":"“I was heading to a design conference in New York and wanted to meet new people,” recalls @aaronykng, creator of @flocknet. “There wasn't an easy way to see all of the designers in my network, so I built one.” Making things like this opened the doors for him to the tech industry."},{"id":"1364275610764201984","text":"If you're newly approved for the Academic Research product track, our next stream is for you.nnThis Thursday, February 25th at 10AM PST @suhemparack will demo how academics can use this track to get started with the new #TwitterAPInnJoin us on Twitch! 👀nhttps://t.co/SQziibOD9P"}],"meta":{"newest_id":"1373001119480344583","oldest_id":"1364275610764201984","result_count":6}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.SearchRecent(TweetSearchOpts{Query: "from%3Atwitterdev%20new%20-is%3Aretweet", MaxResults: 10})
	bc.Equal(len(resp.Data), 6)
	bc.Equal(*resp.Data[0].ID, "1373001119480344583")
}

func (bc *BCSuite) TestSearchAllTweets() {
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/search/all",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1373001119480344583","text":"Looking to get started with the Twitter API but new to APIs in general? @jessicagarson will walk you through everything you need to know in APIs 101 session. She’ll use examples using our v2 endpoints, Tuesday, March 23rd at 1 pm EST.nnJoin us on Twitchnhttps://t.co/GrtBOXyHmB"},{"id":"1372627771717869568","text":"Thanks to everyone who joined and made today a great session! 🙌 nnIf weren't able to attend, we've got you covered. Academic researchers can now sign up for office hours for help using the new product track. See how you can sign up, here 👇nhttps://t.co/duIkd27lPx https://t.co/AP9YY4F8FG"},{"id":"1367519323925843968","text":"Meet Aviary, a modern client for iOS 14 built using the new Twitter API. It has a beautiful UI and great widgets to keep you up to date with the latest Tweets. https://t.co/95cbd253jK"},{"id":"1366832168333234177","text":"The new #TwitterAPI provides the ability to build the Tweet payload with the fields that you want. nnIn this tutorial @suhemparack explains how to build the new Tweet payload and how it compares with the old Tweet payload in v1.1 👇 https://t.co/eQZulq4Ik3"},{"id":"1364984313154916352","text":"“I was heading to a design conference in New York and wanted to meet new people,” recalls @aaronykng, creator of @flocknet. “There wasn't an easy way to see all of the designers in my network, so I built one.” Making things like this opened the doors for him to the tech industry."},{"id":"1364275610764201984","text":"If you're newly approved for the Academic Research product track, our next stream is for you.nnThis Thursday, February 25th at 10AM PST @suhemparack will demo how academics can use this track to get started with the new #TwitterAPInnJoin us on Twitch! 👀nhttps://t.co/SQziibOD9P"}],"meta":{"newest_id":"1373001119480344583","oldest_id":"1364275610764201984","result_count":6}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.SearchAll(TweetSearchOpts{Query: "from%3Atwitterdev%20new%20-is%3Aretweet", MaxResults: 10})
	bc.Equal(len(resp.Data), 6)
	bc.Equal(*resp.Data[0].ID, "1373001119480344583")
}
