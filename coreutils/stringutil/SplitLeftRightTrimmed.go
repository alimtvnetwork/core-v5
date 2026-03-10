package stringutil

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coreindexes"
)

func SplitLeftRightTrimmed(s, separator string) (left, right string) {
	splits := strings.SplitN(
		s, separator,
		ExpectedLeftRightLength)

	length := len(splits)
	first := splits[coreindexes.First]

	if length == ExpectedLeftRightLength {
		return strings.TrimSpace(first), strings.TrimSpace(splits[coreindexes.Second])
	}

	return strings.TrimSpace(first), constants.EmptyString
}
