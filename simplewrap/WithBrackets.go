package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
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
