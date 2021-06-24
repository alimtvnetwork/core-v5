package stringutil

import "gitlab.com/evatix-go/core/constants"

func ToInt16Def(
	s string,
) int16 {
	return ToInt16(s, constants.Zero)
}
