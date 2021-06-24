package stringutil

import "gitlab.com/evatix-go/core/constants"

func ToInt8Def(
	s string,
) int8 {
	return ToInt8(s, constants.Zero)
}
