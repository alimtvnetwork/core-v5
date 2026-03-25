package coresorttests

import "fmt"

// toIntPtrs converts []int to []*int (each element gets its own pointer).
func toIntPtrs(values []int) []*int {
	ptrs := make([]*int, len(values))

	for i := range values {
		v := values[i]
		ptrs[i] = &v
	}

	return ptrs
}

// formatIntPtrs formats []*int to "[v1 v2 v3]" matching fmt.Sprintf("%v", []int{...}).
func formatIntPtrs(ptrs []*int) string {
	vals := make([]int, len(ptrs))

	for i, p := range ptrs {
		vals[i] = *p
	}

	return fmt.Sprintf("%v", vals)
}

// toStrPtrs converts []string to []*string.
func toStrPtrs(values []string) []*string {
	ptrs := make([]*string, len(values))

	for i := range values {
		v := values[i]
		ptrs[i] = &v
	}

	return ptrs
}

// formatStrPtrs formats []*string to "[v1 v2 v3]" matching fmt.Sprintf("%v", []string{...}).
func formatStrPtrs(ptrs []*string) string {
	vals := make([]string, len(ptrs))

	for i, p := range ptrs {
		vals[i] = *p
	}

	return fmt.Sprintf("%v", vals)
}
