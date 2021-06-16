package stringslice

// ToCloneSlicePtr on nil or empty makes new  &[]string{}
// else makes a copy of itself
func ToCloneSlice(slice []string) (slicePtr []string) {
	if len(slice) == 0 {
		return []string{}
	}

	return *MergeNew(slice)
}
