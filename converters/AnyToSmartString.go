package converters

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

// AnyToSmartString
//
//  - If nil return ""
//  - If string return just returns
//  - Or, else return %v of value
func AnyToSmartString(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}
	
	toStr, isSuccess := anyItem.(string)
	
	if isSuccess {
		return toStr
	}
	
	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem)
}
