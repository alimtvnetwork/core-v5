package stringutil

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

func SplitLeftRight(s, separator string) (left, right string) {
	splits := strings.SplitN(
		s, separator,
		ExpectedLeftRightLength)

	length := len(splits)
	first := splits[coreindexes.First]

	if length == ExpectedLeftRightLength {
		return first, splits[coreindexes.Second]
	}

	return first, constants.EmptyString
}
