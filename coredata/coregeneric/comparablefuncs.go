package coregeneric

// =============================================================================
// Comparable constraint functions for Collection[T]
//
// These functions require T to satisfy comparable (==, != operators)
// and provide equality-based searches without custom predicates.
// =============================================================================

// ContainsAll returns true if the collection contains all given items.
func ContainsAll[T comparable](
	source *Collection[T],
	items ...T,
) bool {
	if source == nil {
		return false
	}

	for _, item := range items {
		if !ContainsItem(source, item) {
			return false
		}
	}

	return true
}

// ContainsAny returns true if the collection contains any of the given items.
func ContainsAny[T comparable](
	source *Collection[T],
	items ...T,
) bool {
	if source == nil {
		return false
	}

	for _, item := range items {
		if ContainsItem(source, item) {
			return true
		}
	}

	return false
}

// RemoveItem removes the first occurrence of item. Returns true if found.
func RemoveItem[T comparable](
	source *Collection[T],
	item T,
) bool {
	if source == nil {
		return false
	}

	index := IndexOfItem(source, item)
	if index < 0 {
		return false
	}

	return source.RemoveAt(index)
}

// RemoveAllItems removes all occurrences of item. Returns the count removed.
func RemoveAllItems[T comparable](
	source *Collection[T],
	item T,
) int {
	if source == nil {
		return 0
	}

	removed := 0
	newItems := make([]T, 0, source.Length())

	for _, existing := range source.items {
		if existing == item {
			removed++
		} else {
			newItems = append(newItems, existing)
		}
	}

	source.items = newItems

	return removed
}

// ToHashset converts a Collection[T] to a Hashset[T].
// Requires T to be comparable for map key usage.
func ToHashset[T comparable](
	source *Collection[T],
) *Hashset[T] {
	if source == nil {
		return EmptyHashset[T]()
	}

	return HashsetFrom[T](source.items)
}

// DistinctSimpleSlice returns a new SimpleSlice with duplicates removed.
func DistinctSimpleSlice[T comparable](
	source *SimpleSlice[T],
) *SimpleSlice[T] {
	if source == nil {
		return EmptySimpleSlice[T]()
	}

	seen := make(map[T]bool)
	result := EmptySimpleSlice[T]()

	for _, item := range *source {
		if !seen[item] {
			seen[item] = true
			result.Add(item)
		}
	}

	return result
}

// ContainsSimpleSliceItem checks if a comparable item exists in a SimpleSlice.
func ContainsSimpleSliceItem[T comparable](
	source *SimpleSlice[T],
	item T,
) bool {
	if source == nil {
		return false
	}

	for _, existing := range *source {
		if existing == item {
			return true
		}
	}

	return false
}
