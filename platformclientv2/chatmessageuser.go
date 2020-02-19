package platformclientv2
import (
	"encoding/json"
)

// Chatmessageuser
type Chatmessageuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`


	// Username
	Username *string `json:"username,omitempty"`


	// Images
	Images *[]Userimage `json:"images,omitempty"`

}

// String returns a JSON representation of the model
func (o *Chatmessageuser) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
