package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Openeventnormalizedmessage - Open Messaging rich media message structure
type Openeventnormalizedmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Unique ID of the message generated by Messaging Platform.
	Id *string `json:"id,omitempty"`

	// Channel - Channel-specific information that describes the message and the message channel/provider.
	Channel *Openmessagingchannel `json:"channel,omitempty"`

	// VarType - Message type.
	VarType *string `json:"type,omitempty"`

	// Events - List of event elements.
	Events *[]Openmessageevent `json:"events,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Openeventnormalizedmessage) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Openeventnormalizedmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Openeventnormalizedmessage
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Channel *Openmessagingchannel `json:"channel,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Events *[]Openmessageevent `json:"events,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Channel: o.Channel,
		
		VarType: o.VarType,
		
		Events: o.Events,
		Alias:    (Alias)(o),
	})
}

func (o *Openeventnormalizedmessage) UnmarshalJSON(b []byte) error {
	var OpeneventnormalizedmessageMap map[string]interface{}
	err := json.Unmarshal(b, &OpeneventnormalizedmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OpeneventnormalizedmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Channel, ok := OpeneventnormalizedmessageMap["channel"].(map[string]interface{}); ok {
		ChannelString, _ := json.Marshal(Channel)
		json.Unmarshal(ChannelString, &o.Channel)
	}
	
	if VarType, ok := OpeneventnormalizedmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Events, ok := OpeneventnormalizedmessageMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Openeventnormalizedmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
