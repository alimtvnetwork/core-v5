package stringutil

import "github.com/alimtvnetwork/core/constants"

func ToInt8Def(
	s string,
) int8 {
	return ToInt8(s, constants.Zero)
}
