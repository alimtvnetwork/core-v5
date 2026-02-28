package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

// WithBrackets
//
// [%v]
func WithBrackets(
	source any,
) string {
	toStr := toString(source)
	
	return fmt.Sprintf(
		constants.BracketWrapFormat,
		toStr)
}
