package twitter

import ent "go-twitter/twitter/entities"

/*
	Tweets include api for tweets
*/

type TweetResource Resource

func newTweetResource(cli *Client) *TweetResource {
	return &TweetResource{Cli: cli}
}

type Fields struct {
	Tweet      string `url:"tweet.fields,omitempty"`
	Expansions string `url:"expansions,omitempty"`
	Media      string `url:"media.fields,omitempty"`
	Place      string `url:"place.fields,omitempty"`
	Poll       string `url:"poll.fields,omitempty"`
	User       string `url:"user.fields,omitempty"`
}

type TweetParams struct {
	IDs string `url:"ids,omitempty"`
	Fields
}

type TimelineParams struct {
	Exclude         string `url:"exclude,omitempty"`
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	SinceID         string `url:"since_id,omitempty"`
	UntilID         string `url:"until_id,omitempty"`
	StartTime       string `url:"start_time,omitempty"`
	EndTime         string `url:"end_time,omitempty"`
	Fields
}

type MentionParams struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	SinceID         string `url:"since_id,omitempty"`
	UntilID         string `url:"until_id,omitempty"`
	StartTime       string `url:"start_time,omitempty"`
	EndTime         string `url:"end_time,omitempty"`
	Fields
}

type LikingUserPrams struct {
	Fields
}

type LikedTweetParams struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	Fields
}

type SearchTweetsParams struct {
	Query      string `url:"query"`
	StartTime  string `url:"start_time,omitempty"`
	EndTime    string `url:"end_time,omitempty"`
	SinceID    string `url:"since_id,omitempty"`
	UntilID    string `url:"until_id,omitempty"`
	MaxResults int    `url:"max_results,omitempty"`
	NextToken  string `url:"next_token,omitempty"`
	Fields
}

type TweetsCountsParams struct {
	Query       string `url:"query"`
	Granularity string `url:"granularity,omitempty"`
	StartTime   string `url:"start_time,omitempty"`
	EndTime     string `url:"end_time,omitempty"`
	SinceID     string `url:"since_id,omitempty"`
	UntilID     string `url:"until_id,omitempty"`
}

func (r *TweetResource) LookupByID(id string, params TweetParams) (*ent.TweetResp, *APIError) {
	path := Baseurl + "/tweets/" + id

	resp := new(ent.TweetResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) LookupByIDs(params TweetParams) (*ent.TweetsResp, *APIError) {
	path := Baseurl + "/tweets"

	resp := new(ent.TweetsResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) GetTimelines(id string, params TimelineParams) (*ent.TweetsResp, *APIError) {
	path := Baseurl + "/users/" + id + "/tweets"

	resp := new(ent.TweetsResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) GetMentions(id string, params MentionParams) (*ent.TweetsResp, *APIError) {
	path := Baseurl + "/users/" + id + "/mentions"

	resp := new(ent.TweetsResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) GetLikingUsers(id string, params LikingUserPrams) (*ent.UsersResp, *APIError) {
	path := Baseurl + "/tweets/" + id + "/liking_users"

	resp := new(ent.UsersResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) GetLikedTweets(id string, params LikedTweetParams) (*ent.TweetsResp, *APIError) {
	path := Baseurl + "/users/" + id + "/liked_tweets"

	resp := new(ent.TweetsResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) SearchRecent(params SearchTweetsParams) (*ent.TweetsResp, *APIError) {
	path := Baseurl + "/tweets/search/recent"

	resp := new(ent.TweetsResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) SearchAll(params SearchTweetsParams) (*ent.TweetsResp, *APIError) {
	path := Baseurl + "/tweets/search/all"

	resp := new(ent.TweetsResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) CountsRecent(params TweetsCountsParams) (*ent.TweetsCountsResp, *APIError) {
	path := Baseurl + "/tweets/counts/recent"

	resp := new(ent.TweetsCountsResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *TweetResource) CountsALL(params TweetsCountsParams) (*ent.TweetsCountsResp, *APIError) {
	path := Baseurl + "/tweets/counts/all"

	resp := new(ent.TweetsCountsResp)
	err := r.Cli.DoGet(path, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
