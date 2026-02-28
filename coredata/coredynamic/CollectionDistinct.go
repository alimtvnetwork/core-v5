package coredynamic

// --- Package-level deduplication for comparable types ---

// Distinct returns a new collection with duplicates removed, preserving order.
func Distinct[T comparable](col *Collection[T]) *Collection[T] {
	if col.IsEmpty() {
		return EmptyCollection[T]()
	}

	seen := make(map[T]bool, col.Length())
	result := NewCollection[T](col.Length())
	for _, item := range col.items {
		if !seen[item] {
			seen[item] = true
			result.items = append(result.items, item)
		}
	}
	return result
}

// Unique is an alias for Distinct.
func Unique[T comparable](col *Collection[T]) *Collection[T] {
	return Distinct(col)
}

// DistinctLock returns a new deduplicated collection with mutex protection.
func DistinctLock[T comparable](col *Collection[T]) *Collection[T] {
	col.Lock()
	defer col.Unlock()
	return Distinct(col)
}

// DistinctCount returns the number of unique items.
func DistinctCount[T comparable](col *Collection[T]) int {
	if col.IsEmpty() {
		return 0
	}

	seen := make(map[T]bool, col.Length())
	for _, item := range col.items {
		seen[item] = true
	}
	return len(seen)
}

// IsDistinct returns true if the collection has no duplicates.
func IsDistinct[T comparable](col *Collection[T]) bool {
	return col.Length() == DistinctCount(col)
}
