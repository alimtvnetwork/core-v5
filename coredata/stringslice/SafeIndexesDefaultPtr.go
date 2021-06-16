package stringslice

// SafeIndexesDefaultPtr Only indexes which are present values will be included.
//
// Warning : Not found indexes will not be included in the values.
func SafeIndexesDefaultPtr(slice *[]string, indexes ...int) (values *[]string) {
	if IsEmptyPtr(slice) {
		return &[]string{}
	}

	values2 := SafeIndexesDefault(*slice, indexes...)

	return &values2
}
