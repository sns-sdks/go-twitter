// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package twitter

import (
	"fmt"
	"testing"
)

func TestStringify(t *testing.T) {
	var nilPointer *string

	var tests = []struct {
		in  interface{}
		out string
	}{
		// basic types
		{"foo", `"foo"`},
		{123, `123`},
		{1.5, `1.5`},
		{false, `false`},
		{
			[]string{"a", "b"},
			`["a" "b"]`,
		},
		{
			struct {
				A []string
			}{nil},
			// nil slice is skipped
			`{}`,
		},
		{
			struct {
				A string
			}{"foo"},
			// structs not of a named type get no prefix
			`{A:"foo"}`,
		},

		// pointers
		{nilPointer, `<nil>`},
		{User{ID: String("123456"), Verified: Bool(true)}, `twitter.User{ID:"123456", Verified:true}`},
		{Tweet{ID: String("123456789"), PublicMetrics: &TweetPublicMetrics{LikeCount: Int(10213)}}, `twitter.Tweet{ID:"123456789", PublicMetrics:twitter.TweetPublicMetrics{LikeCount:10213}}`},
		{Space{ID: String("1DXxyRYNejbKM")}, `twitter.Space{ID:"1DXxyRYNejbKM"}`},
		{ComplianceJob{ID: String("1423095206576984067")}, `twitter.ComplianceJob{ID:"1423095206576984067"}`},
	}

	for i, tt := range tests {
		s := Stringify(tt.in)
		if s != tt.out {
			t.Errorf("%d. Stringify(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}

// Directly test the String() methods on various GitHub types. We don't do an
// exaustive test of all the various field types, since TestStringify() above
// takes care of that. Rather, we just make sure that Stringify() is being
// used to build the strings, which we do by verifying that pointers are
// stringified as their underlying value.
func TestString(t *testing.T) {
	var tests = []struct {
		in  interface{}
		out string
	}{
		{User{ID: String("1")}, `twitter.User{ID:"1"}`},
		{Tweet{ID: String("1234")}, `twitter.Tweet{ID:"1234"}`},
		{Media{Type: String("video")}, `twitter.Media{Type:"video"}`},
		{Place{ID: String("123")}, `twitter.Place{ID:"123"}`},
		{Poll{ID: String("123")}, `twitter.Poll{ID:"123"}`},
		{Space{ID: String("1DXxyRYNejbKM")}, `twitter.Space{ID:"1DXxyRYNejbKM"}`},
		{Topic{ID: String("123456"), Name: String("test")}, `twitter.Topic{ID:"123456", Name:"test"}`},
		{List{ID: String("1441162269824405510")}, `twitter.List{ID:"1441162269824405510"}`},
		{ComplianceJob{ID: String("1423095206576984067")}, `twitter.ComplianceJob{ID:"1423095206576984067"}`},
		{TweetsCounts{TweetCount: Int(1)}, `twitter.TweetsCounts{TweetCount:1}`},
		{BlockingStatus{Blocking: Bool(true)}, `twitter.BlockingStatus{Blocking:true}`},
		{FollowingStatus{Following: Bool(true)}, `twitter.FollowingStatus{Following:true}`},
		{MutingStatus{Muting: Bool(true)}, `twitter.MutingStatus{Muting:true}`},
		{LikedStatus{Liked: Bool(true)}, `twitter.LikedStatus{Liked:true}`},
		{RetweetedStatus{Retweeted: Bool(true)}, `twitter.RetweetedStatus{Retweeted:true}`},
		{HiddenStatus{Hidden: Bool(true)}, `twitter.HiddenStatus{Hidden:true}`},
		{ListDeletedStatus{Deleted: Bool(true)}, `twitter.ListDeletedStatus{Deleted:true}`},
		{ListUpdatedStatus{Updated: Bool(true)}, `twitter.ListUpdatedStatus{Updated:true}`},
		{ListMemberStatus{IsMember: Bool(true)}, `twitter.ListMemberStatus{IsMember:true}`},
		{ListFollowingStatus{Following: Bool(true)}, `twitter.ListFollowingStatus{Following:true}`},
		{ListPinnedStatus{Pinned: Bool(true)}, `twitter.ListPinnedStatus{Pinned:true}`},
		{TweetDeletedStatus{Deleted: Bool(true)}, `twitter.TweetDeletedStatus{Deleted:true}`},
		{APIError{Title: "error"}, `twitter.APIError{ClientID:"", RequiredEnrollment:"", RegistrationUrl:"", Title:"error", Detail:"", Reason:"", Type:"", Status:0, Errors:<nil>}`},
		{AuthorizationAPP{ConsumerKey: "123", ConsumerSecret: ""}, `twitter.AuthorizationAPP{ConsumerKey:"123", ConsumerSecret:"", CallbackURL:"", AccessTokenKey:"", AccessTokenSecret:"", RequestSecret:""}`},
		{OAuth2AuthorizationAPP{ClientID: "asfasfa123124"}, `twitter.OAuth2AuthorizationAPP{ClientID:"asfasfa123124", CallbackURL:""}`},
		{UserResp{Data: &User{ID: String("123456")}}, `twitter.UserResp{Data:twitter.User{ID:"123456"}}`},
		{UsersResp{Data: []*User{{ID: String("123456")}}}, `twitter.UsersResp{Data:[twitter.User{ID:"123456"}]}`},
		{TweetResp{Data: &Tweet{ID: String("123")}}, `twitter.TweetResp{Data:twitter.Tweet{ID:"123"}}`},
		{TweetsResp{Data: []*Tweet{{ID: String("123")}}}, `twitter.TweetsResp{Data:[twitter.Tweet{ID:"123"}]}`},
		{TweetsCountsResp{Data: []*TweetsCounts{{TweetCount: Int(1)}}}, `twitter.TweetsCountsResp{Data:[twitter.TweetsCounts{TweetCount:1}]}`},
		{FollowingResp{Data: &FollowingStatus{Following: Bool(true)}}, `twitter.FollowingResp{Data:twitter.FollowingStatus{Following:true}}`},
		{BlockingResp{Data: &BlockingStatus{Blocking: Bool(false)}}, `twitter.BlockingResp{Data:twitter.BlockingStatus{Blocking:false}}`},
		{MutingResp{Data: &MutingStatus{Muting: Bool(false)}}, `twitter.MutingResp{Data:twitter.MutingStatus{Muting:false}}`},
		{LikedResp{Data: &LikedStatus{Liked: Bool(false)}}, `twitter.LikedResp{Data:twitter.LikedStatus{Liked:false}}`},
		{RetweetedResp{Data: &RetweetedStatus{Retweeted: Bool(false)}}, `twitter.RetweetedResp{Data:twitter.RetweetedStatus{Retweeted:false}}`},
		{HiddenResp{Data: &HiddenStatus{Hidden: Bool(false)}}, `twitter.HiddenResp{Data:twitter.HiddenStatus{Hidden:false}}`},
		{SpaceResp{Data: &Space{ID: String("1eaKbnakjkkKX")}}, `twitter.SpaceResp{Data:twitter.Space{ID:"1eaKbnakjkkKX"}}`},
		{SpacesResp{Data: []*Space{{ID: String("1eaKbnakjkkKX"), State: String("scheduled")}}}, `twitter.SpacesResp{Data:[twitter.Space{ID:"1eaKbnakjkkKX", State:"scheduled"}]}`},
		{ComplianceJobResp{Data: &ComplianceJob{ID: String("1423095206576984067")}}, `twitter.ComplianceJobResp{Data:twitter.ComplianceJob{ID:"1423095206576984067"}}`},
		{ComplianceJobsResp{Data: []*ComplianceJob{{ID: String("1421185651106480129")}}}, `twitter.ComplianceJobsResp{Data:[twitter.ComplianceJob{ID:"1421185651106480129"}]}`},
		{ListResp{Data: &List{ID: String("1441162269824405510")}}, `twitter.ListResp{Data:twitter.List{ID:"1441162269824405510"}}`},
		{ListsResp{Data: []*List{{ID: String("1403322685870940160")}}}, `twitter.ListsResp{Data:[twitter.List{ID:"1403322685870940160"}]}`},
		{ListDeletedResp{Data: &ListDeletedStatus{Deleted: Bool(true)}}, `twitter.ListDeletedResp{Data:twitter.ListDeletedStatus{Deleted:true}}`},
		{ListUpdatedResp{Data: &ListUpdatedStatus{Updated: Bool(true)}}, `twitter.ListUpdatedResp{Data:twitter.ListUpdatedStatus{Updated:true}}`},
		{ListMemberResp{Data: &ListMemberStatus{IsMember: Bool(true)}}, `twitter.ListMemberResp{Data:twitter.ListMemberStatus{IsMember:true}}`},
		{ListFollowingResp{Data: &ListFollowingStatus{Following: Bool(true)}}, `twitter.ListFollowingResp{Data:twitter.ListFollowingStatus{Following:true}}`},
		{ListPinnedResp{Data: &ListPinnedStatus{Pinned: Bool(true)}}, `twitter.ListPinnedResp{Data:twitter.ListPinnedStatus{Pinned:true}}`},
		{TweetDeletedResp{Data: &TweetDeletedStatus{Deleted: Bool(true)}}, `twitter.TweetDeletedResp{Data:twitter.TweetDeletedStatus{Deleted:true}}`},
	}

	for i, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			t.Errorf("%d. String() => %q, want %q", i, tt.in, tt.out)
		}
	}
}
