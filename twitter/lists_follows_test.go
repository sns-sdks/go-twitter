package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestFollowList() {
	lid := "1441162269824405510"
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/followed_lists",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.FollowList(uid, lid)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/users/"+uid+"/followed_lists",
		httpmock.NewStringResponder(
			200,
			`{"data":{"following":true}}`,
		),
	)

	resp, _ := bc.Tw.Lists.FollowList(uid, lid)
	bc.Equal(*resp.Data.Following, true)
}

func (bc *BCSuite) TestRemoveFollowedList() {
	lid := "1441162269824405510"
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/followed_lists/"+lid,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.RemoveFollowedList(uid, lid)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/users/"+uid+"/followed_lists/"+lid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"following":false}}`,
		),
	)

	resp, _ := bc.Tw.Lists.RemoveFollowedList(uid, lid)
	bc.Equal(*resp.Data.Following, false)
}

func (bc *BCSuite) TestGetListFollowers() {
	lid := "84839422"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/lists/"+lid+"/followers",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.GetListFollowers(lid, ListFollowersOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/lists/"+lid+"/followers",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"username":"alanbenlee","name":"Alan Lee","created_at":"2020-11-06T22:56:55.000Z","id":"1324848235714736129"},{"pinned_tweet_id":"1452599033625657359","username":"xo_chong","name":"Wilson Chong","created_at":"2020-11-16T15:31:22.000Z","id":"1328359963937259520"},{"username":"SumiraNazir","name":"Sumira Nazir","created_at":"2021-10-22T18:02:31.000Z","id":"1451609880113070085"},{"pinned_tweet_id":"1442182396523257861","username":"JustBorek","name":"Bořek Šindelka(he/him)","created_at":"2021-07-27T16:16:23.000Z","id":"1420055293082415107"},{"username":"w22ZccksRpafZAx","name":"金井憲司","created_at":"2021-06-27T13:09:10.000Z","id":"1409136449803325441"}],"includes":{"tweets":[{"created_at":"2021-10-25T11:32:52.000Z","id":"1452599033625657359","text":"https://t.co/aEuBQLXeuL"},{"created_at":"2021-09-26T17:40:52.000Z","id":"1442182396523257861","text":"Yes couple of days back nI want to kill my self I'm still here because of some amazing people please share this is important to talk about #mentalhealth @JustBorek #wheelchair #DisabilityTwitter #MedTwitter @heatherpsyd @Tweetinggoddess @NashaterS @msfatale @castleDD https://t.co/9hkSPV9NB1"}]},"meta":{"result_count":5,"next_token":"1714209892546977900"}}`,
		),
	)

	resp, _ := bc.Tw.Lists.GetListFollowers(lid, ListFollowersOpts{MaxResults: 5, UserOpts: UserOpts{Expansions: "pinned_tweet_id", UserFields: "created_at", TweetFields: "created_at"}})
	bc.Equal(*resp.Data[0].ID, "1324848235714736129")
	bc.Equal(*resp.Includes.Tweets[0].ID, "1452599033625657359")
}

func (bc *BCSuite) TestGetUserFollowedLists() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/followed_lists",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.GetUserFollowedLists(uid, FollowedListsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/followed_lists",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"follower_count":123,"id":"1630685563471","name":"Test List","owner_id":"1324848235714736129"}],"includes":{"users":[{"username":"alanbenlee","id":"1324848235714736129","created_at":"2009-08-28T18:30:45.000Z","name":"Alan Lee"}]},"meta":{"result_count":1}}`,
		),
	)

	resp, _ := bc.Tw.Lists.GetUserFollowedLists(uid, FollowedListsOpts{ListOpts: ListOpts{Expansions: "owner_id", ListFields: "follower_count", UserFields: "username"}})
	bc.Equal(*resp.Data[0].ID, "1630685563471")
	bc.Equal(*resp.Includes.Users[0].ID, "1324848235714736129")
}
