package twitter

import "github.com/jarcoal/httpmock"

func (uc *UCSuite) TestLookup() {
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/dm_events",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.DirectMessage.Lookup(DMEventOpts{})
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/dm_events",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1585321444547837956","text":"Another photo https://t.co/J5KotyeIyd","event_type":"MessageCreate","dm_conversation_id":"1585094756761149440","created_at":"2022-10-26T17:24:21.000Z","sender_id":"906948460078698496"}]}`,
		),
	)

	resp, _ := uc.Tw.DirectMessage.Lookup(DMEventOpts{})
	uc.Equal(*resp.Data[0].ID, "1585321444547837956")
}

func (uc *UCSuite) TestLookupByParticipant() {
	paritcipantID := "1585321444547837956"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/dm_conversations/with/"+paritcipantID+"/dm_events",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.DirectMessage.LookupByParticipant(paritcipantID, DMEventOpts{})
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/dm_conversations/with/"+paritcipantID+"/dm_events",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1585321444547837956","text":"Another photo https://t.co/J5KotyeIyd","event_type":"MessageCreate","dm_conversation_id":"1585094756761149440","created_at":"2022-10-26T17:24:21.000Z","sender_id":"906948460078698496"}]}`,
		),
	)

	resp, _ := uc.Tw.DirectMessage.LookupByParticipant(paritcipantID, DMEventOpts{})
	uc.Equal(*resp.Data[0].ID, "1585321444547837956")
}

func (uc *UCSuite) TestLookupByConversation() {
	conversationID := "1585094756761149440"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/dm_conversations/"+conversationID+"/dm_events",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.DirectMessage.LookupByConversation(conversationID, DMEventOpts{})
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/dm_conversations/"+conversationID+"/dm_events",
		httpmock.NewStringResponder(
			200,
			`{"data":[{"id":"1585321444547837956","text":"Another photo https://t.co/J5KotyeIyd","event_type":"MessageCreate","dm_conversation_id":"1585094756761149440","created_at":"2022-10-26T17:24:21.000Z","sender_id":"906948460078698496"}]}`,
		),
	)

	resp, _ := uc.Tw.DirectMessage.LookupByConversation(conversationID, DMEventOpts{})
	uc.Equal(*resp.Data[0].ID, "1585321444547837956")
}

func (uc *UCSuite) TestCreateGroupConversation() {
	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/dm_conversations",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.DirectMessage.CreateGroupConversation(CreateGroupConversationOpts{
		ConversationType: "Group",
		ParticipantIDs:   []string{"944480690", "906948460078698496"},
		Message:          CreateGroupMessageOpts{Text: "Hello to you two, this is a new group conversation"},
	})
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/dm_conversations",
		httpmock.NewStringResponder(
			200,
			`{"data":{"dm_conversation_id":"1346889436626259968","dm_event_id":"128341038123"}}`,
		),
	)

	resp, _ := uc.Tw.DirectMessage.CreateGroupConversation(CreateGroupConversationOpts{
		ConversationType: "Group",
		ParticipantIDs:   []string{"944480690", "906948460078698496"},
		Message:          CreateGroupMessageOpts{Text: "Hello to you two, this is a new group conversation"},
	})
	uc.Equal(*resp.Data.DMConversationID, "1346889436626259968")
}

func (uc *UCSuite) TestCreateMessageToParticipant() {
	participantID := "123456789"
	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/dm_conversations/with/"+participantID+"/messages",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.DirectMessage.CreateMessageToParticipant(
		participantID,
		CreateGroupMessageOpts{Text: "This is a one-to-one Direct Message"},
	)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/dm_conversations/with/"+participantID+"/messages",
		httpmock.NewStringResponder(
			200,
			`{"data":{"dm_conversation_id":"1346889436626259968","dm_event_id":"128341038123"}}`,
		),
	)

	resp, _ := uc.Tw.DirectMessage.CreateMessageToParticipant(
		participantID,
		CreateGroupMessageOpts{Text: "This is a one-to-one Direct Message"},
	)
	uc.Equal(*resp.Data.DMConversationID, "1346889436626259968")
}

func (uc *UCSuite) TestCreateMessageToConversation() {
	conversationID := "1346889436626259968"
	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/dm_conversations/"+conversationID+"/messages",
		httpmock.NewStringResponder(
			401,
			`{"title":"Unauthorized","type":"about:blank","status":401,"detail":"Unauthorized"}`,
		),
	)
	_, err := uc.Tw.DirectMessage.CreateMessageToConversation(
		conversationID,
		CreateGroupMessageOpts{Text: "Adding a Direct Message to a conversation by referencing the conversation ID."},
	)
	uc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/dm_conversations/"+conversationID+"/messages",
		httpmock.NewStringResponder(
			200,
			`{"data":{"dm_conversation_id":"1346889436626259968","dm_event_id":"128341038123"}}`,
		),
	)

	resp, _ := uc.Tw.DirectMessage.CreateMessageToConversation(
		conversationID,
		CreateGroupMessageOpts{Text: "Adding a Direct Message to a conversation by referencing the conversation ID."},
	)
	uc.Equal(*resp.Data.DMConversationID, "1346889436626259968")
}
