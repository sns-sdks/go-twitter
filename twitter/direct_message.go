package twitter

/*
	Direct Messages include api for direct message.
*/

type DirectMessageResource Resource

func newDirectMessageResource(cli *Client) *DirectMessageResource {
	return &DirectMessageResource{Cli: cli}
}

// DMEvent represents a dm event for Twitter.
// Refer: https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/dm-events
type DMEvent struct {
	ID               *string               `json:"id,omitempty"`
	EventType        *string               `json:"event_type,omitempty"`
	Text             *string               `json:"text,omitempty"`
	SenderID         *string               `json:"sender_id,omitempty"`
	ParticipantIDs   []*string             `json:"participant_ids,omitempty"`
	DMConversationID *string               `json:"dm_conversation_id,omitempty"`
	CreatedAt        *string               `json:"created_at,omitempty"`
	ReferencedTweets []*DMEReferencedTweet `json:"referenced_tweets,omitempty"`
	Attachments      *DMEAttachments       `json:"attachments,omitempty"`
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
func (r *DirectMessageResource) Lookup(args DMEventOpts) (*DMEventsResp, *APIError) {
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
func (r *DirectMessageResource) LookupByParticipant(id string, args DMEventOpts) (*DMEventsResp, *APIError) {
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
func (r *DirectMessageResource) LookupByConversation(id string, args DMEventOpts) (*DMEventsResp, *APIError) {
	path := "/dm_conversations/" + id + "/dm_events"

	resp := new(DMEventsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DMConversation represents a dm conversation for Twitter.
type DMConversation struct {
	DMConversationID *string `json:"dm_conversation_id,omitempty"`
	DMEventID        *string `json:"dm_event_id,omitempty"`
}

func (d DMConversation) String() string {
	return Stringify(d)
}

// CreateGroupMessageAttachmentOpts specifies the parameters for message attachment
type CreateGroupMessageAttachmentOpts struct {
	MediaID string `url:"media_id,omitempty"`
}

// CreateGroupMessageOpts specifies the parameters for message attachment
type CreateGroupMessageOpts struct {
	Attachments []CreateGroupMessageAttachmentOpts `url:"attachments,omitempty"`
	Text        string                             `url:"text,omitempty"`
}

// CreateGroupConversationOpts specifies the parameters for create conversation.
type CreateGroupConversationOpts struct {
	ConversationType string                 `url:"conversation_type"`
	ParticipantIDs   []string               `url:"participant_ids"`
	Message          CreateGroupMessageOpts `url:"message"`
}

// CreateGroupConversation Creates a new group conversation and adds a Direct Message to it on behalf of an authenticated user.
// Refer: https://developer.twitter.com/en/docs/twitter-api/direct-messages/manage/api-reference/post-dm_conversations
func (r *DirectMessageResource) CreateGroupConversation(args CreateGroupConversationOpts) (*DMConversationResp, *APIError) {
	path := "/dm_conversations"

	resp := new(DMConversationResp)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateMessageToParticipant Creates a one-to-one Direct Message and adds it to the one-to-one conversation.
// Refer: https://developer.twitter.com/en/docs/twitter-api/direct-messages/manage/api-reference/post-dm_conversations-with-participant_id-messages
func (r *DirectMessageResource) CreateMessageToParticipant(participantID string, args CreateGroupMessageOpts) (*DMConversationResp, *APIError) {
	path := "/dm_conversations/with/" + participantID + "/messages"

	resp := new(DMConversationResp)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateMessageToConversation Creates a Direct Message on behalf of an authenticated user, and adds it to the specified conversation.
// Refer: https://developer.twitter.com/en/docs/twitter-api/direct-messages/manage/api-reference/post-dm_conversations-dm_conversation_id-messages
func (r *DirectMessageResource) CreateMessageToConversation(dmConversationID string, args CreateGroupMessageOpts) (*DMConversationResp, *APIError) {
	path := "/dm_conversations/" + dmConversationID + "/messages"

	resp := new(DMConversationResp)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
