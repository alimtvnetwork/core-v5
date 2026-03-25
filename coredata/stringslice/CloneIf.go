package stringslice

import "strings"

func CloneIf(
	isClone bool,
	additionalCap int,
	slice []string,
) (newSlice []string) {
	isSkipClone := !isClone

	if slice == nil && isSkipClone {
		return []string{}
	}

	if isSkipClone {
		return slice
	}

	return CloneUsingCap(additionalCap, slice)
}

// JoinWith
//
// joiner + strings.Join(items, joiner)
func JoinWith(joiner string, items ...string) string {
	if len(items) == 0 {
		return ""
	}

	return joiner + strings.Join(items, joiner)
}

func Joins(joiner string, items ...string) string {
	if len(items) == 0 {
		return ""
	}

	return strings.Join(items, joiner)
}
