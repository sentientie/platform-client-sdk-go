package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queueconversationmessageeventtopicmessagedetails
type Queueconversationmessageeventtopicmessagedetails struct { 
	// Message
	Message *Queueconversationmessageeventtopicurireference `json:"message,omitempty"`


	// MessageTime
	MessageTime *time.Time `json:"messageTime,omitempty"`


	// MessageSegmentCount
	MessageSegmentCount *int32 `json:"messageSegmentCount,omitempty"`


	// MessageStatus
	MessageStatus *string `json:"messageStatus,omitempty"`


	// Media
	Media *[]Queueconversationmessageeventtopicmessagemedia `json:"media,omitempty"`


	// Stickers
	Stickers *[]Queueconversationmessageeventtopicmessagesticker `json:"stickers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicmessagedetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
