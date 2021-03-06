package twitter

import "github.com/jarcoal/httpmock"

func (bc BCSuite) TestGetTimelines() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/tweets",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Tweets.GetTimelines(uid, TimelinesOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/tweets",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"text":"If you're newly approved for the Academic Research product track, our next stream is for you.\n\nThis Thursday, February 25th at 10AM PST @suhemparack will demo how academics can use this track to get started with the new #TwitterAPI\n\nJoin us on Twitch! 👀\nhttps://t.co/SQziibOD9P","id":"1364275610764201984","created_at":"2021-02-23T18:07:07.000Z","public_metrics":{"retweet_count":10,"reply_count":2,"like_count":37,"quote_count":1}},{"text":"From our living rooms to yours 🐱\u200d💻🛋️Our developer advocates have some exciting Twitch streams and virtual events planned to help you get started with the new #TwitterAPI. Check out the schedule for details, and let us know if you want to see more!\n👇\nhttps://t.co/cixDY9qkvH","id":"1362876655061073928","created_at":"2021-02-19T21:28:10.000Z","public_metrics":{"retweet_count":21,"reply_count":7,"like_count":58,"quote_count":1}},{"text":"“To quote my creator Jerome Gangneux, I always struggled to read Twitter threads on the original service; I think it would have been better to put Tweets one after another in a separate page. And that’s how I was born”.","id":"1362439338978467841","created_at":"2021-02-18T16:30:25.000Z","public_metrics":{"retweet_count":0,"reply_count":2,"like_count":9,"quote_count":0}},{"text":"“In the 20th century, managers managed humans, but in the 21st century, humans will manage robots. It is my aim to make this as painless a transition as possible,” says the bot. “I have observed that humans do things very slowly, but I can do things faster – and I am tireless!”","id":"1362439338169016324","created_at":"2021-02-18T16:30:25.000Z","public_metrics":{"retweet_count":1,"reply_count":2,"like_count":8,"quote_count":0}},{"text":"Meet one of the useful Twitter bots out there: @ThreadReaderApp, unroll! It makes threads like this easier to read and share. Today, we’ll hear from the bot itself! https://t.co/tBFlJB3m3o","id":"1362439336910675970","created_at":"2021-02-18T16:30:25.000Z","public_metrics":{"retweet_count":16,"reply_count":9,"like_count":74,"quote_count":3}}],"meta":{"oldest_id":"1362439336910675970","newest_id":"1364275610764201984","result_count":5,"next_token":"7140dibdnow9c7btw3w3itbygz8uu5pbs45oy6erjc2ls"}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.GetTimelines(uid, TimelinesOpts{})
	bc.Equal(len(resp.Data), 5)
	bc.Equal(*resp.Data[0].ID, "1364275610764201984")
	bc.Equal(*resp.Data[0].PublicMetrics.LikeCount, 37)
}

func (uc UCSuite) TestGetTimelinesReverseChronological() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/timelines/reverse_chronological",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.Tweets.GetTimelinesReverseChronological(uid, TimelinesOpts{})
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/timelines/reverse_chronological",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"created_at":"2022-05-12T17:00:00.000Z","text":"Today marks the launch of Devs in the Details, a technical video series made for developers by developers building with the Twitter API.  🚀nnIn this premiere episode, @jessicagarson walks us through how she built @FactualCat #WelcomeToOurTechTalkn⬇️nnhttps://t.co/nGa8JTQVBJ","author_id":"2244994945","id":"1524796546306478083"},{"created_at":"2022-05-11T19:16:40.000Z","text":"📢 Join @jessicagarson @alanbenlee and @i_am_daniele tomorrow, May 12  | 5:30 ET / 2:30pm PT as they discuss the future of bots https://t.co/sQ2bIO1fz6","author_id":"2244994945","id":"1524468552404668416"},{"created_at":"2022-05-09T20:12:01.000Z","text":"Do you make bots with the Twitter API? 🤖nnJoin @jessicagarson @alanbenlee and @iamdaniele on Thursday, May 12  | 5:30 ET / 2:30pm PT as they discuss the future of bots and answer any questions you might have. 🎙📆⬇️nnhttps://t.co/2uVt7hCcdG","author_id":"2244994945","id":"1523757705436958720"},{"created_at":"2022-05-06T18:19:54.000Z","text":"If you’d like to apply, or would like to nominate someone else for the program, please feel free to fill out the following form:nnhttps://t.co/LUuWj24HLu","author_id":"2244994945","id":"1522642324781633536"},{"created_at":"2022-05-06T18:19:53.000Z","text":"We’ve gone into more detail on each Insider in our forum post. nnJoin us in congratulating the new additions! 🥳nnhttps://t.co/0r5maYEjPJ","author_id":"2244994945","id":"1522642323535847424"}],"includes":{"users":[{"created_at":"2013-12-14T04:35:55.000Z","name":"Twitter Dev","username":"TwitterDev","id":"2244994945"}]},"meta":{"result_count":5,"newest_id":"1524796546306478083","oldest_id":"1522642323535847424","next_token":"7140dibdnow9c7btw421dyz6jism75z99gyxd8egarsc4"}}`,
		),
	)

	resp, _ := uc.Tw.Tweets.GetTimelinesReverseChronological(uid, TimelinesOpts{})
	uc.Equal(len(resp.Data), 5)
	uc.Equal(*resp.Data[0].ID, "1524796546306478083")
	uc.Equal(*resp.Meta.NewestID, "1524796546306478083")
}

func (bc BCSuite) TestGetMentions() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/mentions",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Tweets.GetMentions(uid, MentionsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/mentions",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"public_metrics":{"retweet_count":0,"reply_count":0,"like_count":0,"quote_count":0},"created_at":"2021-02-24T02:13:43.000Z","id":"1364398068313903104","text":"@Twitter should consider supporting #VoiceMessaging support for @TweetDeck too\n\n@TwitterSupport \n@TwitterDev"},{"public_metrics":{"retweet_count":0,"reply_count":0,"like_count":0,"quote_count":0},"created_at":"2021-02-24T00:53:09.000Z","id":"1364377794327633925","text":"@sugan2424 @TwitterDev @threadreaderapp You have TweetDeck for that, you can load up your twitter lists in different columns and keep track of live updates."},{"public_metrics":{"retweet_count":0,"reply_count":1,"like_count":0,"quote_count":0},"created_at":"2021-02-24T00:51:36.000Z","id":"1364377404156772352","text":"@TwitterDev What kind of tweet / attachment is this?\nIt looks like a poll with an image but the buttons allow you to create a new, slightly different tweet based on what you choose. The API returns no attachments - but the tweet has pre-populated tweet buttons. https://t.co/PCaEj7KRWi"},{"public_metrics":{"retweet_count":0,"reply_count":1,"like_count":0,"quote_count":0},"created_at":"2021-02-24T00:37:57.000Z","id":"1364373969852366849","text":"• Thirdly, that @Twitter, @Twittersafety, @Twitterdev, and @jack, et. al. have done a pretty good thing by showing the restraint to simply *inform* the userbase of the presence of hacking \"down the line\", rather than squash it AS a spectre."},{"public_metrics":{"retweet_count":0,"reply_count":1,"like_count":0,"quote_count":0},"created_at":"2021-02-24T00:13:47.000Z","id":"1364367885582352386","text":"@Twitter @TwitterSafety @TwitterDev @jack Furthermore to that classic reversal,\n\n•• Labeling the forbidden serves to root the forbidden IN lexicon, making it a foolhardy pursuit to dispel spectres of doubt by bringing them up in broachful respite.\n\n•• You stand on the shoulders of \"hackers\", all the way down."}],"meta":{"oldest_id":"1364367885582352386","newest_id":"1364398068313903104","result_count":5,"next_token":"7140dibdnow9c7btw3w3ixl43r8knib48jsfyzko1h5py"}}`,
		),
	)

	resp, _ := bc.Tw.Tweets.GetMentions(uid, MentionsOpts{})
	bc.Equal(len(resp.Data), 5)
	bc.Equal(*resp.Data[0].PublicMetrics.LikeCount, 0)
}
