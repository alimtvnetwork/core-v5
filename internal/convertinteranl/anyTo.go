package convertinteranl

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

type anyTo struct{}

// String
//
// If nil then returns ""
// Or, returns %v of the value given.
func (it anyTo) String(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}
	
	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem)
}

// SmartString
//
//  - If nil return ""
//  - If string return just returns
//  - Or, else return %v of value
func (it anyTo) SmartString(anyItem interface{}) string {
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
