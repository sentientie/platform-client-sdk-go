package platformclientv2
import (
	"encoding/json"
)

// Outboundmessagingmessagingcampaignconfigchangeresponseref
type Outboundmessagingmessagingcampaignconfigchangeresponseref struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangeresponseref) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
