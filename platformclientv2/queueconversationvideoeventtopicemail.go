package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queueconversationvideoeventtopicemail
type Queueconversationvideoeventtopicemail struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// AutoGenerated
	AutoGenerated *bool `json:"autoGenerated,omitempty"`


	// Subject
	Subject *string `json:"subject,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId
	PeerId *string `json:"peerId,omitempty"`


	// MessagesSent
	MessagesSent *int `json:"messagesSent,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationvideoeventtopicerrordetails `json:"errorInfo,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// MessageId
	MessageId *string `json:"messageId,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// DraftAttachments
	DraftAttachments *[]Queueconversationvideoeventtopicattachment `json:"draftAttachments,omitempty"`


	// Spam
	Spam *bool `json:"spam,omitempty"`


	// Wrapup
	Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicemail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
