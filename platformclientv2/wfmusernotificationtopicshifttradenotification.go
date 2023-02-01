package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmusernotificationtopicshifttradenotification
type Wfmusernotificationtopicshifttradenotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// WeekDate
	WeekDate *string `json:"weekDate,omitempty"`

	// TradeId
	TradeId *string `json:"tradeId,omitempty"`

	// OneSided
	OneSided *bool `json:"oneSided,omitempty"`

	// NewState
	NewState *string `json:"newState,omitempty"`

	// InitiatingUser
	InitiatingUser *Wfmusernotificationtopicuserreference `json:"initiatingUser,omitempty"`

	// InitiatingShiftDate
	InitiatingShiftDate *time.Time `json:"initiatingShiftDate,omitempty"`

	// ReceivingUser
	ReceivingUser *Wfmusernotificationtopicuserreference `json:"receivingUser,omitempty"`

	// ReceivingShiftDate
	ReceivingShiftDate *time.Time `json:"receivingShiftDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmusernotificationtopicshifttradenotification) SetField(field string, fieldValue interface{}) {
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

func (o Wfmusernotificationtopicshifttradenotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "InitiatingShiftDate","ReceivingShiftDate", }
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
	type Alias Wfmusernotificationtopicshifttradenotification
	
	InitiatingShiftDate := new(string)
	if o.InitiatingShiftDate != nil {
		
		*InitiatingShiftDate = timeutil.Strftime(o.InitiatingShiftDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		InitiatingShiftDate = nil
	}
	
	ReceivingShiftDate := new(string)
	if o.ReceivingShiftDate != nil {
		
		*ReceivingShiftDate = timeutil.Strftime(o.ReceivingShiftDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReceivingShiftDate = nil
	}
	
	return json.Marshal(&struct { 
		WeekDate *string `json:"weekDate,omitempty"`
		
		TradeId *string `json:"tradeId,omitempty"`
		
		OneSided *bool `json:"oneSided,omitempty"`
		
		NewState *string `json:"newState,omitempty"`
		
		InitiatingUser *Wfmusernotificationtopicuserreference `json:"initiatingUser,omitempty"`
		
		InitiatingShiftDate *string `json:"initiatingShiftDate,omitempty"`
		
		ReceivingUser *Wfmusernotificationtopicuserreference `json:"receivingUser,omitempty"`
		
		ReceivingShiftDate *string `json:"receivingShiftDate,omitempty"`
		Alias
	}{ 
		WeekDate: o.WeekDate,
		
		TradeId: o.TradeId,
		
		OneSided: o.OneSided,
		
		NewState: o.NewState,
		
		InitiatingUser: o.InitiatingUser,
		
		InitiatingShiftDate: InitiatingShiftDate,
		
		ReceivingUser: o.ReceivingUser,
		
		ReceivingShiftDate: ReceivingShiftDate,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmusernotificationtopicshifttradenotification) UnmarshalJSON(b []byte) error {
	var WfmusernotificationtopicshifttradenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmusernotificationtopicshifttradenotificationMap)
	if err != nil {
		return err
	}
	
	if WeekDate, ok := WfmusernotificationtopicshifttradenotificationMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
    
	if TradeId, ok := WfmusernotificationtopicshifttradenotificationMap["tradeId"].(string); ok {
		o.TradeId = &TradeId
	}
    
	if OneSided, ok := WfmusernotificationtopicshifttradenotificationMap["oneSided"].(bool); ok {
		o.OneSided = &OneSided
	}
    
	if NewState, ok := WfmusernotificationtopicshifttradenotificationMap["newState"].(string); ok {
		o.NewState = &NewState
	}
    
	if InitiatingUser, ok := WfmusernotificationtopicshifttradenotificationMap["initiatingUser"].(map[string]interface{}); ok {
		InitiatingUserString, _ := json.Marshal(InitiatingUser)
		json.Unmarshal(InitiatingUserString, &o.InitiatingUser)
	}
	
	if initiatingShiftDateString, ok := WfmusernotificationtopicshifttradenotificationMap["initiatingShiftDate"].(string); ok {
		InitiatingShiftDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", initiatingShiftDateString)
		o.InitiatingShiftDate = &InitiatingShiftDate
	}
	
	if ReceivingUser, ok := WfmusernotificationtopicshifttradenotificationMap["receivingUser"].(map[string]interface{}); ok {
		ReceivingUserString, _ := json.Marshal(ReceivingUser)
		json.Unmarshal(ReceivingUserString, &o.ReceivingUser)
	}
	
	if receivingShiftDateString, ok := WfmusernotificationtopicshifttradenotificationMap["receivingShiftDate"].(string); ok {
		ReceivingShiftDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", receivingShiftDateString)
		o.ReceivingShiftDate = &ReceivingShiftDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmusernotificationtopicshifttradenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
