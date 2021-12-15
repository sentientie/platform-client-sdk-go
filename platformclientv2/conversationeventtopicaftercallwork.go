package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicaftercallwork - A communication's after-call work data.
type Conversationeventtopicaftercallwork struct { 
	// State - The communication's after-call work state.
	State *string `json:"state,omitempty"`


	// StartTime - The timestamp when this communication started after-call work in the cloud clock.
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The timestamp when this communication ended after-call work in the cloud clock.
	EndTime *time.Time `json:"endTime,omitempty"`

}

func (o *Conversationeventtopicaftercallwork) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicaftercallwork
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		StartTime: StartTime,
		
		EndTime: EndTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationeventtopicaftercallwork) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicaftercallworkMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicaftercallworkMap)
	if err != nil {
		return err
	}
	
	if State, ok := ConversationeventtopicaftercallworkMap["state"].(string); ok {
		o.State = &State
	}
	
	if startTimeString, ok := ConversationeventtopicaftercallworkMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if endTimeString, ok := ConversationeventtopicaftercallworkMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicaftercallwork) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
