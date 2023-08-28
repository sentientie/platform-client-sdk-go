package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Gamificationscorecardchangetopicperformancemetric
type Gamificationscorecardchangetopicperformancemetric struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Metric
	Metric *Gamificationscorecardchangetopicmetric `json:"metric,omitempty"`

	// Points
	Points *int `json:"points,omitempty"`

	// Value
	Value *float32 `json:"value,omitempty"`

	// PunctualityEvents
	PunctualityEvents *[]Gamificationscorecardchangetopicpunctualityevent `json:"punctualityEvents,omitempty"`

	// EvaluationDetails
	EvaluationDetails *[]Gamificationscorecardchangetopicevaluationdetail `json:"evaluationDetails,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Gamificationscorecardchangetopicperformancemetric) SetField(field string, fieldValue interface{}) {
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

func (o Gamificationscorecardchangetopicperformancemetric) MarshalJSON() ([]byte, error) {
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
	type Alias Gamificationscorecardchangetopicperformancemetric
	
	return json.Marshal(&struct { 
		Metric *Gamificationscorecardchangetopicmetric `json:"metric,omitempty"`
		
		Points *int `json:"points,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		
		PunctualityEvents *[]Gamificationscorecardchangetopicpunctualityevent `json:"punctualityEvents,omitempty"`
		
		EvaluationDetails *[]Gamificationscorecardchangetopicevaluationdetail `json:"evaluationDetails,omitempty"`
		Alias
	}{ 
		Metric: o.Metric,
		
		Points: o.Points,
		
		Value: o.Value,
		
		PunctualityEvents: o.PunctualityEvents,
		
		EvaluationDetails: o.EvaluationDetails,
		Alias:    (Alias)(o),
	})
}

func (o *Gamificationscorecardchangetopicperformancemetric) UnmarshalJSON(b []byte) error {
	var GamificationscorecardchangetopicperformancemetricMap map[string]interface{}
	err := json.Unmarshal(b, &GamificationscorecardchangetopicperformancemetricMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := GamificationscorecardchangetopicperformancemetricMap["metric"].(map[string]interface{}); ok {
		MetricString, _ := json.Marshal(Metric)
		json.Unmarshal(MetricString, &o.Metric)
	}
	
	if Points, ok := GamificationscorecardchangetopicperformancemetricMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	
	if Value, ok := GamificationscorecardchangetopicperformancemetricMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
    
	if PunctualityEvents, ok := GamificationscorecardchangetopicperformancemetricMap["punctualityEvents"].([]interface{}); ok {
		PunctualityEventsString, _ := json.Marshal(PunctualityEvents)
		json.Unmarshal(PunctualityEventsString, &o.PunctualityEvents)
	}
	
	if EvaluationDetails, ok := GamificationscorecardchangetopicperformancemetricMap["evaluationDetails"].([]interface{}); ok {
		EvaluationDetailsString, _ := json.Marshal(EvaluationDetails)
		json.Unmarshal(EvaluationDetailsString, &o.EvaluationDetails)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Gamificationscorecardchangetopicperformancemetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
