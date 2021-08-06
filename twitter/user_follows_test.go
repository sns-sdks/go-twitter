package twitter

import "github.com/jarcoal/httpmock"

func (bc BCSuite) TestGetUserFollowers() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/followers",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Users.GetFollowers(uid, FollowsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/followers",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1356836061637468161","username":"Aman64539368","name":"Aman"},{"id":"1288755170738728961","username":"ziauddi45395796","name":"ziauddin"},{"id":"1326483734145720320","username":"destekportali","name":"Destek Portalı"},{"id":"2375888137","pinned_tweet_id":"1273689600490430464","username":"kaizhu717","name":"Kai Zhu"},{"id":"1345482756545982467","username":"Akpop81","name":"Akpop8"}],"includes":{"tweets":[{"id":"1273689600490430464","text":"Thrilled to see my first research project finally being published! It's a pleasure to work on an issue that I truly care about. https://t.co/UDq2yEvaj9"}]},"meta":{"result_count":5,"next_token":"H1411NSKD9R1EZZZ"}}`,
		),
	)

	resp, _ := bc.Tw.Users.GetFollowers(uid, FollowsOpts{})
	bc.Equal(len(resp.Data), 5)
	bc.Equal(*resp.Data[0].ID, "1356836061637468161")
}

func (bc BCSuite) TestGetUserFollowing() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/following",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Users.GetFollowing(uid, FollowsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/following",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1062359582","name":"Sheryl Klein Lavi🏴󠁧󠁢󠁳󠁣󠁴󠁿","pinned_tweet_id":"1257022566297812992","username":"TheSherylKlein"},{"id":"243665363","name":"Liliana Aidé Monge (She/Her/Ella)","pinned_tweet_id":"1356661433602252801","username":"mongeliliana"},{"id":"459860328","name":"julie✨","pinned_tweet_id":"1329503681855725568","username":"JulieMendoza206"},{"id":"273830767","name":"🄿🅄🅂🄷","username":"rahul_pushkarna"},{"id":"2240152338","name":"Max","username":"maxfwinter"}],"includes":{"tweets":[{"id":"1257022566297812992","text":"Twitter, the last 3 1/2 yrs are a gift my ❤️ will always cherish. From Infra Eng HRBP to 1st impression of culture, I feel proud I onboarded 3k Tweeps, transformed NHO to virtual #flightschoolremote & traveled the 🌎 to deliver trainings! Figaro & I will KIT about our adventure🏴󠁧󠁢󠁳󠁣󠁴󠁿 https://t.co/TPCvLYssWH"},{"id":"1356661433602252801","text":"Good Tech and Cause Show - Interview with me  (Lilian Aide Monge) from @WeSabio  (... https://t.co/jSnlBQyFDa via @YouTube"},{"id":"1329503681855725568","text":"Quick tip for folks that DM recruiters:\nA. Tell them WHY you are reaching out (keep this short & sweet 😊)\nB. do your own research the company & potential roles of interest (1-3 job links 🔗 in message)"}]},"meta":{"result_count":5,"next_token":"EFBO2DR4U531EZZZ"}}`,
		),
	)

	resp, _ := bc.Tw.Users.GetFollowing(uid, FollowsOpts{})
	bc.Equal(len(resp.Data), 5)
	bc.Equal(*resp.Data[0].ID, "1062359582")
}
