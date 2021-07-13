package stringslice

import "sort"

func SortIf(
	isSort bool,
	slice []string,
) []string {
	if isSort {
		sort.Strings(slice)
	}

	return slice
}
