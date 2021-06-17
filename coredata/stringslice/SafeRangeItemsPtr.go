package stringslice

func SafeRangeItemsPtr(
	slice *[]string,
	start, end int,
) *[]string {
	if slice == nil || *slice == nil {
		return &[]string{}
	}

	results := SafeRangeItems(*slice, start, end)

	return &results
}
