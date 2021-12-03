package twitter

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestAddListMember() {
	lid := "1448302476780871685"
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/lists/"+lid+"/members",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.AddListMember(lid, uid)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/lists/"+lid+"/members",
		httpmock.NewStringResponder(
			200,
			`{"data":{"is_member":true}}`,
		),
	)

	resp, _ := bc.Tw.Lists.AddListMember(lid, uid)
	bc.Equal(*resp.Data.IsMember, true)
}

func (bc *BCSuite) TestRemoveListMember() {
	lid := "1448302476780871685"
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/lists/"+lid+"/members/"+uid,
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.RemoveListMember(lid, uid)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/lists/"+lid+"/members/"+uid,
		httpmock.NewStringResponder(
			200,
			`{"data":{"is_member":false}}`,
		),
	)

	resp, _ := bc.Tw.Lists.RemoveListMember(lid, uid)
	bc.Equal(*resp.Data.IsMember, false)
}

func (bc *BCSuite) TestGetListMembers() {
	lid := "84839422"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/lists/"+lid+"/members",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.GetListMembers(lid, ListMembersOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/lists/"+lid+"/members",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"name":"Birdwatch","id":"1319036828964454402","username":"birdwatch","pinned_tweet_id":"1353789891348475905"},{"name":"Twitter Thailand","id":"1244731491088809984","username":"TwitterThailand"},{"name":"Twitter Retweets","id":"1194267639100723200","username":"TwitterRetweets"},{"name":"Twitter Able","id":"1168976680867762177","username":"TwitterAble"},{"name":"Spaces","id":"1065249714214457345","username":"TwitterSpaces","pinned_tweet_id":"1451239134798942208"}],"includes":{"tweets":[{"id":"1353789891348475905","text":"Want to help build a new community-driven approach to tackling misleading information? Join us — sign up for Birdwatch! nnhttps://t.co/FSsqNznPy1"},{"id":"1451239134798942208","text":"the time has arrived -- we’re now rolling out the ability for everyone on iOS and Android to host a Spacennif this is your first time hosting, welcome! here’s a refresher on how https://t.co/cLH8z0bocy"}]},"meta":{"result_count":5,"next_token":"5676935732641845249"}}`,
		),
	)

	resp, _ := bc.Tw.Lists.GetListMembers(lid, ListMembersOpts{MaxResults: 5, UserOpts: UserOpts{Expansions: "pinned_tweet_id", UserFields: "username"}})
	bc.Equal(*resp.Data[0].ID, "1319036828964454402")
	bc.Equal(*resp.Includes.Tweets[0].ID, "1353789891348475905")
}

func (bc *BCSuite) TestGetUserJoinedLists() {
	uid := "2244994945"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/list_memberships",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := bc.Tw.Lists.GetUserJoinedLists(uid, JoinedListsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/users/"+uid+"/list_memberships",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"follower_count":5,"id":"1451951974291689472","name":"Twitter","owner_id":"1227213680120479745"}],"includes":{"users":[{"name":"구돆","created_at":"2020-02-11T12:52:11.000Z","id":"1227213680120479745","username":"Follow__Y0U"}]},"meta":{"result_count":1}}`,
		),
	)

	resp, _ := bc.Tw.Lists.GetUserJoinedLists(uid, JoinedListsOpts{ListOpts: ListOpts{Expansions: "owner_id", ListFields: "follower_count", UserFields: "created_at"}})
	bc.Equal(*resp.Data[0].ID, "1451951974291689472")
	bc.Equal(*resp.Includes.Users[0].ID, "1227213680120479745")
}
