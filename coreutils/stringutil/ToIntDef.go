package stringutil

import (
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

func ToIntDef(
	s string,
) int {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return constants.Zero
	}

	return toInt
}
