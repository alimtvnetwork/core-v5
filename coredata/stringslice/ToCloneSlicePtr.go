package stringslice

// ToCloneSlicePtr on nil or empty makes new  &[]string{}
// else makes a copy of itself
func ToCloneSlicePtr(slice *[]string) (slicePtr *[]string) {
	if IsEmptyPtr(slice) {
		return &[]string{}
	}

	return MergeNewPointers(slice)
}
