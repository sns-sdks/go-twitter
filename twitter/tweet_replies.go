package twitter

type hideReplyOpts struct {
	Hidden bool `json:"hidden"`
}

// HiddenStatus represents the status for hidden reply tweet
type HiddenStatus struct {
	Hidden *bool `json:"hidden,omitempty"`
}

func (h HiddenStatus) String() string {
	return Stringify(h)
}

// HiddenResp represents the response for hidden reply tweet
type HiddenResp struct {
	Data *HiddenStatus `json:"data,omitempty"`
}

func (h HiddenResp) String() string {
	return Stringify(h)
}

// HideReply Hide a reply to a Tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/hide-replies/api-reference/put-tweets-id-hidden
func (r *TweetResource) HideReply(id string) (*HiddenResp, *APIError) {
	path := Baseurl + "/tweets/" + id + "/hidden"
	postArgs := hideReplyOpts{Hidden: true}

	resp := new(HiddenResp)
	err := r.Cli.DoPut(path, postArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// HideReplyDestroy Remove hide a reply to a Tweet.
// Refer: https://developer.twitter.com/en/docs/twitter-api/tweets/hide-replies/api-reference/put-tweets-id-hidden
func (r *TweetResource) HideReplyDestroy(id string) (*HiddenResp, *APIError) {
	path := Baseurl + "/tweets/" + id + "/hidden"
	postArgs := hideReplyOpts{Hidden: false}

	resp := new(HiddenResp)
	err := r.Cli.DoPut(path, postArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
