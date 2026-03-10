package stringutil

import (
	"github.com/alimtvnetwork/core/coredata/corestr"
)

func SplitLeftRightTypeTrimmed(s, separator string) *corestr.LeftRight {
	left, right := SplitLeftRightTrimmed(s, separator)

	return corestr.NewLeftRight(left, right)
}
