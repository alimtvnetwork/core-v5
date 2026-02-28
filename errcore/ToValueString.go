package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func ToValueString(reference any) string {
	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		reference)
}
