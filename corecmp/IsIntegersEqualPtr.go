package corecmp

func IsIntegersEqualPtr(leftSlicePtr, rightSlicePtr *[]int) bool {
	if leftSlicePtr == nil && rightSlicePtr == nil {
		return true
	}

	if leftSlicePtr == nil || rightSlicePtr == nil {
		return false
	}

	return IsIntegersEqual(*leftSlicePtr, *rightSlicePtr)
}
