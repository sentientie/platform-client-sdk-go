package platformclientv2
import (
	"time"
	"encoding/json"
)

// Conversationemaileventtopicemailmediaparticipant
type Conversationemaileventtopicemailmediaparticipant struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Address
	Address *string `json:"address,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// EndTime
	EndTime *time.Time `json:"endTime,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// Purpose
	Purpose *string `json:"purpose,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// WrapupRequired
	WrapupRequired *bool `json:"wrapupRequired,omitempty"`


	// WrapupPrompt
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`


	// User
	User *Conversationemaileventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Conversationemaileventtopicurireference `json:"queue,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Conversationemaileventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Conversationemaileventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int32 `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int32 `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Conversationemaileventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Conversationemaileventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Conversationemaileventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Conversationemaileventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Conversationemaileventtopicjourneycontext `json:"journeyContext,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// Subject
	Subject *string `json:"subject,omitempty"`


	// MessagesSent
	MessagesSent *int32 `json:"messagesSent,omitempty"`


	// AutoGenerated
	AutoGenerated *bool `json:"autoGenerated,omitempty"`


	// MessageId
	MessageId *string `json:"messageId,omitempty"`


	// DraftAttachments
	DraftAttachments *[]Conversationemaileventtopicattachment `json:"draftAttachments,omitempty"`


	// Spam
	Spam *bool `json:"spam,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicemailmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
