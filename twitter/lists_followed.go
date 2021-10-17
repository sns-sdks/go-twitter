package twitter

/*
	Lists followed include api for user follow or unfollow a list
*/

// ListFollowingStatus represents the status for user follow a list
type ListFollowingStatus struct {
	Following *bool `json:"following,omitempty"`
}

func (f ListFollowingStatus) String() string {
	return Stringify(f)
}

// ListFollowingResp represents the response for user follow a list
type ListFollowingResp struct {
	Data *ListFollowingStatus `json:"data,omitempty"`
}

func (f ListFollowingResp) String() string {
	return Stringify(f)
}

// followListOpts Inline parameters for follow a list
type followListOpts struct {
	ListID string `json:"list_id"`
}

// FollowList Enables the authenticated user to follow a List.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-users-id-followed-lists
func (r *ListsResource) FollowList(id, ListID string) (*ListFollowingResp, *APIError) {
	path := "/users/" + id + "/followed_lists"
	args := followListOpts{ListID: ListID}

	resp := new(ListFollowingResp)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RemoveFollowList Enables the authenticated user to unfollow a List.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-users-id-followed-lists-list_id
func (r *ListsResource) RemoveFollowList(id, ListID string) (*ListFollowingResp, *APIError) {
	path := "/users/" + id + "/followed_lists/" + ListID

	resp := new(ListFollowingResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
