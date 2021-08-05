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

type UserOpts struct {
	UserFields  string `url:"user.fields,omitempty"`
	Expansions  string `url:"expansions,omitempty"`
	TweetFields string `url:"tweet.fields,omitempty"`
}

func (r *UserResource) LookupByID(id string, args UserOpts) (*UserResp, *APIError) {
	path := Baseurl + "/users/" + id
	resp := new(UserResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type userOptsByIDs struct {
	IDs string `url:"ids"`
	UserOpts
}

func (r *UserResource) LookupByIDs(ids string, args UserOpts) (*UsersResp, *APIError) {
	path := Baseurl + "/users"
	newArgs := userOptsByIDs{ids, args}

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, newArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *UserResource) LookupByUsername(username string, args UserOpts) (*UserResp, *APIError) {
	path := Baseurl + "/users/by/username/" + username

	resp := new(UserResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type userOptsByUsernames struct {
	Usernames string `url:"usernames"`
	UserOpts
}

func (r *UserResource) LookupByUsernames(usernames string, args UserOpts) (*UsersResp, *APIError) {
	path := Baseurl + "/users/by"
	newArgs := userOptsByUsernames{usernames, args}

	resp := new(UsersResp)
	err := r.Cli.DoGet(path, newArgs, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
