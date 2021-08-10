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
		{TweetsCounts{TweetCount: Int(1)}, `twitter.TweetsCounts{TweetCount:1}`},
		{BlockingStatus{Blocking: Bool(true)}, `twitter.BlockingStatus{Blocking:true}`},
		{FollowingStatus{Following: Bool(true)}, `twitter.FollowingStatus{Following:true}`},
		{MutingStatus{Muting: Bool(true)}, `twitter.MutingStatus{Muting:true}`},
		{LikedStatus{Liked: Bool(true)}, `twitter.LikedStatus{Liked:true}`},
		{RetweetedStatus{Retweeted: Bool(true)}, `twitter.RetweetedStatus{Retweeted:true}`},
		{APIError{Title: "error"}, `twitter.APIError{ClientID:"", RequiredEnrollment:"", RegistrationUrl:"", Title:"error", Detail:"", Reason:"", Type:"", Status:0}`},
		{AuthorizationAPP{ConsumerKey: "123", ConsumerSecret: ""}, `twitter.AuthorizationAPP{ConsumerKey:"123", ConsumerSecret:"", CallbackURL:"", AccessTokenKey:"", AccessTokenSecret:"", RequestSecret:""}`},
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
	}

	for i, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			t.Errorf("%d. String() => %q, want %q", i, tt.in, tt.out)
		}
	}
}
