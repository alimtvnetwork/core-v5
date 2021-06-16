package stringslice

// ToCloneSimpleSliceToPointers on nil or empty makes new  &[]string{}
// else makes a copy of itself
func ToCloneSimpleSliceToPointers(slice []string) (slicePtr *[]string) {
	if IsEmpty(slice) {
		return &[]string{}
	}

	return MergeNew(slice)
}
