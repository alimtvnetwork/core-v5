package stringutil

import (
	"strconv"

	"github.com/alimtvnetwork/core/constants"
)

func ToByteDefault(
	s string,
) byte {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return constants.Zero
	}

	if toInt >= constants.Zero && toInt <= constants.MaxUnit8AsInt {
		return byte(toInt)
	}

	return constants.Zero
}
