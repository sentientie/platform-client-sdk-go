package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicemail
type Queueconversationvideoeventtopicemail struct { 
	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`


	// Held - True if this call is held and the person on this side hears silence.
	Held *bool `json:"held,omitempty"`


	// AutoGenerated - Indicates that the email was auto-generated like an Out of Office reply.
	AutoGenerated *bool `json:"autoGenerated,omitempty"`


	// Subject - The subject for the initial email that started this conversation.
	Subject *string `json:"subject,omitempty"`


	// Provider - The source provider of the email.
	Provider *string `json:"provider,omitempty"`


	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`


	// MessagesSent - The number of email messages sent by this participant.
	MessagesSent *int `json:"messagesSent,omitempty"`


	// ErrorInfo - Detailed information about an error response.
	ErrorInfo *Queueconversationvideoeventtopicerrordetails `json:"errorInfo,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime - The timestamp the email was placed on hold in the cloud clock if the email is currently on hold.
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// ConnectedTime - The timestamp when this communication was connected in the cloud clock.
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock.
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// MessageId - A globally unique identifier for the stored content of this communication.
	MessageId *string `json:"messageId,omitempty"`


	// Direction - Whether a call is inbound or outbound.
	Direction *string `json:"direction,omitempty"`


	// DraftAttachments - A list of uploaded attachments on the email draft.
	DraftAttachments *[]Queueconversationvideoeventtopicattachment `json:"draftAttachments,omitempty"`


	// Spam - Indicates if the inbound email was marked as spam.
	Spam *bool `json:"spam,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork - A communication's after-call work data.
	AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired - Indicates if after-call is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

}

func (o *Queueconversationvideoeventtopicemail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicemail
	
	StartHoldTime := new(string)
	if o.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(o.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	DisconnectedTime := new(string)
	if o.DisconnectedTime != nil {
		
		*DisconnectedTime = timeutil.Strftime(o.DisconnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DisconnectedTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		AutoGenerated *bool `json:"autoGenerated,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		MessagesSent *int `json:"messagesSent,omitempty"`
		
		ErrorInfo *Queueconversationvideoeventtopicerrordetails `json:"errorInfo,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		MessageId *string `json:"messageId,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DraftAttachments *[]Queueconversationvideoeventtopicattachment `json:"draftAttachments,omitempty"`
		
		Spam *bool `json:"spam,omitempty"`
		
		Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		Held: o.Held,
		
		AutoGenerated: o.AutoGenerated,
		
		Subject: o.Subject,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		MessagesSent: o.MessagesSent,
		
		ErrorInfo: o.ErrorInfo,
		
		DisconnectType: o.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		MessageId: o.MessageId,
		
		Direction: o.Direction,
		
		DraftAttachments: o.DraftAttachments,
		
		Spam: o.Spam,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopicemail) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicemailMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicemailMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationvideoeventtopicemailMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if State, ok := QueueconversationvideoeventtopicemailMap["state"].(string); ok {
		o.State = &State
	}
	
	if Held, ok := QueueconversationvideoeventtopicemailMap["held"].(bool); ok {
		o.Held = &Held
	}
	
	if AutoGenerated, ok := QueueconversationvideoeventtopicemailMap["autoGenerated"].(bool); ok {
		o.AutoGenerated = &AutoGenerated
	}
	
	if Subject, ok := QueueconversationvideoeventtopicemailMap["subject"].(string); ok {
		o.Subject = &Subject
	}
	
	if Provider, ok := QueueconversationvideoeventtopicemailMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ScriptId, ok := QueueconversationvideoeventtopicemailMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
	
	if PeerId, ok := QueueconversationvideoeventtopicemailMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
	
	if MessagesSent, ok := QueueconversationvideoeventtopicemailMap["messagesSent"].(float64); ok {
		MessagesSentInt := int(MessagesSent)
		o.MessagesSent = &MessagesSentInt
	}
	
	if ErrorInfo, ok := QueueconversationvideoeventtopicemailMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if DisconnectType, ok := QueueconversationvideoeventtopicemailMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if startHoldTimeString, ok := QueueconversationvideoeventtopicemailMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if connectedTimeString, ok := QueueconversationvideoeventtopicemailMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationvideoeventtopicemailMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if MessageId, ok := QueueconversationvideoeventtopicemailMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
	
	if Direction, ok := QueueconversationvideoeventtopicemailMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if DraftAttachments, ok := QueueconversationvideoeventtopicemailMap["draftAttachments"].([]interface{}); ok {
		DraftAttachmentsString, _ := json.Marshal(DraftAttachments)
		json.Unmarshal(DraftAttachmentsString, &o.DraftAttachments)
	}
	
	if Spam, ok := QueueconversationvideoeventtopicemailMap["spam"].(bool); ok {
		o.Spam = &Spam
	}
	
	if Wrapup, ok := QueueconversationvideoeventtopicemailMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationvideoeventtopicemailMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationvideoeventtopicemailMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicemail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
