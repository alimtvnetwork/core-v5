package converters

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

// AnyToValueString
//
// If nil then returns ""
// Or, returns %v of the value given.
func AnyToValueString(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}
	
	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem)
}
