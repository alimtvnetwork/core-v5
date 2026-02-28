package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

type printMessage struct{}

func (it printMessage) FailedExpected(
	isFailed bool,
	when,
	actual,
	expected any,
	counter int,
) {
	if isFailed {
		message := GetAssert.Quick(when, actual, expected, counter)

		fmt.Println(message)
	}
}

// PrintNameValue
//
// Print using msgformats.PrintValuesFormat
func (it printMessage) NameValue(header string, anyItem any) {
	toString := ToStringNameValues(anyItem)

	fmt.Printf(
		msgformats.PrintValuesFormat,
		header,
		anyItem,
		toString,
	)
}

// PrintValue
//
// Print values using msgformats.PrintValuesFormat
func (it printMessage) Value(header string, anyItem any) {
	toString := ToStringValues(anyItem)

	fmt.Printf(
		msgformats.PrintValuesFormat,
		header,
		anyItem,
		toString,
	)
}
