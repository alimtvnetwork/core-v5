package stringutil

import "strconv"

func ToByte(
	s string,
	defVal byte,
) byte {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return defVal
	}

	return byte(toInt)
}
