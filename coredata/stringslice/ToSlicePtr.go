package stringslice

// ToSlicePtr on nil or empty makes new  &[]string{}
// else makes a pointer to self and returns
func ToSlicePtr(slice []string) (slicePtr *[]string) {
	if len(slice) == 0 {
		return &[]string{}
	}

	return &slice
}
