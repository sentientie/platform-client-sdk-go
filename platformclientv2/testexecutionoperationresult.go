package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Testexecutionoperationresult
type Testexecutionoperationresult struct { 
	// Step - The step number to indicate the order in which the operation was performed
	Step *int `json:"step,omitempty"`


	// Name - Name of the operation performed
	Name *string `json:"name,omitempty"`


	// Success - Indicated whether or not the operation was successful
	Success *bool `json:"success,omitempty"`


	// Result - The result of the operation
	Result *interface{} `json:"result,omitempty"`


	// VarError - Error that occurred during the operation
	VarError *Errorbody `json:"error,omitempty"`

}

// String returns a JSON representation of the model
func (o *Testexecutionoperationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
