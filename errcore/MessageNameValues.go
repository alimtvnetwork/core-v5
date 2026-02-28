package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/namevalue"
)

func MessageNameValues(
	message string,
	nameValues ...namevalue.StringAny,
) string {
	if len(nameValues) == 0 {
		return message
	}

	compiledMap := VarNameValues(nameValues...)

	return fmt.Sprintf(
		messageMapFormat,
		message,
		compiledMap)
}
