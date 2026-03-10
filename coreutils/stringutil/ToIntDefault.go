package stringutil

import (
	"strconv"

	"github.com/alimtvnetwork/core/constants"
)

func ToIntDefault(
	s string,
) int {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return constants.Zero
	}

	return toInt
}
