package stringutil

import "strconv"

func ToInt32(
	s string,
	defVal int32,
) int32 {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return defVal
	}

	return int32(toInt)
}
