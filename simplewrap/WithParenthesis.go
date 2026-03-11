package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
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
