package twitter

/*
	Direct Messages include api for direct message.
*/

type DirectMessageResource Resource

func newDirectMessageResource(cli *Client) *DirectMessageResource {
	return &DirectMessageResource{Cli: cli}
}

// DMEvent represents a dm event for twitter.
// Refer: https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/dm-events
type DMEvent struct {
	ID               *string             `json:"id,omitempty"`
	EventType        *string             `json:"event_type,omitempty"`
	Text             *string             `json:"text,omitempty"`
	SenderID         *string             `json:"sender_id,omitempty"`
	ParticipantIDs   []*string           `json:"participant_ids,omitempty"`
	DMConversationID *string             `json:"dm_conversation_id,omitempty"`
	CreatedAt        *string             `json:"created_at,omitempty"`
	ReferencedTweets *DMEReferencedTweet `json:"referenced_tweets,omitempty"`
	Attachments      *DMEAttachments     `json:"attachments,omitempty"`
}

func (d DMEvent) String() string {
	return Stringify(d)
}

type DMEReferencedTweet struct {
	ID *string `json:"id"`
}

func (d DMEReferencedTweet) String() string {
	return Stringify(d)
}

type DMEAttachments struct {
	MediaKeys []*string `json:"media_keys,omitempty"`
}

func (d DMEAttachments) String() string {
	return Stringify(d)
}

// DMEventOpts specifies the parameters for get DM events
type DMEventOpts struct {
	DMEventFields   string `url:"dm_event.fields,omitempty"`
	EventTypes      string `url:"event_types,omitempty"`
	MaxResults      int    `url:"max_results,omitempty"`
	PaginationToken string `url:"pagination_token,omitempty"`
	Expansions      string `url:"expansions,omitempty"`
	UserFields      string `url:"user.fields,omitempty"`
	MediaFields     string `url:"media.fields,omitempty"`
	TweetFields     string `url:"tweet.fields,omitempty"`
}

// LookUp Returns a list of Direct Messages for the authenticated user, both sent and received.
// Refer: https://developer.twitter.com/en/docs/twitter-api/direct-messages/lookup/api-reference/get-dm_events
func (r *DirectMessageResource) LookUp(args DMEventOpts) (*DMEventsResp, *APIError) {
	path := "/dm_events"

	resp := new(DMEventsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LookUpByParticipant Returns a list of Direct Messages (DM) events within a 1-1 conversation with the user specified in the participant_id path parameter.
// Refer: https://developer.twitter.com/en/docs/twitter-api/direct-messages/lookup/api-reference/get-dm_conversations-with-participant_id-dm_events
func (r *DirectMessageResource) LookUpByParticipant(id string, args DMEventOpts) (*DMEventsResp, *APIError) {
	path := "/dm_conversations/with/" + id + "/dm_events"

	resp := new(DMEventsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LookUpByConversation Returns a list of Direct Messages within a conversation specified in the dm_conversation_id path parameter.
// Refer: https://developer.twitter.com/en/docs/twitter-api/direct-messages/lookup/api-reference/get-dm_conversations-dm_conversation_id-dm_events
func (r *DirectMessageResource) LookUpByConversation(id string, args DMEventOpts) (*DMEventsResp, *APIError) {
	path := "/dm_conversations/" + id + "/dm_events"

	resp := new(DMEventsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
