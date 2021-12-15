package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedmutopicurireference
type Wfmuserscheduleadherenceupdatedmutopicurireference struct { 
	// Id - The ID of the resource
	Id *string `json:"id,omitempty"`


	// Name - The name of the resource
	Name *string `json:"name,omitempty"`

}

func (o *Wfmuserscheduleadherenceupdatedmutopicurireference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmuserscheduleadherenceupdatedmutopicurireference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmuserscheduleadherenceupdatedmutopicurireference) UnmarshalJSON(b []byte) error {
	var WfmuserscheduleadherenceupdatedmutopicurireferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmuserscheduleadherenceupdatedmutopicurireferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmuserscheduleadherenceupdatedmutopicurireferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := WfmuserscheduleadherenceupdatedmutopicurireferenceMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedmutopicurireference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
