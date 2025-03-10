package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignruleactionentities
type Campaignruleactionentities struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Campaigns - The list of campaigns for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a campaign.
	Campaigns *[]Domainentityref `json:"campaigns,omitempty"`

	// Sequences - The list of sequences for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a sequence.
	Sequences *[]Domainentityref `json:"sequences,omitempty"`

	// EmailCampaigns - The list of Email campaigns for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a Email campaign.
	EmailCampaigns *[]Domainentityref `json:"emailCampaigns,omitempty"`

	// SmsCampaigns - The list of SMS campaigns for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a SMS campaign.
	SmsCampaigns *[]Domainentityref `json:"smsCampaigns,omitempty"`

	// UseTriggeringEntity - If true, the CampaignRuleAction will apply to the same entity that triggered the CampaignRuleCondition.
	UseTriggeringEntity *bool `json:"useTriggeringEntity,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaignruleactionentities) SetField(field string, fieldValue interface{}) {
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

func (o Campaignruleactionentities) MarshalJSON() ([]byte, error) {
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
	type Alias Campaignruleactionentities
	
	return json.Marshal(&struct { 
		Campaigns *[]Domainentityref `json:"campaigns,omitempty"`
		
		Sequences *[]Domainentityref `json:"sequences,omitempty"`
		
		EmailCampaigns *[]Domainentityref `json:"emailCampaigns,omitempty"`
		
		SmsCampaigns *[]Domainentityref `json:"smsCampaigns,omitempty"`
		
		UseTriggeringEntity *bool `json:"useTriggeringEntity,omitempty"`
		Alias
	}{ 
		Campaigns: o.Campaigns,
		
		Sequences: o.Sequences,
		
		EmailCampaigns: o.EmailCampaigns,
		
		SmsCampaigns: o.SmsCampaigns,
		
		UseTriggeringEntity: o.UseTriggeringEntity,
		Alias:    (Alias)(o),
	})
}

func (o *Campaignruleactionentities) UnmarshalJSON(b []byte) error {
	var CampaignruleactionentitiesMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignruleactionentitiesMap)
	if err != nil {
		return err
	}
	
	if Campaigns, ok := CampaignruleactionentitiesMap["campaigns"].([]interface{}); ok {
		CampaignsString, _ := json.Marshal(Campaigns)
		json.Unmarshal(CampaignsString, &o.Campaigns)
	}
	
	if Sequences, ok := CampaignruleactionentitiesMap["sequences"].([]interface{}); ok {
		SequencesString, _ := json.Marshal(Sequences)
		json.Unmarshal(SequencesString, &o.Sequences)
	}
	
	if EmailCampaigns, ok := CampaignruleactionentitiesMap["emailCampaigns"].([]interface{}); ok {
		EmailCampaignsString, _ := json.Marshal(EmailCampaigns)
		json.Unmarshal(EmailCampaignsString, &o.EmailCampaigns)
	}
	
	if SmsCampaigns, ok := CampaignruleactionentitiesMap["smsCampaigns"].([]interface{}); ok {
		SmsCampaignsString, _ := json.Marshal(SmsCampaigns)
		json.Unmarshal(SmsCampaignsString, &o.SmsCampaigns)
	}
	
	if UseTriggeringEntity, ok := CampaignruleactionentitiesMap["useTriggeringEntity"].(bool); ok {
		o.UseTriggeringEntity = &UseTriggeringEntity
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignruleactionentities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
