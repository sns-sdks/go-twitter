package twitter

/*
	Lists members include api for lists members
*/

// ListMemberStatus represents the status for member is belonged to the list
type ListMemberStatus struct {
	IsMember *bool `json:"is_member,omitempty"`
}

func (m ListMemberStatus) String() string {
	return Stringify(m)
}

// ListMemberResp represents the response for member is belonged to the list
type ListMemberResp struct {
	Data *ListMemberStatus `json:"data,omitempty"`
}

func (m ListMemberResp) String() string {
	return Stringify(m)
}

// addMemberOpts Inline parameters for add members
type addMemberOpts struct {
	UserID string `json:"user_id"`
}

// AddListMember Enables the authenticated user to add a member to a List they own.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-lists-id-members
func (r *ListsResource) AddListMember(id, userID string) (*ListMemberResp, *APIError) {
	path := "/lists/" + id + "/members"
	args := addMemberOpts{UserID: userID}

	resp := new(ListMemberResp)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RemoveListMember Enables the authenticated user to remove a member from a List they own.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-lists-id-members-user_id
func (r *ListsResource) RemoveListMember(id, userID string) (*ListMemberResp, *APIError) {
	path := "/lists/" + id + "/members/" + userID

	resp := new(ListMemberResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListMembersOpts specifies the parameters for get list members.
type ListMembersOpts struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	UserOpts
}

// GetListMembers Returns a list of users who are members of the specified List.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/list-members/api-reference/get-lists-id-members
func (r *ListsResource) GetListMembers(id string, args ListMembersOpts) (*UsersResp, *APIError) {
	path := "/lists/" + id + "/members"

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// JoinedListsOpts Specifies the parameters for get joined lists.
type JoinedListsOpts struct {
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	ListOpts
}

// GetUserJoinedLists Returns all Lists a specified user is a member of.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/list-members/api-reference/get-users-id-list_memberships
func (r *ListsResource) GetUserJoinedLists(id string, args JoinedListsOpts) (*ListsResp, *APIError) {
	path := "/users/" + id + "/list_memberships"

	resp := new(ListsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
