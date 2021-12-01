package twitter

// ManageListOpts represents the parameters for manage list create and update
type ManageListOpts struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private,omitempty"`
}

// ListCreate Enables the authenticated user to create a List.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-lists
func (r *ListsResource) ListCreate(args ManageListOpts) (*ListResp, *APIError) {
	path := "/lists"

	resp := new(ListResp)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListDeletedStatus represents the status for update list
type ListDeletedStatus struct {
	Deleted *bool `json:"deleted,omitempty"`
}

func (l ListDeletedStatus) String() string {
	return Stringify(l)
}

// ListDeletedResp represents the response for delete list
type ListDeletedResp struct {
	Data *ListDeletedStatus `json:"data,omitempty"`
}

func (l ListDeletedResp) String() string {
	return Stringify(l)
}

// ListDelete Enables the authenticated user to delete a List that they own.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-lists-id
func (r *ListsResource) ListDelete(id string) (*ListDeletedResp, *APIError) {
	path := "/lists/" + id

	resp := new(ListDeletedResp)
	err := r.Cli.DoDelete(path, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListUpdatedStatus represents the status for update list
type ListUpdatedStatus struct {
	Updated *bool `json:"updated,omitempty"`
}

func (l ListUpdatedStatus) String() string {
	return Stringify(l)
}

// ListUpdatedResp represents the response for update list
type ListUpdatedResp struct {
	Data *ListUpdatedStatus `json:"data,omitempty"`
}

func (l ListUpdatedResp) String() string {
	return Stringify(l)
}

// ListUpdate Enables the authenticated user to update the metadata of a specified List that they own.
// Refer: https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/put-lists-id
func (r *ListsResource) ListUpdate(id string, args ManageListOpts) (*ListUpdatedResp, *APIError) {
	path := "/lists/" + id

	resp := new(ListUpdatedResp)
	err := r.Cli.DoPut(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
