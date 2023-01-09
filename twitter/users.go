package twitter

/*
	Users API
*/

type UserResource Resource

func newUserResource(cli *Client) *UserResource {
	return &UserResource{
		Cli: cli,
	}
}

// User represent a Twitter user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/user
type User struct {
	ID              *string            `json:"id,omitempty"`
	Name            *string            `json:"name,omitempty"`
	Username        *string            `json:"username,omitempty"`
	CreatedAt       *string            `json:"created_at,omitempty"`
	Description     *string            `json:"description,omitempty"`
	Location        *string            `json:"location,omitempty"`
	PinnedTweetID   *string            `json:"pinned_tweet_id,omitempty"`
	ProfileImageUrl *string            `json:"profile_image_url,omitempty"`
	Protected       *string            `json:"protected,omitempty"`
	URL             *string            `json:"url,omitempty,omitempty"`
	Verified        *bool              `json:"verified,omitempty"`
	VerifiedType    *string            `json:"verified_type,omitempty"`
	Entities        *UserEntities      `json:"entities,omitempty"`
	PublicMetrics   *UserPublicMetrics `json:"public_metrics,omitempty"`
	Withheld        *UserWithheld      `json:"withheld,omitempty"`
}

type UserPublicMetrics struct {
	FollowerCount  *int `json:"follower_count,omitempty"`
	FollowingCount *int `json:"following_count,omitempty"`
	TweetCount     *int `json:"tweet_count,omitempty"`
	ListCount      *int `json:"list_count,omitempty"`
}

type UserWithheld struct {
	Scope        *string   `json:"scope,omitempty"`
	CountryCodes []*string `json:"country_codes,omitempty"`
}

func (u User) String() string {
	return Stringify(u)
}

// UserOpts specifies the parameters for get user
type UserOpts struct {
	UserFields  string `url:"user.fields,omitempty"`
	Expansions  string `url:"expansions,omitempty"`
	TweetFields string `url:"tweet.fields,omitempty"`
}

// LookupByID Returns a variety of information about a single user specified by the requested ID.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-id
func (r *UserResource) LookupByID(id string, args UserOpts) (*UserResp, *APIError) {
	path := "/users/" + id
	resp := new(UserResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// userOptsByIDs specifies the parameters for get user by ids
type userOptsByIDs struct {
	IDs string `url:"ids"`
	UserOpts
}

// LookupByIDs Returns a variety of information about one or more users specified by the requested IDs.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
func (r *UserResource) LookupByIDs(ids string, args UserOpts) (*UsersResp, *APIError) {
	path := "/users"
	newArgs := userOptsByIDs{ids, args}

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, newArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LookupByUsername Returns a variety of information about a single user specified by his username.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
func (r *UserResource) LookupByUsername(username string, args UserOpts) (*UserResp, *APIError) {
	path := "/users/by/username/" + username

	resp := new(UserResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// userOptsByUsernames specifies the parameters for get user by username
type userOptsByUsernames struct {
	Usernames string `url:"usernames"`
	UserOpts
}

// LookupByUsernames Returns a variety of information about one or more users specified by their usernames.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by
func (r *UserResource) LookupByUsernames(usernames string, args UserOpts) (*UsersResp, *APIError) {
	path := "/users/by"
	newArgs := userOptsByUsernames{usernames, args}

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, newArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LookupMe Returns information about an authorized user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-me
func (r *UserResource) LookupMe(args UserOpts) (*UserResp, *APIError) {
	path := "/users/me"

	resp := new(UserResp)
	err := r.Cli.DoGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
