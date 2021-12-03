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

// RemoveFollowedList Enables the authenticated user to unfollow a List.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-users-id-followed-lists-list_id
func (r *ListsResource) RemoveFollowedList(id, ListID string) (*ListFollowingResp, *APIError) {
	path := "/users/" + id + "/followed_lists/" + ListID

	resp := new(ListFollowingResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListFollowersOpts Specifies the parameters for get list followers
type ListFollowersOpts struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	UserOpts
}

// GetListFollowers Returns a list of users who are followers of the specified List.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/list-follows/api-reference/get-lists-id-followers
func (r *ListsResource) GetListFollowers(id string, args ListFollowersOpts) (*UsersResp, *APIError) {
	path := "/lists/" + id + "/followers"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FollowedListsOpts Specifies the parameters for get user followed lists.
type FollowedListsOpts struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	ListOpts
}

// GetUserFollowedLists Returns all Lists a specified user follows.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/list-follows/api-reference/get-users-id-followed_lists
func (r *ListsResource) GetUserFollowedLists(id string, args FollowedListsOpts) (*ListsResp, *APIError) {
	path := "/users/" + id + "/followed_lists"

	resp := new(ListsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
