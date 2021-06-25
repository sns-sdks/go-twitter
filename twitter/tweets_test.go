package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestTweetsByID() {
	tid := "1067094924124872705"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/"+tid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"id":"1067094924124872705","text":"Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN"}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.LookupByID(tid, TweetParams{})
	bc.Equal(*resp.Data.ID, tid)
}

func (bc *BCSuite) TestTweetsByIDs() {
	params := TweetParams{IDs: "1261326399320715264,1278347468690915330"}

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1261326399320715264","text":"Tune in to the @MongoDB @Twitch stream featuring our very own @suhemparack to learn about Twitter Developer Labs - starting now! https://t.co/fAWpYi3o5O"},{"id":"1278347468690915330","text":"Good news and bad news: nn2020 is half over"}]}`),
	)

	resp, _ := bc.Tw.Tweets.LookupByIDs(params)
	bc.Equal(len(resp.Data), 2)
}

func (bc *BCSuite) TestTweetsTimelines() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/tweets",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1338971066773905408","text":"💡 Using Twitter data for academic research? Join our next livestream this Friday @ 9am PT on https://t.co/GrtBOXh5Y1!n n@SuhemParack will show how to get started with recent search &amp; filtered stream endpoints on the #TwitterAPI v2, the new Tweet payload, annotations, &amp; more. https://t.co/IraD2Z7wEg"},{"id":"1338923691497959425","text":"📈 Live now with @jessicagarson and @i_am_daniele! https://t.co/Y1AFzsTTxb"},{"id":"1337498609819021312","text":"Thanks to everyone who tuned in today to make music with the #TwitterAPI!nnNext week on Twitch - @iamdaniele and @jessicagarson will show you how to integrate the #TwitterAPI and Google Sheets 📈. Tuesday, Dec 15th at 2pm ET. nnhttps://t.co/SQziic6eyp"},{"id":"1337464482654793740","text":"🎧💻 We're live! Tune in! 🎶 https://t.co/FSYP4rJdHr"},{"id":"1337122535188652033","text":"👂We want to hear what you think about our plans. As we continue to build our new product tracks, your feedback is essential to shaping the future of the Twitter API. Share your thoughts on this survey: https://t.co/dkIqFGPji7"},{"id":"1337122534173663235","text":"Is 2020 over yet?nDespite everything that happened this year, thousands of you still made the time to learn, play, and build incredible things on the new #TwitterAPI.nWe want to share some of your stories and give you a preview of what’s to come next year.nhttps://t.co/VpOKT22WgF"},{"id":"1336463248510623745","text":"🎧 Headphones on: watch @jessicagarson build an interactive app to write music using SuperCollider, Python, FoxDot, and the new Twitter API. Streaming Friday 1:30 ET on our new Twitch channel 🎶💻 https://t.co/SQziic6eyp"},{"id":"1334987486343299072","text":"console.log('Happy birthday, JavaScript!');"},{"id":"1334920270587584521","text":"Live now!nJoin the first ever @Twitch stream from TwitterDev https://t.co/x33fiVIi7B"},{"id":"1334564488884862976","text":"Before we release new #TwitterAPI endpoints, we let developers test drive a prototype of our intended design. @i_am_daniele takes you behind the scenes of an endpoint in the making. https://t.co/NNTDnciwNq"}],"meta":{"oldest_id":"1334564488884862976","newest_id":"1338971066773905408","result_count":10,"next_token":"7140dibdnow9c7btw3w29grvxfcgvpb9n9coehpk7xz5i"}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.GetTimelines(uid, TimelineParams{})
	bc.Equal(len(resp.Data), 10)
	bc.Equal(*resp.Data[0].ID, "1338971066773905408")
	bc.Equal(*resp.Meta.ResultCount, 10)
}

func (bc *BCSuite) TestTweetsMentions() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/mentions",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1375152598945312768","text":"@LeBraat @TwitterDev @LeGuud There's enough @twitterdev love to go around, @LeBraat"},{"id":"1375152449594523649","text":"Apparently I'm not the only one (of my test accounts) that loves @TwitterDev nn@LeStaache @LeGuud"},{"id":"1375152043455873027","text":"Can I join this @twitterdev love party?!"},{"id":"1375151947360174082","text":"I love me some @twitterdev too!"},{"id":"1375151827189137412","text":"This is a test, but also a good excuse to express my love for @TwitterDev 😍"}],"meta":{"oldest_id":"1375151827189137412","newest_id":"1375152598945312768","result_count":5,"next_token":"7140dibdnow9c7btw3w3y5b6jqjnk3k4g5zkmfkvqwa42"}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.GetMentions(uid, MentionParams{})
	bc.Equal(len(resp.Data), 5)
	bc.Equal(*resp.Data[0].ID, "1375152598945312768")
}

func (bc *BCSuite) TestLikingUsers() {
	tid := "1354143047324299264"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/"+tid+"/liking_users",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1065249714214457345","name":"Spaces","username":"TwitterSpaces"},{"id":"783214","name":"Twitter","username":"Twitter"},{"id":"1526228120","name":"Twitter Data","username":"TwitterData"},{"id":"2244994945","name":"Twitter Dev","username":"TwitterDev"},{"id":"6253282","name":"Twitter API","username":"TwitterAPI"}]}`,
		),
	)

	resp, _ := bc.Tw.Tweets.GetLikingUsers(tid, LikingUserPrams{})
	bc.Equal(len(resp.Data), 5)
}

func (bc *BCSuite) TestLikedTweets() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/liked_tweets",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1362449997430542337","text":"Honored to be the first developer to be featured in @TwitterDev's love fest 🥰♥️😍 https://t.co/g8TsPoZsij"},{"id":"1365416026435854338","text":"We're so happy for our Official Partner @Brandwatch and their big news. https://t.co/3DwWBNSq0o https://t.co/bDUGbgPkKO"},{"id":"1296487407475462144","text":"Check out this feature on @TwitterDev to learn more about how we're mining social media data to make sense of this evolving #publichealth crisis https://t.co/sIFLXRSvEX."},{"id":"1294346980072624128","text":"I awake from five years of slumber https://t.co/OEPVyAFcfB"},{"id":"1283153843367206912","text":"@wongmjane Wish we could tell you more, but I’m only a teapot 👀"}]}`,
		),
	)
	resp, _ := bc.Tw.Tweets.GetLikedTweets(uid, LikedTweetParams{})
	bc.Equal(len(resp.Data), 5)
}

func (bc *BCSuite) TestSearchRecentTweets() {
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/tweets/search/recent",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1373001119480344583","text":"Looking to get started with the Twitter API but new to APIs in general? @jessicagarson will walk you through everything you need to know in APIs 101 session. She’ll use examples using our v2 endpoints, Tuesday, March 23rd at 1 pm EST.nnJoin us on Twitchnhttps://t.co/GrtBOXyHmB"},{"id":"1372627771717869568","text":"Thanks to everyone who joined and made today a great session! 🙌 nnIf weren't able to attend, we've got you covered. Academic researchers can now sign up for office hours for help using the new product track. See how you can sign up, here 👇nhttps://t.co/duIkd27lPx https://t.co/AP9YY4F8FG"},{"id":"1367519323925843968","text":"Meet Aviary, a modern client for iOS 14 built using the new Twitter API. It has a beautiful UI and great widgets to keep you up to date with the latest Tweets. https://t.co/95cbd253jK"},{"id":"1366832168333234177","text":"The new #TwitterAPI provides the ability to build the Tweet payload with the fields that you want. nnIn this tutorial @suhemparack explains how to build the new Tweet payload and how it compares with the old Tweet payload in v1.1 👇 https://t.co/eQZulq4Ik3"},{"id":"1364984313154916352","text":"“I was heading to a design conference in New York and wanted to meet new people,” recalls @aaronykng, creator of @flocknet. “There wasn't an easy way to see all of the designers in my network, so I built one.” Making things like this opened the doors for him to the tech industry."},{"id":"1364275610764201984","text":"If you're newly approved for the Academic Research product track, our next stream is for you.nnThis Thursday, February 25th at 10AM PST @suhemparack will demo how academics can use this track to get started with the new #TwitterAPInnJoin us on Twitch! 👀nhttps://t.co/SQziibOD9P"}],"meta":{"newest_id":"1373001119480344583","oldest_id":"1364275610764201984","result_count":6}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.SearchRecent(SearchTweetsParams{Query: "from%3Atwitterdev%20new%20-is%3Aretweet", MaxResults: 10})
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

	resp, _ := bc.Tw.Tweets.SearchAll(SearchTweetsParams{Query: "from%3Atwitterdev%20new%20-is%3Aretweet", MaxResults: 10})
	bc.Equal(len(resp.Data), 6)
	bc.Equal(*resp.Data[0].ID, "1373001119480344583")
}
