package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

// WithBracketsQuotation
//
// [\"%v\"]
func WithBracketsQuotation(
	source any,
) string {
	return fmt.Sprintf(
		constants.BracketQuotationWrapFormat,
		toString(source))
}
