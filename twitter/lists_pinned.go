package twitter

/*
	Lists followed include api for user pinned or unpin a list
*/

// ListPinnedStatus represents the status for user pinned a list
type ListPinnedStatus struct {
	Pinned *bool `json:"pinned,omitempty"`
}

func (p ListPinnedStatus) String() string {
	return Stringify(p)
}

// ListPinnedResp represents the response for user pin a list
type ListPinnedResp struct {
	Data *ListPinnedStatus `json:"data,omitempty"`
}

func (p ListPinnedResp) String() string {
	return Stringify(p)
}

// pinListOpts Inline parameters for pin a list
type pinListOpts struct {
	ListID string `json:"list_id"`
}

// PinList Enables the authenticated user to pin a List.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-users-id-pinned-lists
func (r *ListsResource) PinList(id, ListID string) (*ListPinnedResp, *APIError) {
	path := "/users/" + id + "/pinned_lists"
	args := pinListOpts{ListID: ListID}

	resp := new(ListPinnedResp)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RemovePinnedList Enables the authenticated user to unfollow a List.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-users-id-pinned-lists-list_id
func (r *ListsResource) RemovePinnedList(id, ListID string) (*ListPinnedResp, *APIError) {
	path := "/users/" + id + "/pinned_lists/" + ListID

	resp := new(ListPinnedResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetUserPinnedLists Returns the Lists pinned by a specified user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/pinned-lists/api-reference/get-users-id-pinned_lists
func (r *ListsResource) GetUserPinnedLists(id string, args ListOpts) (*ListsResp, *APIError) {
	path := "/users/" + id + "/pinned_lists"

	resp := new(ListsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
