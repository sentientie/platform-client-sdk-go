package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Twitterdataingestionruleversionresponse
type Twitterdataingestionruleversionresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - ID of the data ingestion rule.
	Id *string `json:"id,omitempty"`

	// Name - The name of the data ingestion rule.
	Name *string `json:"name,omitempty"`

	// Description - A description of the data ingestion rule.
	Description *string `json:"description,omitempty"`

	// Status - The status of the data ingestion rule.
	Status *string `json:"status,omitempty"`

	// Version - The version number of the data ingestion rule.
	Version *int `json:"version,omitempty"`

	// DateCreated - Timestamp indicating when the data ingestion rule was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Timestamp indicating when the data ingestion rule was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Platform - The platform of the data ingestion rule.
	Platform *string `json:"platform,omitempty"`

	// SearchTerms - Search terms for X (formally Twitter).
	SearchTerms *string `json:"searchTerms,omitempty"`

	// Countries - ISO 3166-1 alpha-2 country codes where Data Ingestion Rules should apply. Defaults to worldwide.
	Countries *[]string `json:"countries,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Twitterdataingestionruleversionresponse) SetField(field string, fieldValue interface{}) {
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

func (o Twitterdataingestionruleversionresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Twitterdataingestionruleversionresponse
	
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
		
		Status *string `json:"status,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Platform *string `json:"platform,omitempty"`
		
		SearchTerms *string `json:"searchTerms,omitempty"`
		
		Countries *[]string `json:"countries,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Status: o.Status,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Platform: o.Platform,
		
		SearchTerms: o.SearchTerms,
		
		Countries: o.Countries,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Twitterdataingestionruleversionresponse) UnmarshalJSON(b []byte) error {
	var TwitterdataingestionruleversionresponseMap map[string]interface{}
	err := json.Unmarshal(b, &TwitterdataingestionruleversionresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TwitterdataingestionruleversionresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TwitterdataingestionruleversionresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := TwitterdataingestionruleversionresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Status, ok := TwitterdataingestionruleversionresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Version, ok := TwitterdataingestionruleversionresponseMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := TwitterdataingestionruleversionresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := TwitterdataingestionruleversionresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Platform, ok := TwitterdataingestionruleversionresponseMap["platform"].(string); ok {
		o.Platform = &Platform
	}
    
	if SearchTerms, ok := TwitterdataingestionruleversionresponseMap["searchTerms"].(string); ok {
		o.SearchTerms = &SearchTerms
	}
    
	if Countries, ok := TwitterdataingestionruleversionresponseMap["countries"].([]interface{}); ok {
		CountriesString, _ := json.Marshal(Countries)
		json.Unmarshal(CountriesString, &o.Countries)
	}
	
	if SelfUri, ok := TwitterdataingestionruleversionresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Twitterdataingestionruleversionresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
