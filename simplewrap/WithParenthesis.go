package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

// WithParenthesis
//
// (%v)
func WithParenthesis(
	source any,
) string {
	return fmt.Sprintf(
		constants.ParenthesisWrapFormat,
		toString(source))
}
