package twitter

/*
	Lists include api for lists
*/

type ListsResource Resource

func newListsResource(cli *Client) *ListsResource {
	return &ListsResource{Cli: cli}
}

// List represents a list for twitter.
// Refer: https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/lists
type List struct {
	ID            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	CreatedAt     *string `json:"created_at,omitempty"`
	Description   *string `json:"description,omitempty"`
	FollowerCount *int    `json:"follower_count,omitempty"`
	MemberCount   *int    `json:"member_count,omitempty"`
	Private       *bool   `json:"private,omitempty"`
	OwnerID       *string `json:"owner_id,omitempty"`
}

func (l List) String() string {
	return Stringify(l)
}

// ListResp represents the response for a list
type ListResp struct {
	Data *List `json:"data,omitempty"`
}

func (l ListResp) String() string {
	return Stringify(l)
}

// ListsResp represents the response for multi lists
type ListsResp struct {
	Data *[]List `json:"data,omitempty"`
	*BaseData
}

func (l ListsResp) String() string {
	return Stringify(l)
}

// ListOpts specifies the parameter for get list
type ListOpts struct {
	List       string `url:"list.fields,omitempty"`
	Expansions string `url:"expansions,omitempty"`
	User       string `url:"user.fields,omitempty"`
}

// LookupByID Returns the details of a specified List.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/list-lookup/api-reference/get-lists-id
func (r *ListsResource) LookupByID(id string, args ListOpts) (*ListResp, *APIError) {
	path := "/lists/" + id

	resp := new(ListResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// OwnedListsOpts Specifies the parameter for get owned lists.
type OwnedListsOpts struct {
	MaxResults      string `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	ListOpts
}

// GetOwnedLists Returns all Lists owned by the specified user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/list-lookup/api-reference/get-users-id-owned_lists
func (r *ListsResource) GetOwnedLists(id string, args OwnedListsOpts) (*ListsResp, *APIError) {
	path := "/users/" + id + "/owned_lists"

	resp := new(ListsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
