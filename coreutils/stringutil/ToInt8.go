package stringutil

import "strconv"

func ToInt8(
	s string,
	defVal int8,
) int8 {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return defVal
	}

	return int8(toInt)
}
