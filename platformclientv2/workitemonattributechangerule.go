package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemonattributechangerule
type Workitemonattributechangerule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// VarType - The type of the rule.
	VarType *string `json:"type,omitempty"`

	// Action - The rules action. If the condition criteria is met this action will be executed.
	Action *Workitemruleaction `json:"action,omitempty"`

	// Worktype - The Worktype containing the rule.
	Worktype *Worktypereference `json:"worktype,omitempty"`

	// Condition - The rules condition. If the condition criteria is met the rules action will be executed.
	Condition *Workitemonattributechangecondition `json:"condition,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workitemonattributechangerule) SetField(field string, fieldValue interface{}) {
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

func (o Workitemonattributechangerule) MarshalJSON() ([]byte, error) {
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
	type Alias Workitemonattributechangerule
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Action *Workitemruleaction `json:"action,omitempty"`
		
		Worktype *Worktypereference `json:"worktype,omitempty"`
		
		Condition *Workitemonattributechangecondition `json:"condition,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		Action: o.Action,
		
		Worktype: o.Worktype,
		
		Condition: o.Condition,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Workitemonattributechangerule) UnmarshalJSON(b []byte) error {
	var WorkitemonattributechangeruleMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemonattributechangeruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkitemonattributechangeruleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorkitemonattributechangeruleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := WorkitemonattributechangeruleMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Action, ok := WorkitemonattributechangeruleMap["action"].(map[string]interface{}); ok {
		ActionString, _ := json.Marshal(Action)
		json.Unmarshal(ActionString, &o.Action)
	}
	
	if Worktype, ok := WorkitemonattributechangeruleMap["worktype"].(map[string]interface{}); ok {
		WorktypeString, _ := json.Marshal(Worktype)
		json.Unmarshal(WorktypeString, &o.Worktype)
	}
	
	if Condition, ok := WorkitemonattributechangeruleMap["condition"].(map[string]interface{}); ok {
		ConditionString, _ := json.Marshal(Condition)
		json.Unmarshal(ConditionString, &o.Condition)
	}
	
	if SelfUri, ok := WorkitemonattributechangeruleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemonattributechangerule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
