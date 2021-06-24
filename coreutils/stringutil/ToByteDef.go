package stringutil

import "gitlab.com/evatix-go/core/constants"

func ToByteDef(
	s string,
) byte {
	return ToByte(s, constants.Zero)
}
