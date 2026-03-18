package corecmp

func IsIntegersEqual(array, other []int) bool {
	if array == nil && other == nil {
		return true
	}

	if array == nil || other == nil {
		return false
	}

	if len(array) != len(other) {
		return false
	}

	for i := range array {
		if array[i] != other[i] {
			return false
		}
	}

	return true
}
