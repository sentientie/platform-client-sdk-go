package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodylistitemproperties
type Documentbodylistitemproperties struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// BackgroundColor - The background color for the list item. The valid values in hex color code representation. For example black color - #000000
	BackgroundColor *string `json:"backgroundColor,omitempty"`

	// Align - The align type for the list item.
	Align *string `json:"align,omitempty"`

	// Indentation - The indentation property for the list item. The valid values in 'em'.
	Indentation *float32 `json:"indentation,omitempty"`

	// FontSize - The font size for the list item. The valid values in 'em'.
	FontSize *string `json:"fontSize,omitempty"`

	// FontType - The font type for the list item.
	FontType *string `json:"fontType,omitempty"`

	// TextColor - The text color for the list item. The valid values in hex color code representation. For example black color - #000000
	TextColor *string `json:"textColor,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Documentbodylistitemproperties) SetField(field string, fieldValue interface{}) {
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

func (o Documentbodylistitemproperties) MarshalJSON() ([]byte, error) {
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
	type Alias Documentbodylistitemproperties
	
	return json.Marshal(&struct { 
		BackgroundColor *string `json:"backgroundColor,omitempty"`
		
		Align *string `json:"align,omitempty"`
		
		Indentation *float32 `json:"indentation,omitempty"`
		
		FontSize *string `json:"fontSize,omitempty"`
		
		FontType *string `json:"fontType,omitempty"`
		
		TextColor *string `json:"textColor,omitempty"`
		Alias
	}{ 
		BackgroundColor: o.BackgroundColor,
		
		Align: o.Align,
		
		Indentation: o.Indentation,
		
		FontSize: o.FontSize,
		
		FontType: o.FontType,
		
		TextColor: o.TextColor,
		Alias:    (Alias)(o),
	})
}

func (o *Documentbodylistitemproperties) UnmarshalJSON(b []byte) error {
	var DocumentbodylistitempropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodylistitempropertiesMap)
	if err != nil {
		return err
	}
	
	if BackgroundColor, ok := DocumentbodylistitempropertiesMap["backgroundColor"].(string); ok {
		o.BackgroundColor = &BackgroundColor
	}
    
	if Align, ok := DocumentbodylistitempropertiesMap["align"].(string); ok {
		o.Align = &Align
	}
    
	if Indentation, ok := DocumentbodylistitempropertiesMap["indentation"].(float64); ok {
		IndentationFloat32 := float32(Indentation)
		o.Indentation = &IndentationFloat32
	}
	
	if FontSize, ok := DocumentbodylistitempropertiesMap["fontSize"].(string); ok {
		o.FontSize = &FontSize
	}
    
	if FontType, ok := DocumentbodylistitempropertiesMap["fontType"].(string); ok {
		o.FontType = &FontType
	}
    
	if TextColor, ok := DocumentbodylistitempropertiesMap["textColor"].(string); ok {
		o.TextColor = &TextColor
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodylistitemproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
