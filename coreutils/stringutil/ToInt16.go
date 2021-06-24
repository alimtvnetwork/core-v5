package stringutil

import "strconv"

func ToInt16(
	s string,
	defVal int16,
) int16 {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return defVal
	}

	return int16(toInt)
}
