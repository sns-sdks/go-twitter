package twitter

/*
	Bookmarks api collection.
*/

// GetBookmarksOpts specifies the parameters for user get bookmark tweets.
type GetBookmarksOpts struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	TweetOpts
}

// GetBookmarks Allows you to get information about an authenticated user’s 800 most recent bookmarked Tweets.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/get-users-id-bookmarks
func (r *TweetResource) GetBookmarks(id string, args GetBookmarksOpts) (*TweetsResp, *APIError) {
	path := "/users/" + id + "/bookmarks"

	resp := new(TweetsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// bookmarkTweetOpts specifies the parameters for bookmark tweet
type bookmarkTweetOpts struct {
	TweetID string `json:"tweet_id"`
}

// BookmarkedStatus represents the status for bookmark tweet
type BookmarkedStatus struct {
	Bookmarked *bool `json:"bookmarked,omitempty"`
}

func (s BookmarkedStatus) String() string {
	return Stringify(s)
}

// BookmarkedResp represents the response for bookmark tweet
type BookmarkedResp struct {
	Data *BookmarkedStatus `json:"data,omitempty"`
}

func (r BookmarkedResp) String() string {
	return Stringify(r)
}

// BookmarkTweet Causes the user ID of an authenticated user identified in the path parameter to Bookmark the target Tweet provided in the request body.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/post-users-id-bookmarks
func (r *TweetResource) BookmarkTweet(id, tweetID string) (*BookmarkedResp, *APIError) {
	path := "/users/" + id + "/bookmarks"
	postArgs := bookmarkTweetOpts{TweetID: tweetID}

	resp := new(BookmarkedResp)
	err := r.Cli.DoPost(path, postArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BookmarkTweetRemove Allows a user or authenticated user ID to remove a Bookmark of a Tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/delete-users-id-bookmarks-tweet_id
func (r *TweetResource) BookmarkTweetRemove(id, tweetID string) (*BookmarkedResp, *APIError) {
	path := "/users/" + id + "/bookmarks/" + tweetID

	resp := new(BookmarkedResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
