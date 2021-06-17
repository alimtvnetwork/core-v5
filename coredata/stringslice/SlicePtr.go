package stringslice

// SlicePtr on nil or empty makes new  &[]string{}
// else makes a pointer to self and returns
func SlicePtr(slice []string) (slicePtr *[]string) {
	if len(slice) == 0 {
		return &[]string{}
	}

	return &slice
}
