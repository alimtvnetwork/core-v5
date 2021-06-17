package stringutil

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

func SplitLeftRightType(s, separator string) *corestr.LeftRight {
	splits := strings.SplitN(
		s, separator,
		constants.Two)

	return corestr.LeftRightUsingSlice(splits)
}
