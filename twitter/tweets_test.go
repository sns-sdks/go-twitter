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
