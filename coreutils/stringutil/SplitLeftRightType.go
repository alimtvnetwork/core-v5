package stringutil

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

func SplitLeftRightType(s, separator string) *corestr.LeftRight {
	splits := strings.SplitN(
		s, separator,
		constants.Two)

	return corestr.LeftRightUsingSlice(splits)
}
