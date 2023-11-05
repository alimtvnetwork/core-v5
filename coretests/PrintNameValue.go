package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

// PrintNameValue
//
// Print using msgformats.PrintValuesFormat
func PrintNameValue(header string, any interface{}) {
	toString := ToStringNameValues(any)

	fmt.Printf(
		msgformats.PrintValuesFormat,
		header,
		any,
		toString)
}
