package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationsettings
type Evaluationsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RevisionsEnabled - Whether revisions are allowed for evaluations. When enabled, rescoring creates a new version of the evaluation and retracts the existing evaluation version. Does not apply for calibration evaluations.
	RevisionsEnabled *bool `json:"revisionsEnabled,omitempty"`

	// DisputesEnabled - Whether disputes are allowed for evaluations. Does not apply for calibration evaluations.
	DisputesEnabled *bool `json:"disputesEnabled,omitempty"`

	// DisputesAllowedPerEvaluation - The maximum number of disputes allowed for an evaluation.
	DisputesAllowedPerEvaluation *int `json:"disputesAllowedPerEvaluation,omitempty"`

	// DisputesAssignees - A list of assignees responsible for handling each dispute. This list size needs to be equal to disputesAllowedPerEvaluation.
	DisputesAssignees *[]Evaluationsettingsassignee `json:"disputesAssignees,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationsettings) SetField(field string, fieldValue interface{}) {
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

func (o Evaluationsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Evaluationsettings
	
	return json.Marshal(&struct { 
		RevisionsEnabled *bool `json:"revisionsEnabled,omitempty"`
		
		DisputesEnabled *bool `json:"disputesEnabled,omitempty"`
		
		DisputesAllowedPerEvaluation *int `json:"disputesAllowedPerEvaluation,omitempty"`
		
		DisputesAssignees *[]Evaluationsettingsassignee `json:"disputesAssignees,omitempty"`
		Alias
	}{ 
		RevisionsEnabled: o.RevisionsEnabled,
		
		DisputesEnabled: o.DisputesEnabled,
		
		DisputesAllowedPerEvaluation: o.DisputesAllowedPerEvaluation,
		
		DisputesAssignees: o.DisputesAssignees,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluationsettings) UnmarshalJSON(b []byte) error {
	var EvaluationsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationsettingsMap)
	if err != nil {
		return err
	}
	
	if RevisionsEnabled, ok := EvaluationsettingsMap["revisionsEnabled"].(bool); ok {
		o.RevisionsEnabled = &RevisionsEnabled
	}
    
	if DisputesEnabled, ok := EvaluationsettingsMap["disputesEnabled"].(bool); ok {
		o.DisputesEnabled = &DisputesEnabled
	}
    
	if DisputesAllowedPerEvaluation, ok := EvaluationsettingsMap["disputesAllowedPerEvaluation"].(float64); ok {
		DisputesAllowedPerEvaluationInt := int(DisputesAllowedPerEvaluation)
		o.DisputesAllowedPerEvaluation = &DisputesAllowedPerEvaluationInt
	}
	
	if DisputesAssignees, ok := EvaluationsettingsMap["disputesAssignees"].([]interface{}); ok {
		DisputesAssigneesString, _ := json.Marshal(DisputesAssignees)
		json.Unmarshal(DisputesAssigneesString, &o.DisputesAssignees)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
