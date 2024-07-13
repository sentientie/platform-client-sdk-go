package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webeventdevice
type Webeventdevice struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Category - Device category.
	Category *string `json:"category,omitempty"`

	// VarType - Device type (e.g. iPad, iPhone, Other).
	VarType *string `json:"type,omitempty"`

	// IsMobile - Flag that is true for mobile devices.
	IsMobile *bool `json:"isMobile,omitempty"`

	// ScreenHeight - Device's screen height.
	ScreenHeight *int `json:"screenHeight,omitempty"`

	// ScreenWidth - Device's screen width.
	ScreenWidth *int `json:"screenWidth,omitempty"`

	// ScreenDensity - Device's screen density, measured as a scale factor where a value of 1 represents a baseline 1:1 ratio of pixels to logical (device-independent) pixels.
	ScreenDensity *int `json:"screenDensity,omitempty"`

	// OsFamily - Operating system family.
	OsFamily *string `json:"osFamily,omitempty"`

	// OsVersion - Operating system version.
	OsVersion *string `json:"osVersion,omitempty"`

	// Manufacturer - Manufacturer of the device.
	Manufacturer *string `json:"manufacturer,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webeventdevice) SetField(field string, fieldValue interface{}) {
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

func (o Webeventdevice) MarshalJSON() ([]byte, error) {
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
	type Alias Webeventdevice
	
	return json.Marshal(&struct { 
		Category *string `json:"category,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		IsMobile *bool `json:"isMobile,omitempty"`
		
		ScreenHeight *int `json:"screenHeight,omitempty"`
		
		ScreenWidth *int `json:"screenWidth,omitempty"`
		
		ScreenDensity *int `json:"screenDensity,omitempty"`
		
		OsFamily *string `json:"osFamily,omitempty"`
		
		OsVersion *string `json:"osVersion,omitempty"`
		
		Manufacturer *string `json:"manufacturer,omitempty"`
		Alias
	}{ 
		Category: o.Category,
		
		VarType: o.VarType,
		
		IsMobile: o.IsMobile,
		
		ScreenHeight: o.ScreenHeight,
		
		ScreenWidth: o.ScreenWidth,
		
		ScreenDensity: o.ScreenDensity,
		
		OsFamily: o.OsFamily,
		
		OsVersion: o.OsVersion,
		
		Manufacturer: o.Manufacturer,
		Alias:    (Alias)(o),
	})
}

func (o *Webeventdevice) UnmarshalJSON(b []byte) error {
	var WebeventdeviceMap map[string]interface{}
	err := json.Unmarshal(b, &WebeventdeviceMap)
	if err != nil {
		return err
	}
	
	if Category, ok := WebeventdeviceMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if VarType, ok := WebeventdeviceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if IsMobile, ok := WebeventdeviceMap["isMobile"].(bool); ok {
		o.IsMobile = &IsMobile
	}
    
	if ScreenHeight, ok := WebeventdeviceMap["screenHeight"].(float64); ok {
		ScreenHeightInt := int(ScreenHeight)
		o.ScreenHeight = &ScreenHeightInt
	}
	
	if ScreenWidth, ok := WebeventdeviceMap["screenWidth"].(float64); ok {
		ScreenWidthInt := int(ScreenWidth)
		o.ScreenWidth = &ScreenWidthInt
	}
	
	if ScreenDensity, ok := WebeventdeviceMap["screenDensity"].(float64); ok {
		ScreenDensityInt := int(ScreenDensity)
		o.ScreenDensity = &ScreenDensityInt
	}
	
	if OsFamily, ok := WebeventdeviceMap["osFamily"].(string); ok {
		o.OsFamily = &OsFamily
	}
    
	if OsVersion, ok := WebeventdeviceMap["osVersion"].(string); ok {
		o.OsVersion = &OsVersion
	}
    
	if Manufacturer, ok := WebeventdeviceMap["manufacturer"].(string); ok {
		o.Manufacturer = &Manufacturer
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webeventdevice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
