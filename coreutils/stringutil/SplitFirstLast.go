package stringutil

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coreindexes"
)

func SplitFirstLast(s, separator string) (first, last string) {
	splits := strings.Split(
		s, separator,
	)

	length := len(splits)
	first = splits[coreindexes.First]

	if length > 1 {
		return first, splits[length-1]
	}

	return first, constants.EmptyString
}
