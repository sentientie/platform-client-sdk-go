package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Maskingrule
type Maskingrule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - Masking rule name.
	Name *string `json:"name,omitempty"`

	// Description - Description about masking rule.
	Description *string `json:"description,omitempty"`

	// SubstituteCharacter - Replacement character for masked text character.
	SubstituteCharacter *string `json:"substituteCharacter,omitempty"`

	// Definition - Definition of masking rule (a valid regex or builtin AI based mask name).
	Definition *string `json:"definition,omitempty"`

	// Enabled - True/False
	Enabled *bool `json:"enabled,omitempty"`

	// VarType - Masking rule type
	VarType *string `json:"type,omitempty"`

	// Integrations - Associated integration channels
	Integrations *[]string `json:"integrations,omitempty"`

	// DateCreated - Date when the rule was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Date when the rule was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Maskingrule) SetField(field string, fieldValue interface{}) {
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

func (o Maskingrule) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Maskingrule
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SubstituteCharacter *string `json:"substituteCharacter,omitempty"`
		
		Definition *string `json:"definition,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Integrations *[]string `json:"integrations,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		SubstituteCharacter: o.SubstituteCharacter,
		
		Definition: o.Definition,
		
		Enabled: o.Enabled,
		
		VarType: o.VarType,
		
		Integrations: o.Integrations,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		Alias:    (Alias)(o),
	})
}

func (o *Maskingrule) UnmarshalJSON(b []byte) error {
	var MaskingruleMap map[string]interface{}
	err := json.Unmarshal(b, &MaskingruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MaskingruleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MaskingruleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := MaskingruleMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if SubstituteCharacter, ok := MaskingruleMap["substituteCharacter"].(string); ok {
		o.SubstituteCharacter = &SubstituteCharacter
	}
    
	if Definition, ok := MaskingruleMap["definition"].(string); ok {
		o.Definition = &Definition
	}
    
	if Enabled, ok := MaskingruleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if VarType, ok := MaskingruleMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Integrations, ok := MaskingruleMap["integrations"].([]interface{}); ok {
		IntegrationsString, _ := json.Marshal(Integrations)
		json.Unmarshal(IntegrationsString, &o.Integrations)
	}
	
	if dateCreatedString, ok := MaskingruleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := MaskingruleMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Maskingrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
