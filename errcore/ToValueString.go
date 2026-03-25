package errcore

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func ToValueString(reference any) string {
	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		reference)
}
