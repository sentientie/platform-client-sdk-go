package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodyblockwithhighlight
type Documentbodyblockwithhighlight struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of the block for the body. This determines which body block object (paragraph, list, video, image or table) would have a value.
	VarType *string `json:"type,omitempty"`

	// Image - Image. It must contain a value if the type of the block is Image.
	Image *Documentbodyimage `json:"image,omitempty"`

	// Video - Video. It must contain a value if the type of the block is Video.
	Video *Documentbodyvideo `json:"video,omitempty"`

	// Paragraph - Paragraph. It must contain a value if the type of the block is Paragraph.
	Paragraph *Documentbodyparagraphwithhighlight `json:"paragraph,omitempty"`

	// List - List. It must contain a value if the type of the block is UnorderedList or OrderedList.
	List *Documentbodylistwithhighlight `json:"list,omitempty"`

	// Table - Table. It must contain a value if type of the block is Table.
	Table *Documentbodytablewithhighlight `json:"table,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Documentbodyblockwithhighlight) SetField(field string, fieldValue interface{}) {
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

func (o Documentbodyblockwithhighlight) MarshalJSON() ([]byte, error) {
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
	type Alias Documentbodyblockwithhighlight
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Image *Documentbodyimage `json:"image,omitempty"`
		
		Video *Documentbodyvideo `json:"video,omitempty"`
		
		Paragraph *Documentbodyparagraphwithhighlight `json:"paragraph,omitempty"`
		
		List *Documentbodylistwithhighlight `json:"list,omitempty"`
		
		Table *Documentbodytablewithhighlight `json:"table,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Image: o.Image,
		
		Video: o.Video,
		
		Paragraph: o.Paragraph,
		
		List: o.List,
		
		Table: o.Table,
		Alias:    (Alias)(o),
	})
}

func (o *Documentbodyblockwithhighlight) UnmarshalJSON(b []byte) error {
	var DocumentbodyblockwithhighlightMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodyblockwithhighlightMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DocumentbodyblockwithhighlightMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Image, ok := DocumentbodyblockwithhighlightMap["image"].(map[string]interface{}); ok {
		ImageString, _ := json.Marshal(Image)
		json.Unmarshal(ImageString, &o.Image)
	}
	
	if Video, ok := DocumentbodyblockwithhighlightMap["video"].(map[string]interface{}); ok {
		VideoString, _ := json.Marshal(Video)
		json.Unmarshal(VideoString, &o.Video)
	}
	
	if Paragraph, ok := DocumentbodyblockwithhighlightMap["paragraph"].(map[string]interface{}); ok {
		ParagraphString, _ := json.Marshal(Paragraph)
		json.Unmarshal(ParagraphString, &o.Paragraph)
	}
	
	if List, ok := DocumentbodyblockwithhighlightMap["list"].(map[string]interface{}); ok {
		ListString, _ := json.Marshal(List)
		json.Unmarshal(ListString, &o.List)
	}
	
	if Table, ok := DocumentbodyblockwithhighlightMap["table"].(map[string]interface{}); ok {
		TableString, _ := json.Marshal(Table)
		json.Unmarshal(TableString, &o.Table)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodyblockwithhighlight) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
