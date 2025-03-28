package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationdevice
type Journeysessioneventsnotificationdevice struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType
	VarType *string `json:"type,omitempty"`

	// IsMobile
	IsMobile *bool `json:"isMobile,omitempty"`

	// ScreenHeight
	ScreenHeight *int `json:"screenHeight,omitempty"`

	// ScreenWidth
	ScreenWidth *int `json:"screenWidth,omitempty"`

	// ScreenDensity
	ScreenDensity *int `json:"screenDensity,omitempty"`

	// Fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// OsFamily
	OsFamily *string `json:"osFamily,omitempty"`

	// OsVersion
	OsVersion *string `json:"osVersion,omitempty"`

	// Category
	Category *string `json:"category,omitempty"`

	// Manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeysessioneventsnotificationdevice) SetField(field string, fieldValue interface{}) {
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

func (o Journeysessioneventsnotificationdevice) MarshalJSON() ([]byte, error) {
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
	type Alias Journeysessioneventsnotificationdevice
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		IsMobile *bool `json:"isMobile,omitempty"`
		
		ScreenHeight *int `json:"screenHeight,omitempty"`
		
		ScreenWidth *int `json:"screenWidth,omitempty"`
		
		ScreenDensity *int `json:"screenDensity,omitempty"`
		
		Fingerprint *string `json:"fingerprint,omitempty"`
		
		OsFamily *string `json:"osFamily,omitempty"`
		
		OsVersion *string `json:"osVersion,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		Manufacturer *string `json:"manufacturer,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		IsMobile: o.IsMobile,
		
		ScreenHeight: o.ScreenHeight,
		
		ScreenWidth: o.ScreenWidth,
		
		ScreenDensity: o.ScreenDensity,
		
		Fingerprint: o.Fingerprint,
		
		OsFamily: o.OsFamily,
		
		OsVersion: o.OsVersion,
		
		Category: o.Category,
		
		Manufacturer: o.Manufacturer,
		Alias:    (Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationdevice) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationdeviceMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationdeviceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := JourneysessioneventsnotificationdeviceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if IsMobile, ok := JourneysessioneventsnotificationdeviceMap["isMobile"].(bool); ok {
		o.IsMobile = &IsMobile
	}
    
	if ScreenHeight, ok := JourneysessioneventsnotificationdeviceMap["screenHeight"].(float64); ok {
		ScreenHeightInt := int(ScreenHeight)
		o.ScreenHeight = &ScreenHeightInt
	}
	
	if ScreenWidth, ok := JourneysessioneventsnotificationdeviceMap["screenWidth"].(float64); ok {
		ScreenWidthInt := int(ScreenWidth)
		o.ScreenWidth = &ScreenWidthInt
	}
	
	if ScreenDensity, ok := JourneysessioneventsnotificationdeviceMap["screenDensity"].(float64); ok {
		ScreenDensityInt := int(ScreenDensity)
		o.ScreenDensity = &ScreenDensityInt
	}
	
	if Fingerprint, ok := JourneysessioneventsnotificationdeviceMap["fingerprint"].(string); ok {
		o.Fingerprint = &Fingerprint
	}
    
	if OsFamily, ok := JourneysessioneventsnotificationdeviceMap["osFamily"].(string); ok {
		o.OsFamily = &OsFamily
	}
    
	if OsVersion, ok := JourneysessioneventsnotificationdeviceMap["osVersion"].(string); ok {
		o.OsVersion = &OsVersion
	}
    
	if Category, ok := JourneysessioneventsnotificationdeviceMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if Manufacturer, ok := JourneysessioneventsnotificationdeviceMap["manufacturer"].(string); ok {
		o.Manufacturer = &Manufacturer
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationdevice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
