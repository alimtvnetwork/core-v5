package corecsv

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func StringersToStringDefault(
	stringerFunctions ...fmt.Stringer,
) string {
	return StringersToString(
		constants.CommaSpace,
		true,
		false,
		stringerFunctions...)
}
