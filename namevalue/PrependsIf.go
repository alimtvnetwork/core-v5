package namevalue

// PrependsIf prepends items to a generic Instance slice when isAdd is true.
func PrependsIf[K comparable, V any](
	isAdd bool,
	nameValues []Instance[K, V],
	prependingItems ...Instance[K, V],
) []Instance[K, V] {
	if !isAdd || len(prependingItems) == 0 {
		return nameValues
	}

	nameValues = append(prependingItems, nameValues...)

	return nameValues
}
