package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationmessageeventtopicdivisionentityref
type Queueconversationmessageeventtopicdivisionentityref struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

	// DateDivisionUpdated - The time the entity division was last updated.
	DateDivisionUpdated *time.Time `json:"dateDivisionUpdated,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueconversationmessageeventtopicdivisionentityref) SetField(field string, fieldValue interface{}) {
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

func (o Queueconversationmessageeventtopicdivisionentityref) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateDivisionUpdated", }
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
	type Alias Queueconversationmessageeventtopicdivisionentityref
	
	DateDivisionUpdated := new(string)
	if o.DateDivisionUpdated != nil {
		
		*DateDivisionUpdated = timeutil.Strftime(o.DateDivisionUpdated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDivisionUpdated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		DateDivisionUpdated *string `json:"dateDivisionUpdated,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		DateDivisionUpdated: DateDivisionUpdated,
		Alias:    (Alias)(o),
	})
}

func (o *Queueconversationmessageeventtopicdivisionentityref) UnmarshalJSON(b []byte) error {
	var QueueconversationmessageeventtopicdivisionentityrefMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationmessageeventtopicdivisionentityrefMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationmessageeventtopicdivisionentityrefMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := QueueconversationmessageeventtopicdivisionentityrefMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if dateDivisionUpdatedString, ok := QueueconversationmessageeventtopicdivisionentityrefMap["dateDivisionUpdated"].(string); ok {
		DateDivisionUpdated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDivisionUpdatedString)
		o.DateDivisionUpdated = &DateDivisionUpdated
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicdivisionentityref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
