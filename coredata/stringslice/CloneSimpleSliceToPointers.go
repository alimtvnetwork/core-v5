package stringslice

// CloneSimpleSliceToPointers on nil or empty makes new  &[]string{}
// else makes a copy of itself
func CloneSimpleSliceToPointers(slice []string) (slicePtr *[]string) {
	if IsEmpty(slice) {
		return &[]string{}
	}

	return MergeNew(slice)
}
