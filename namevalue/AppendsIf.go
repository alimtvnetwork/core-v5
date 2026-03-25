package namevalue

// AppendsIf appends items to a generic Instance slice when isAdd is true.
func AppendsIf[K comparable, V any](
	isAdd bool,
	nameValues []Instance[K, V],
	appendingItems ...Instance[K, V],
) []Instance[K, V] {
	if !isAdd || len(appendingItems) == 0 {
		return nameValues
	}

	nameValues = append(
		nameValues,
		appendingItems...)

	return nameValues
}
