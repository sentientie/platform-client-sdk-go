package platformclientv2
import (
	"time"
	"encoding/json"
)

// Callbackmediaparticipant
type Callbackmediaparticipant struct { 
	// Id - The unique participant ID.
	Id *string `json:"id,omitempty"`


	// Name - The display friendly name of the participant.
	Name *string `json:"name,omitempty"`


	// Address - The participant address.
	Address *string `json:"address,omitempty"`


	// StartTime - The time when this participant first joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// ConnectedTime - The time when this participant went connected for this media (eg: video connected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// EndTime - The time when this participant went disconnected for this media (eg: video disconnected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// StartHoldTime - The time when this participant's hold started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// Purpose - The participant's purpose.  Values can be: 'agent', 'user', 'customer', 'external', 'acd', 'ivr
	Purpose *string `json:"purpose,omitempty"`


	// State - The participant's state.  Values can be: 'alerting', 'connected', 'disconnected', 'dialing', 'contacting
	State *string `json:"state,omitempty"`


	// Direction - The participant's direction.  Values can be: 'inbound' or 'outbound'
	Direction *string `json:"direction,omitempty"`


	// DisconnectType - The reason the participant was disconnected from the conversation.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// Held - Value is true when the participant is on hold.
	Held *bool `json:"held,omitempty"`


	// WrapupRequired - Value is true when the participant requires wrap-up.
	WrapupRequired *bool `json:"wrapupRequired,omitempty"`


	// WrapupPrompt - The wrap-up prompt indicating the type of wrap-up to be performed.
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`


	// User - The PureCloud user for this participant.
	User *Domainentityref `json:"user,omitempty"`


	// Queue - The PureCloud queue for this participant.
	Queue *Domainentityref `json:"queue,omitempty"`


	// Team - The PureCloud team for this participant.
	Team *Domainentityref `json:"team,omitempty"`


	// Attributes - A list of ad-hoc attributes for the participant.
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo - If the conversation ends in error, contains additional error details.
	ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`


	// Script - The Engage script that should be used by this participant.
	Script *Domainentityref `json:"script,omitempty"`


	// WrapupTimeoutMs - The amount of time the participant has to complete wrap-up.
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped - Value is true when the participant has skipped wrap-up.
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs - Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider - The source provider for the communication.
	Provider *string `json:"provider,omitempty"`


	// ExternalContact - If this participant represents an external contact, then this will be the reference for the external contact.
	ExternalContact *Domainentityref `json:"externalContact,omitempty"`


	// ExternalOrganization - If this participant represents an external org, then this will be the reference for the external org.
	ExternalOrganization *Domainentityref `json:"externalOrganization,omitempty"`


	// Wrapup - Wrapup for this participant, if it has been applied.
	Wrapup *Wrapup `json:"wrapup,omitempty"`


	// Peer - The peer communication corresponding to a matching leg for this communication.
	Peer *string `json:"peer,omitempty"`


	// FlaggedReason - The reason specifying why participant flagged the conversation.
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext - Journey System data/context that is applicable to this communication.  When used for historical purposes, the context should be immutable.  When null, there is no applicable Journey System context.
	JourneyContext *Journeycontext `json:"journeyContext,omitempty"`


	// ConversationRoutingData - Information on how a communication should be routed to an agent.
	ConversationRoutingData *Conversationroutingdata `json:"conversationRoutingData,omitempty"`


	// StartAcwTime - The timestamp when this participant started after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime - The timestamp when this participant ended after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// OutboundPreview - The outbound preview associated with this callback.
	OutboundPreview *Dialerpreview `json:"outboundPreview,omitempty"`


	// Voicemail - The voicemail associated with this callback.
	Voicemail *Voicemail `json:"voicemail,omitempty"`


	// CallbackNumbers - The list of phone number to use for this callback.
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`


	// CallbackUserName - The name of the callback target.
	CallbackUserName *string `json:"callbackUserName,omitempty"`


	// ExternalCampaign - True if the call for the callback uses external dialing.
	ExternalCampaign *bool `json:"externalCampaign,omitempty"`


	// SkipEnabled - If true, the callback can be skipped.
	SkipEnabled *bool `json:"skipEnabled,omitempty"`


	// TimeoutSeconds - Duration in seconds before the callback will be auto-dialed.
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`


	// AutomatedCallbackConfigId - The id of the config for automatically placing the callback (and handling the disposition). If absent, the callback will not be placed automatically but routed to an agent as per normal.
	AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`


	// CallbackScheduledTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callbackmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
