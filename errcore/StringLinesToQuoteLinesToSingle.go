package errcore

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

// StringLinesToQuoteLinesToSingle
//
// Each line will be wrapped with "\"%s\", quotation and comma
func StringLinesToQuoteLinesToSingle(lines []string) string {
	slice := StringLinesToQuoteLines(lines)

	return strings.Join(
		slice,
		constants.NewLineUnix)
}
