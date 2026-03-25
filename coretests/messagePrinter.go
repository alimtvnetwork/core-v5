package coretests

import (
	"log/slog"
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

		slog.Warn("test failed", "message", message)
	}
}

// PrintNameValue
//
// Print using msgformats.PrintValuesFormat
func (it printMessage) NameValue(header string, anyItem any) {
	toString := ToStringNameValues(anyItem)

	slog.Info("name-value",
		"header", header,
		"value", anyItem,
		"toString", toString,
	)
}

// PrintValue
//
// Print values using msgformats.PrintValuesFormat
func (it printMessage) Value(header string, anyItem any) {
	toString := ToStringValues(anyItem)

	slog.Info("value",
		"header", header,
		"value", anyItem,
		"toString", toString,
	)
}
