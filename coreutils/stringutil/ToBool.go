package stringutil

import "strconv"

func ToBool(
	s string,
) bool {
	if s == "" {
		return false
	}

	switch s {
	case "yes", "Yes", "YES":
		return true
	case "no", "NO", "No":
		return false
	}

	isBool, err := strconv.ParseBool(s)

	if err != nil {
		return false
	}

	return isBool
}
